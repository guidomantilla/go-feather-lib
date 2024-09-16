package messaging

import (
	"context"
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
	_ MessagingContext                         = (*DefaultMessagingContext)(nil)
	_ MessagingConnection[*amqp.Connection]    = (*RabbitMQConnection[*amqp.Connection])(nil)
	_ MessagingConnection[*stream.Environment] = (*RabbitMQConnection[*stream.Environment])(nil)
	_ MessagingListener[*amqp.Delivery]        = (*RabbitMQListener)(nil)
	_ MessagingListener[*samqp.Message]        = (*RabbitMQStreamsListener)(nil)
	_ MessagingListener[*nats.Msg]             = (*NatsListener)(nil)
	_ MessagingConsumer                        = (*RabbitMQConsumer)(nil)
	_ MessagingConsumer                        = (*RabbitMQStreamsConsumer)(nil)
	_ MessagingContext                         = (*MockMessagingContext)(nil)
	_ MessagingConnection[*amqp.Connection]    = (*MockMessagingConnection[*amqp.Connection])(nil)
	_ MessagingConnection[*stream.Environment] = (*MockMessagingConnection[*stream.Environment])(nil)
	_ MessagingListener[*amqp.Delivery]        = (*MockMessagingListener[*amqp.Delivery])(nil)
	_ MessagingListener[*samqp.Message]        = (*MockMessagingListener[*samqp.Message])(nil)
	_ MessagingListener[*nats.Msg]             = (*MockMessagingListener[*nats.Msg])(nil)
)

type MessagingContext interface {
	Url() string
	Server() string
}

type MessagingConnectionTypes interface {
	*amqp.Connection | *stream.Environment | *nats.Conn
	IsClosed() bool
	Close() error
}

type MessagingConnectionDialer[T MessagingConnectionTypes] func(url string) (T, error)

type MessagingConnection[T MessagingConnectionTypes] interface {
	MessagingContext() MessagingContext
	Connect() (T, error)
	Close()
}

type MessagingListenerTypes interface {
	*amqp.Delivery | *samqp.Message | *nats.Msg
}

type MessagingPublishingTypes interface {
	//TODO: Add the types for the different messaging systems
}

type MessagingListener[T MessagingListenerTypes] interface {
	OnMessage(message T) error
}

type MessagingTargetTypes interface {
	*amqp.Channel | *stream.Environment
}

type MessagingEvent = chan string

type MessagingConsumer interface {
	MessagingContext() MessagingContext
	Consume(ctx context.Context) (MessagingEvent, error)
	Close()
}

type MessagingProducer[T MessagingPublishingTypes] interface {
	MessagingContext() MessagingContext
	Produce(ctx context.Context, message T) error
	Close()
}
