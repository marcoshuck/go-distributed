package queue

import "github.com/streadway/amqp"

const (
	SimulationRequests   string = "simulation-requests"
	SimulationDeployment string = "simulation-deployment"
)

func NewAMQPQueue(channel *amqp.Channel, name string) (amqp.Queue, error) {
	return channel.QueueDeclare(name, false, false, false, false, nil)
}
