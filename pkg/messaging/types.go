package messaging

import (
	"time"

	nats "github.com/nats-io/nats.go"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/rabbitmq/rabbitmq-stream-go-client/pkg/stream"
)

const (
	makeConnectionDelay = 2 * time.Second
)

var (
	_ RabbitMQContext                         = (*DefaultRabbitMQContext)(nil)
	_ RabbitMQConnection[*amqp.Connection]    = (*DefaultRabbitMQConnection)(nil)
	_ RabbitMQConnection[*stream.Environment] = (*DefaultRabbitMQStreamsConnection)(nil)
	_ RabbitMQChannel                         = (*DefaultRabbitMQChannel)(nil)
	_ RabbitMQQueue                           = (*DefaultRabbitMQQueue)(nil)
	_ RabbitMQMessageListener                 = (*DefaultRabbitMQMessageListener)(nil)
	_ NatsSubjectConnection                   = (*DefaultNatsSubjectConnection)(nil)
	_ NatsMessageListener                     = (*DefaultNatsMessageListener)(nil)
	_ RabbitMQContext                         = (*MockRabbitMQContext)(nil)
	_ RabbitMQConnection[*amqp.Connection]    = (*MockRabbitMQConnection[*amqp.Connection])(nil)
	_ RabbitMQConnection[*stream.Environment] = (*MockRabbitMQConnection[*stream.Environment])(nil)
	_ RabbitMQChannel                         = (*MockRabbitMQChannel)(nil)
	_ RabbitMQQueue                           = (*MockRabbitMQQueue)(nil)
	_ RabbitMQMessageListener                 = (*MockRabbitMQMessageListener)(nil)
	_ NatsSubjectConnection                   = (*MockNatsSubjectConnection)(nil)
	_ NatsMessageListener                     = (*MockNatsMessageListener)(nil)
)

// Generic

type RabbitMQContext interface {
	Url() string
	Server() string
	VHost() string
}

type RabbitMQConnectionTypes interface {
	*amqp.Connection | *stream.Environment
}

type RabbitMQConnection[T RabbitMQConnectionTypes] interface {
	RabbitMQContext() RabbitMQContext
	Connect() (T, error)
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

type RabbitMQMessageListener interface {
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
