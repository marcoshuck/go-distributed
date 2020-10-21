package main

import (
	"fmt"
	"github.com/marcoshuck/go-distributed/challenge_2/pkg/consumer"
	"github.com/marcoshuck/go-distributed/challenge_2/pkg/queue"
	"github.com/marcoshuck/go-distributed/challenge_2/pkg/sender"
	"github.com/marcoshuck/go-distributed/challenge_2/pkg/simulations"
	"github.com/streadway/amqp"
	"log"
	"os"
)

func main() {
	user := os.Getenv("RABBITMQ_USER")
	password := os.Getenv("RABBITMQ_PASSWORD")
	host := os.Getenv("RABBITMQ_HOST")
	port := os.Getenv("RABBITMQ_PORT")

	log.Printf("Connecting to the AMQP server on %s:%s\n", host, port)

	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%s/", user, password, host, port))
	defer conn.Close()
	if err != nil {
		log.Fatal("error while connecting to RabbitMQ:", err)
	}

	log.Println("Attempting to open new channel")

	channel, err := conn.Channel()
	defer channel.Close()
	if err != nil {
		log.Fatal("error while creating new channel:", err)
	}

	log.Println("Attempting to declare or use queue:", queue.SimulationRequests)
	_, err = queue.NewAMQPQueue(channel, queue.SimulationRequests)
	if err != nil {
		log.Fatal("error while opening queue:", err)
	}

	log.Println("Attempting to declare or use queue:", queue.SimulationDeployment)
	_, err = queue.NewAMQPQueue(channel, queue.SimulationDeployment)
	if err != nil {
		log.Fatal("error while opening queue:", err)
	}

	log.Println("Setting up new sender to connect to the worker pool microservice")
	send := sender.NewSender(channel)

	log.Println("Setting up new consumer to connect to the http microservice")
	cons := consumer.NewConsumer(channel)

	repository := simulations.NewRepository()
	srv := simulations.NewService(repository, send, cons)

	log.Println("Initializing simulations microservice")
	err = srv.Init()
	if err != nil {
		log.Fatal("error while initializing simulation service:", err)
	}

	log.Println("Simulations microservice initialized")

	select {}
}
