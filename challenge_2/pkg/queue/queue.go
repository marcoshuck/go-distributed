package queue

import "github.com/streadway/amqp"

const (
	SimulationRequests   string = "simulation-requests"
	SimulationDeployment string = "simulation-deployment"
	SimulationMachines   string = "simulation-machines"
)

func NewAMQPQueue(channel *amqp.Channel, name string) (amqp.Queue, error) {
	return channel.QueueDeclare(name, false, false, false, false, nil)
}
