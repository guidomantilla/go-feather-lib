package amqp

import (
	"context"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
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

type ConnectionDialer func(url string) (*amqp.Connection, error)

type Connection interface {
	Context() Context
	Connect(ctx context.Context) (*amqp.Connection, error)
	Close(ctx context.Context)
}

type Listener interface {
	OnMessage(ctx context.Context, message *amqp.Delivery) error
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
	Produce(ctx context.Context, message *amqp.Publishing) error
	Close(ctx context.Context)
	Set(property string, value any)
}
