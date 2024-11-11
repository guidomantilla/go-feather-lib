package streams

import (
	"context"
	"time"

	samqp "github.com/rabbitmq/rabbitmq-stream-go-client/pkg/amqp"
	"github.com/rabbitmq/rabbitmq-stream-go-client/pkg/stream"
)

var (
	_ Context    = (*context_)(nil)
	_ Connection = (*connection)(nil)
	_ Consumer   = (*consumer)(nil)
	_ Listener   = (*listener)(nil)

	_ Context    = (*MockContext)(nil)
	_ Connection = (*MockConnection)(nil)
	_ Listener   = (*MockListener)(nil)
)

const (
	Delay = 2 * time.Second
)

type Context interface {
	Url() string
	Server() string
	Set(property string, value string)
}

type ConnectionDialer func(url string) (*stream.Environment, error)

type Connection interface {
	Context() Context
	Connect(ctx context.Context) (*stream.Environment, error)
	Close(ctx context.Context)
}

type Listener interface {
	OnMessage(ctx context.Context, message *samqp.Message) error
}

type Event = chan string

type Consumer interface {
	Context() Context
	Consume(ctx context.Context) (Event, error)
	Close(ctx context.Context)
	Set(property string, value any)
}

type Producer interface {
	Context() Context
	Produce(ctx context.Context, message *samqp.AMQP10) error
	Close(ctx context.Context)
	Set(property string, value any)
}
