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
	GetUrl() string
	GetServer() string
}

//

type RabbitMQConnection interface {
	Connect() (*amqp.Connection, error)
	Close()
	RabbitMQContext() RabbitMQContext
}

type RabbitMQChannel interface {
	Connect() (*amqp.Channel, error)
	Close()
	RabbitMQContext() RabbitMQContext
}

type RabbitMQQueue interface {
	Connect() (*amqp.Queue, error)
	Close()
	RabbitMQContext() RabbitMQContext
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
