package amqp

import (
	"context"
	"fmt"
	"sync"
	"time"

	retry "github.com/avast/retry-go/v4"
	amqp "github.com/rabbitmq/amqp091-go"

	"github.com/guidomantilla/go-feather-lib/pkg/common/assert"
	"github.com/guidomantilla/go-feather-lib/pkg/common/log"
)

type connection struct {
	context          Context
	connectionDialer ConnectionDialer
	connection       *amqp.Connection
	mu               sync.RWMutex
}

func NewConnection(context Context, connectionDialer ConnectionDialer) *connection {
	assert.NotNil(context, "starting up - error setting up rabbitmq connection: context is nil")
	assert.NotNil(connectionDialer, "starting up - error setting up rabbitmq connection: connectionDialer is nil")

	connection := &connection{
		context:          context,
		connectionDialer: connectionDialer,
	}

	return connection
}

func (connection *connection) Connect(ctx context.Context) (*amqp.Connection, error) {

	connection.mu.Lock()
	defer connection.mu.Unlock()

	if connection.connection != nil && !connection.connection.IsClosed() {
		log.Debug(ctx, fmt.Sprintf("rabbitmq connection - already connected to %s", connection.context.Server()))
		return connection.connection, nil
	}

	err := retry.Do(connection.connect, retry.Attempts(5), retry.Delay(Delay),
		retry.LastErrorOnly(true), retry.OnRetry(func(n uint, err error) {
			log.Warn(ctx, fmt.Sprintf("rabbitmq connection - failed to connect: %s", err.Error()))
		}),
	)

	if err != nil {
		log.Error(ctx, fmt.Sprintf("rabbitmq connection - failed connection to %s", connection.context.Server()))
		return nil, err
	}

	log.Info(ctx, fmt.Sprintf("rabbitmq connection - connected to %s", connection.context.Server()))

	return connection.connection, nil
}

func (connection *connection) connect() error {

	var err error
	if connection.connection, err = connection.connectionDialer(connection.context.Url()); err != nil {
		return err
	}

	return nil
}

func (connection *connection) Close(ctx context.Context) {
	time.Sleep(Delay)

	if connection.connection != nil && !connection.connection.IsClosed() {
		log.Debug(ctx, "rabbitmq connection - closing connection")
		if err := connection.connection.Close(); err != nil {
			log.Error(ctx, fmt.Sprintf("rabbitmq connection - failed to close connection to %s: %s", connection.context.Server(), err.Error()))
		}
	}
	connection.connection = nil
	log.Debug(ctx, fmt.Sprintf("rabbitmq connection - closed connection to %s", connection.context.Server()))
}

func (connection *connection) Context() Context {
	return connection.context
}
