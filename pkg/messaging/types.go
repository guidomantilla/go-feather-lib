package messaging

import (
	"context"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	samqp "github.com/rabbitmq/rabbitmq-stream-go-client/pkg/amqp"
	"github.com/rabbitmq/rabbitmq-stream-go-client/pkg/stream"
)

const (
	Delay = 2 * time.Second
)

var (
	_ MessagingContext                         = (*DefaultMessagingContext)(nil)
	_ MessagingConnection[*amqp.Connection]    = (*RabbitMQConnection[*amqp.Connection])(nil)
	_ MessagingConnection[*stream.Environment] = (*RabbitMQConnection[*stream.Environment])(nil)
	_ MessagingListener[*amqp.Delivery]        = (*RabbitMQListener)(nil)
	_ MessagingListener[*samqp.Message]        = (*RabbitMQStreamsListener)(nil)
	_ MessagingConsumer                        = (*RabbitMQConsumer)(nil)
	_ MessagingConsumer                        = (*RabbitMQStreamsConsumer)(nil)
	_ MessagingContext                         = (*MockMessagingContext)(nil)
	_ MessagingConnection[*amqp.Connection]    = (*MockMessagingConnection[*amqp.Connection])(nil)
	_ MessagingConnection[*stream.Environment] = (*MockMessagingConnection[*stream.Environment])(nil)
	_ MessagingListener[*amqp.Delivery]        = (*MockMessagingListener[*amqp.Delivery])(nil)
	_ MessagingListener[*samqp.Message]        = (*MockMessagingListener[*samqp.Message])(nil)
)

type MessagingContext interface {
	Url() string
	Server() string
}

type MessagingConnectionTypes interface {
	*amqp.Connection | *stream.Environment
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
	*amqp.Delivery | *samqp.Message
}

type MessagingListener[T MessagingListenerTypes] interface {
	OnMessage(ctx context.Context, message T) error
}

type MessagingEvent = chan string

type MessagingConsumer interface {
	MessagingContext() MessagingContext
	Consume(ctx context.Context) (MessagingEvent, error)
	Close()
}

type MessagingPublishingTypes interface {
	*amqp.Publishing | *samqp.AMQP10
}

type MessagingProducer[T MessagingPublishingTypes] interface {
	MessagingContext() MessagingContext
	Produce(ctx context.Context, message T) error
	Close()
}
