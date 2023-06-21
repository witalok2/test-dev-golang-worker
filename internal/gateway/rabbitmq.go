package rabbitmq

import (
	"errors"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/witalok2/test-dev-golang-worker/config"
)

type QueueClient struct {
	Conn    *amqp.Connection
	Channel *amqp.Channel

	QueueName string
}

func NewQueueClient(env config.Queue) (*QueueClient, error) {
	conn, err := amqp.Dial(env.URI)
	if err != nil {
		return nil, errors.New("failed to connect to RabbitMQ")
	}

	channel, err := conn.Channel()
	if err != nil {
		return nil, errors.New("failed to open a channel")
	}

	return &QueueClient{
		Conn:      conn,
		Channel:   channel,
		QueueName: env.QueueName,
	}, nil
}
