package messaging

import (
	"time"

	nats "github.com/nats-io/nats.go"
	amqp "github.com/rabbitmq/amqp091-go"
)

const (
	makeConnectionDelay = 2 * time.Second
)

//

var _ RabbitMQQueueConnection = (*DefaultRabbitMQQueueConnection)(nil)

type RabbitMQQueueConnection interface {
	Start()
	Close()
	Connect() (*amqp.Connection, *amqp.Channel, *amqp.Queue, <-chan amqp.Delivery, error)
}

var _ RabbitMQQueueMessageListener = (*DefaultRabbitMQQueueMessageListener)(nil)

type RabbitMQQueueMessageListener interface {
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
