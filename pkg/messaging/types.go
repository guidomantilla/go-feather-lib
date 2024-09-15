package messaging

import (
	"time"

	nats "github.com/nats-io/nats.go"
	amqp "github.com/rabbitmq/amqp091-go"
)

const (
	makeConnectionDelay = 2 * time.Second
)

var (
	_ RabbitMQContext              = (*DefaultRabbitMQContext)(nil)
	_ RabbitMQConnection           = (*DefaultRabbitMQConnection)(nil)
	_ RabbitMQChannel              = (*DefaultRabbitMQChannel)(nil)
	_ RabbitMQQueue                = (*DefaultRabbitMQQueue)(nil)
	_ RabbitMQQueueMessageListener = (*DefaultRabbitMQQueueMessageListener)(nil)
	_ NatsSubjectConnection        = (*DefaultNatsSubjectConnection)(nil)
	_ NatsMessageListener          = (*DefaultNatsMessageListener)(nil)
	_ RabbitMQContext              = (*MockRabbitMQContext)(nil)
	_ RabbitMQConnection           = (*MockRabbitMQConnection)(nil)
	_ RabbitMQChannel              = (*MockRabbitMQChannel)(nil)
	_ RabbitMQQueue                = (*MockRabbitMQQueue)(nil)
	_ RabbitMQQueueMessageListener = (*MockRabbitMQQueueMessageListener)(nil)
	_ NatsSubjectConnection        = (*MockNatsSubjectConnection)(nil)
	_ NatsMessageListener          = (*MockNatsMessageListener)(nil)
)

type RabbitMQContext interface {
	Url() string
	Server() string
}

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
	Connect() (*amqp.Channel, error)
	Close()
	Name() string
	Consumer() string
}

type RabbitMQQueueMessageListener interface {
	OnMessage(message *amqp.Delivery) error
}

//

type NatsSubjectConnection interface {
	Start()
	Close()
	Connect() (*nats.Conn, *nats.Subscription, chan *nats.Msg, error)
}

type NatsMessageListener interface {
	OnMessage(message *nats.Msg) error
}
