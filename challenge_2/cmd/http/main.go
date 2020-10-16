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

	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%s/", user, password, host, port))
	defer conn.Close()
	if err != nil {
		log.Fatal("error while connecting to RabbitMQ:", err)
	}

	channel, err := conn.Channel()
	defer channel.Close()
	if err != nil {
		log.Fatal("error while creating new channel:", err)
	}

	_, err = queue.NewAMQPQueue(channel, queue.SimulationRequests)
	if err != nil {
		log.Fatal("error while opening queue:", queue.SimulationRequests)
	}

	send := sender.NewSender(channel)

	ctrl := simulations.NewController(send)

	server.Route("POST", "/simulations", ctrl.Create)

	err = server.Run(8000)
	if err != nil {
		log.Fatal("error while running http server:", err)
	}
}
