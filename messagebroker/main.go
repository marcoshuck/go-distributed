package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	conn *amqp.Connection
)

func init() {
	connectMessageBroker("localhost", "guest", "guest", 5672)
}

func connectMessageBroker(host, user, password string, port uint16) {
	var err error
	conn, err = amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%d/", user, password, host, port))
	if err != nil {
		log.Fatalf("Error while connecting to RabbitMQ: %s", err)
	}
}

func main() {
	defer conn.Close()

	// Set up Ctrl + C signal termination
	var exit bool
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		log.Println("\rClosing server...")
		exit = true
	}()

	// Set up channel
	ch, err := setupChannel(conn)
	if err != nil {
		log.Println(fmt.Sprintf("Failed to initialize channel: %s", err))
		return
	}
	defer ch.Close()

	// Set up queue to send and receive messages
	q, err := setupQueue(ch)
	if err != nil {
		log.Println(fmt.Sprintf("Failed to initialize channel: %s", err))
		return
	}

	// Set up workers
	for i := 1; i <= 10; i++ {
		go func(publisherID int) {
			for i := 0; i < 999999; i++ {
				publish(ch, q, fmt.Sprintf("Message [%d] from publisher [%d]", i, publisherID))
				log.Println(fmt.Sprintf("Publisher [%d]: [%d]", publisherID, i))
				time.Sleep(time.Duration(publisherID) * time.Second)
			}
		}(i)
	}

	// Set up consumers
	msgs, err := setupConsumer(ch, q)
	if err != nil {
		log.Println(fmt.Sprintf("Failed to initialize channel: %s", err))
		return
	}

	go func() {
		for d := range msgs {
			log.Printf("Consumer: %s", d.Body)
		}
	}()

	for i := 0; !exit; i++ {
		log.Println(fmt.Sprintf("APPLICATION STATUS REPORT [%d], running for: %d seconds", i, (i+1)*5))
		time.Sleep(5 * time.Second)
	}
}

func setupChannel(c *amqp.Connection) (*amqp.Channel, error) {
	return c.Channel()
}

func setupQueue(ch *amqp.Channel) (amqp.Queue, error) {
	return ch.QueueDeclare("app", false, false, false, false, nil)
}

func publish(ch *amqp.Channel, q amqp.Queue, message string) error {
	msg := amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(message),
	}
	return ch.Publish("", q.Name, false, false, msg)
}

func setupConsumer(ch *amqp.Channel, q amqp.Queue) (<-chan amqp.Delivery, error) {
	return ch.Consume(q.Name, "", true, false, false, false, nil)
}
