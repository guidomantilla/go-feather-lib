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
	_ Context                         = (*context_)(nil)
	_ Connection[*amqp.Connection]    = (*RabbitMQConnection[*amqp.Connection])(nil)
	_ Connection[*stream.Environment] = (*RabbitMQConnection[*stream.Environment])(nil)
	_ Listener[*amqp.Delivery]        = (*RabbitMQListener)(nil)
	_ Listener[*samqp.Message]        = (*RabbitMQStreamsListener)(nil)
	_ Consumer                        = (*RabbitMQConsumer)(nil)
	_ Consumer                        = (*RabbitMQStreamsConsumer)(nil)
	_ Context                         = (*MockContext)(nil)
	_ Connection[*amqp.Connection]    = (*MockConnection[*amqp.Connection])(nil)
	_ Connection[*stream.Environment] = (*MockConnection[*stream.Environment])(nil)
	_ Listener[*amqp.Delivery]        = (*MockListener[*amqp.Delivery])(nil)
	_ Listener[*samqp.Message]        = (*MockListener[*samqp.Message])(nil)
)

type Context interface {
	Url() string
	Server() string
	set(property string, value string)
}

type ConnectionTypes interface {
	*amqp.Connection | *stream.Environment
	IsClosed() bool
	Close() error
}

type ConnectionDialer[T ConnectionTypes] func(url string) (T, error)

type Connection[T ConnectionTypes] interface {
	Context() Context
	Connect() (T, error)
	Close()
}

type ListenerTypes interface {
	*amqp.Delivery | *samqp.Message
}

type Listener[T ListenerTypes] interface {
	OnMessage(ctx context.Context, message T) error
}

type Event = chan string

type Consumer interface {
	Context() Context
	Consume(ctx context.Context) (Event, error)
	Close()
	set(property string, value any)
}

type PublishingTypes interface {
	*amqp.Publishing | *samqp.AMQP10
}

type Producer[T PublishingTypes] interface {
	Context() Context
	Produce(ctx context.Context, message T) error
	Close()
}
