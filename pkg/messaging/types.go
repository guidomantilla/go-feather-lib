package messaging

import (
	"time"

	nats "github.com/nats-io/nats.go"
	amqp "github.com/rabbitmq/amqp091-go"
	samqp "github.com/rabbitmq/rabbitmq-stream-go-client/pkg/amqp"
	"github.com/rabbitmq/rabbitmq-stream-go-client/pkg/stream"
)

const (
	makeConnectionDelay = 2 * time.Second
)

var (
	_ MessagingContext                        = (*DefaultRabbitMQContext)(nil)
	_ RabbitMQConnection[*amqp.Connection]    = (*DefaultRabbitMQConnection)(nil)
	_ RabbitMQConnection[*stream.Environment] = (*DefaultRabbitMQStreamsConnection)(nil)
	_ RabbitMQChannel                         = (*DefaultRabbitMQChannel)(nil)
	_ RabbitMQQueue                           = (*DefaultRabbitMQQueue)(nil)
	_ RabbitMQMessageListener[*amqp.Delivery] = (*DefaultRabbitMQListener)(nil)
	_ RabbitMQMessageListener[*samqp.Message] = (*DefaultRabbitMQStreamsListener)(nil)
	_ NatsSubjectConnection                   = (*DefaultNatsSubjectConnection)(nil)
	_ NatsMessageListener                     = (*DefaultNatsMessageListener)(nil)
	_ MessagingContext                        = (*MockMessagingContext)(nil)
	_ RabbitMQConnection[*amqp.Connection]    = (*MockRabbitMQConnection[*amqp.Connection])(nil)
	_ RabbitMQConnection[*stream.Environment] = (*MockRabbitMQConnection[*stream.Environment])(nil)
	_ RabbitMQChannel                         = (*MockRabbitMQChannel)(nil)
	_ RabbitMQQueue                           = (*MockRabbitMQQueue)(nil)
	_ RabbitMQMessageListener[*amqp.Delivery] = (*MockRabbitMQMessageListener[*amqp.Delivery])(nil)
	_ RabbitMQMessageListener[*samqp.Message] = (*MockRabbitMQMessageListener[*samqp.Message])(nil)
	_ NatsSubjectConnection                   = (*MockNatsSubjectConnection)(nil)
	_ NatsMessageListener                     = (*MockNatsMessageListener)(nil)
)

type MessagingContext interface {
	Url() string
	Server() string
}

// RabbitGeneric

type RabbitMQConnectionTypes interface {
	*amqp.Connection | *stream.Environment
}

type RabbitMQConnection[T RabbitMQConnectionTypes] interface {
	MessagingContext() MessagingContext
	Connect() (T, error)
	Close()
}

type RabbitMQMessageListenerTypes interface {
	*amqp.Delivery | *samqp.Message
}

type RabbitMQMessageListener[T RabbitMQMessageListenerTypes] interface {
	OnMessage(message T) error
}

// RabbitMQ Classic

type RabbitMQChannel interface {
	MessagingContext() MessagingContext
	Connect() (*amqp.Channel, error)
	Close()
}

type RabbitMQQueue interface {
	MessagingContext() MessagingContext
	Connect() (*amqp.Channel, error)
	Close()
	Name() string
	Consumer() string
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
