package worker

import (
	"context"
	"encoding/json"
	"github.com/marcoshuck/go-distributed/challenge_2/pkg/cloud"
	"github.com/marcoshuck/go-distributed/challenge_2/pkg/consumer"
	"github.com/marcoshuck/go-distributed/challenge_2/pkg/queue"
	"github.com/marcoshuck/go-distributed/challenge_2/pkg/sender"
	"github.com/marcoshuck/go-distributed/challenge_2/pkg/simulations"
	"github.com/streadway/amqp"
	"log"
	"time"
)

type Worker interface {
	Init() error
	Do(simulation simulations.CreateSimulationOutput)
}

type worker struct {
	consumer      consumer.Consumer
	sender        sender.Sender
	cloudProvider cloud.CloudClient
}

func (w *worker) Init() error {
	err := w.consumer.Consume(queue.SimulationDeployment, w.createWorkerActivity)
	if err != nil {
		return err
	}
	return nil
}

func (w *worker) Do(simulation simulations.CreateSimulationOutput) {
	log.Printf("Launching simulation [%s] created by [%s] and with image [%s]", simulation.Name, simulation.Owner, simulation.Image)

	input := cloud.CreateMachinesRequest{
		Name:     simulation.Name,
		Provider: cloud.Provider_AWS,
		Kind:     "g3.4xlarge",
		Min:      1,
		Max:      10,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	response, err := w.cloudProvider.CreateMachines(ctx, &input)

	if err != nil {
		log.Printf("Error while creating machines, error: %s", err)
		return
	}

	log.Printf("Simulation [%s] created by [%s] with image [%s]\n", simulation.Name, simulation.Owner, simulation.Image)
	log.Printf("\t\tKind: g3.4xlarge | Count: %d\n", response.GetAmount())
	for i, instance := range response.GetInstances() {
		log.Printf("\t\t[%d] %s\n", i, instance)
	}
}

func (w *worker) createWorkerActivity(deliveries <-chan amqp.Delivery) {
	var sim simulations.CreateSimulationOutput
	for d := range deliveries {
		err := json.Unmarshal(d.Body, &sim)
		if err != nil {
			log.Println("error while unmarshalling body message received from queue:", err)
			continue
		}
		w.Do(sim)
	}
}

func NewWorker(consumer consumer.Consumer, cloudProvider cloud.CloudClient) Worker {
	return &worker{
		cloudProvider: cloudProvider,
		consumer:      consumer,
	}
}
