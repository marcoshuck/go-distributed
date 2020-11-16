package main

import (
	"context"
	"fmt"
	"github.com/marcoshuck/go-distributed/challenge_2/pkg/cloud"
	"github.com/marcoshuck/go-distributed/challenge_2/pkg/consumer"
	"github.com/marcoshuck/go-distributed/challenge_2/pkg/queue"
	"github.com/marcoshuck/go-distributed/challenge_2/pkg/worker"
	"github.com/streadway/amqp"
	"google.golang.org/grpc"
	"log"
	"os"
	"time"
)

func main() {
	user := os.Getenv("RABBITMQ_USER")
	password := os.Getenv("RABBITMQ_PASSWORD")
	host := os.Getenv("RABBITMQ_HOST")
	port := os.Getenv("RABBITMQ_PORT")

	grpcHost := os.Getenv("CLOUD_PROVIDER_HOST")
	grpcPort := os.Getenv("CLOUD_PROVIDER_PORT")
	grpcAddr := fmt.Sprintf("%s:%s", grpcHost, grpcPort)

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

	log.Println("Attempting to declare or use queue:", queue.SimulationDeployment)
	_, err = queue.NewAMQPQueue(channel, queue.SimulationDeployment)
	if err != nil {
		log.Fatal("error while opening queue:", err)
	}

	log.Println("Setting up new consumer to connect to the simulations microservice")
	c := consumer.NewConsumer(channel)

	log.Println("Connecting to cloud provider microservice")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	grpcConn, err := grpc.DialContext(ctx, grpcAddr, grpc.WithInsecure(), grpc.WithBlock())
	defer grpcConn.Close()
	if err != nil {
		log.Fatal("error while connecting to cloud provider microservice:", err)
	}

	cloudProvider := cloud.NewCloudClient(grpcConn)

	w := worker.NewWorker(c, cloudProvider)

	log.Println("Initializing worker microservice")
	err = w.Init()
	if err != nil {
		log.Fatal("error while initializing worker service:", err)
	}

	log.Println("Worker microservice initialized")

	select {}
}
