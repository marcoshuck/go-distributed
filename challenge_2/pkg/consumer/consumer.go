package consumer

import (
	"github.com/streadway/amqp"
	"log"
)

type Consumer interface {
	Consume(key string, fn func(<-chan amqp.Delivery)) error
}

type consumer struct {
	channel *amqp.Channel
}

func (r *consumer) Consume(key string, fn func(<-chan amqp.Delivery)) error {
	messages, err := r.channel.Consume(key, "", true, false, false, false, nil)
	if err != nil {
		log.Println("failed to set up a new consumer:", err)
		return err
	}
	go fn(messages)
	return nil
}

func NewConsumer(channel *amqp.Channel) Consumer {
	return &consumer{
		channel: channel,
	}
}
