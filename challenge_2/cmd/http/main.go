package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/marcoshuck/go-distributed/challenge_2/pkg/http"
	"github.com/marcoshuck/go-distributed/challenge_2/pkg/queue"
	"github.com/marcoshuck/go-distributed/challenge_2/pkg/sender"
	"github.com/marcoshuck/go-distributed/challenge_2/pkg/simulations"
	"github.com/streadway/amqp"
	"log"
	"os"
)

func main() {
	r := gin.Default()
	server := http.NewServer(r)

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
		log.Fatal("error while opening new channel:", err)
	}

	log.Println("Attempting to declare or use queue:", queue.SimulationRequests)
	_, err = queue.NewAMQPQueue(channel, queue.SimulationRequests)
	if err != nil {
		log.Fatal("error while using queue:", queue.SimulationRequests)
	}

	log.Println("Setting up new sender to connect to the simulations microservice")
	send := sender.NewSender(channel)

	log.Println("Initializing simulations controller")
	ctrl := simulations.NewController(send)

	log.Println("Adding simulation routes to server router")
	server.Route("POST", "/simulations", ctrl.Create)

	err = server.Run(8000)
	if err != nil {
		log.Fatal("error while running http server:", err)
	}
}
