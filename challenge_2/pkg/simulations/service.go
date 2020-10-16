package simulations

import (
	"encoding/json"
	"github.com/marcoshuck/go-distributed/challenge_2/pkg/consumer"
	"github.com/marcoshuck/go-distributed/challenge_2/pkg/queue"
	"github.com/marcoshuck/go-distributed/challenge_2/pkg/sender"
	"github.com/streadway/amqp"
	"log"
)

type Service interface {
	Init() error
	Create(input CreateSimulationInput) (CreateSimulationOutput, error)
}

type service struct {
	repository Repository
	sender     sender.Sender
	consumer   consumer.Consumer
}

func (s *service) Init() error {
	if err := s.initializeConsumer(); err != nil {
		return err
	}
	return nil
}

func (s *service) Create(input CreateSimulationInput) (CreateSimulationOutput, error) {
	sim := Simulation{
		Name:  input.Name,
		Owner: input.Owner,
		Image: input.Image,
	}

	var err error
	sim, err = s.repository.Create(sim)
	if err != nil {
		return CreateSimulationOutput{}, err
	}

	return CreateSimulationOutput{
		ID:    sim.ID,
		Name:  sim.Name,
		Owner: sim.Owner,
		Image: sim.Image,
	}, nil
}

func (s *service) initializeConsumer() error {
	err := s.consumer.Consume(queue.SimulationRequests, s.createSimulationRequest)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) createSimulationRequest(deliveries <-chan amqp.Delivery) {
	var simulationCreate CreateSimulationInput
	for d := range deliveries {
		err := json.Unmarshal(d.Body, &simulationCreate)
		if err != nil {
			log.Println("error while unmarshalling body message received from queue:", err)
			continue
		}

		output, err := s.Create(simulationCreate)
		if err != nil {
			log.Println("error while creating simulation:", err)
			continue
		}

		b, err := json.Marshal(&output)
		if err != nil {
			log.Println("error while marshalling body message to create simulation deployment:", err)
			continue
		}

		err = s.sender.Send("", queue.SimulationDeployment, b, "text/json")
		if err != nil {
			log.Println("error while sending simulation deployment to the queue:", err)
			continue
		}
	}
}

func NewService(repository Repository, sender sender.Sender, consumer consumer.Consumer) Service {
	return &service{
		repository: repository,
		sender:     sender,
		consumer:   consumer,
	}
}
