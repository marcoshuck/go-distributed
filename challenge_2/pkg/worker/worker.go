package worker

import (
	"encoding/json"
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
	consumer consumer.Consumer
	sender   sender.Sender
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
	time.Sleep(time.Minute * 3)
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

func NewWorker(sender sender.Sender, consumer consumer.Consumer) Worker {
	return &worker{
		sender:   sender,
		consumer: consumer,
	}
}
