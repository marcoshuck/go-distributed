package sender

import (
	"github.com/streadway/amqp"
)

type Sender interface {
	Send(exchange, key string, message []byte, contentType string) error
}

type sender struct {
	channel *amqp.Channel
}

func (s *sender) Send(exchange, key string, message []byte, contentType string) error {
	msg := amqp.Publishing{
		ContentType: contentType,
		Body:        message,
	}
	return s.channel.Publish(exchange, key, false, false, msg)
}

func NewSender(channel *amqp.Channel) Sender {
	return &sender{
		channel: channel,
	}
}
