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
	"os/signal"
	"syscall"
	"time"
)

func main() {
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
		log.Fatal("error while opening queue:", err)
	}

	_, err = queue.NewAMQPQueue(channel, queue.SimulationDeployment)
	if err != nil {
		log.Fatal("error while opening queue:", err)
	}

	send := sender.NewSender(channel)
	cons := consumer.NewConsumer(channel)

	repository := simulations.NewRepository()

	srv := simulations.NewService(repository, send, cons)

	err = srv.Init()
	if err != nil {
		log.Fatal("error while initializing simulation service:", err)
	}

	var exit bool
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		log.Println("\rClosing simulation service...")
		exit = true
	}()

	for !exit {
		time.Sleep(5 * time.Second)
	}
}
