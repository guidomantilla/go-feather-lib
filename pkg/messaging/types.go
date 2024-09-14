package messaging

import (
	"time"

	nats "github.com/nats-io/nats.go"
	amqp "github.com/rabbitmq/amqp091-go"
)

const (
	makeConnectionDelay = 2 * time.Second
)

type RabbitMQContext interface {
	Url() string
	Server() string
	FailOver() bool
}

//

type RabbitMQConnection interface {
	RabbitMQContext() RabbitMQContext
	Connect() (*amqp.Connection, error)
	Close()
}

type RabbitMQChannel interface {
	RabbitMQContext() RabbitMQContext
	Connect() (*amqp.Channel, error)
	Close()
}

type RabbitMQQueue interface {
	RabbitMQContext() RabbitMQContext
	Connect() (*amqp.Queue, error)
	Close()
}

//

var _ RabbitMQQueueConnection = (*DefaultRabbitMQQueueConnection)(nil)

type RabbitMQQueueConnection interface {
	Start()
	Close()
	Connect() (*amqp.Connection, *amqp.Channel, *amqp.Queue, <-chan amqp.Delivery, error)
}

var _ RabbitMQQueueMessageListener = (*DefaultRabbitMQQueueMessageListener)(nil)

type RabbitMQQueueMessageListener interface {
	Queue() string
	OnMessage(message *amqp.Delivery) error
}

//

type NatsSubjectConnection interface {
	Start()
	Close()
	Connect() (*nats.Conn, *nats.Subscription, chan *nats.Msg, error)
}

var _ NatsMessageListener = (*DefaultNatsMessageListener)(nil)

type NatsMessageListener interface {
	OnMessage(message *nats.Msg) error
}
