package rabbitmq

import (
	"context"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	samqp "github.com/rabbitmq/rabbitmq-stream-go-client/pkg/amqp"
	"github.com/rabbitmq/rabbitmq-stream-go-client/pkg/stream"
)

var (
	_ Context                         = (*context_)(nil)
	_ Connection[*amqp.Connection]    = (*connection[*amqp.Connection])(nil)
	_ Connection[*stream.Environment] = (*connection[*stream.Environment])(nil)
	_ Consumer                        = (*AmqpConsumer)(nil)
	_ Consumer                        = (*StreamsConsumer)(nil)
	_ Listener[*amqp.Delivery]        = (*AmqpListener)(nil)
	_ Listener[*samqp.Message]        = (*StreamsListener)(nil)

	_ Context                         = (*MockContext)(nil)
	_ Connection[*amqp.Connection]    = (*MockConnection[*amqp.Connection])(nil)
	_ Connection[*stream.Environment] = (*MockConnection[*stream.Environment])(nil)
	_ Listener[*amqp.Delivery]        = (*MockListener[*amqp.Delivery])(nil)
	_ Listener[*samqp.Message]        = (*MockListener[*samqp.Message])(nil)
)

const (
	Delay = 2 * time.Second
)

type Context interface {
	Url() string
	Server() string
	Set(property string, value string)
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
	Set(property string, value any)
}

type PublishingTypes interface {
	*amqp.Publishing | *samqp.AMQP10
}

type Producer[T PublishingTypes] interface {
	Context() Context
	Produce(ctx context.Context, message T) error
	Close()
}
