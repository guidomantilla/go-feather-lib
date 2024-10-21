package rabbitmq

import (
	"fmt"
	"sync"
	"time"

	retry "github.com/avast/retry-go/v4"

	"github.com/guidomantilla/go-feather-lib/pkg/common/assert"
	"github.com/guidomantilla/go-feather-lib/pkg/common/log"
	"github.com/guidomantilla/go-feather-lib/pkg/messaging"
)

type Connection[T messaging.ConnectionTypes] struct {
	context          messaging.Context
	connectionDialer messaging.ConnectionDialer[T]
	connection       T
	mu               sync.RWMutex
}

func NewConnection[T messaging.ConnectionTypes](context messaging.Context, connectionDialer messaging.ConnectionDialer[T]) *Connection[T] {
	assert.NotNil(context, "starting up - error setting up rabbitmq connection: context is nil")
	assert.NotNil(connectionDialer, "starting up - error setting up rabbitmq connection: connectionDialer is nil")

	connection := &Connection[T]{
		context:          context,
		connectionDialer: connectionDialer,
	}

	return connection
}

func (connection *Connection[T]) Connect() (T, error) {

	connection.mu.Lock()
	defer connection.mu.Unlock()

	if connection.connection != nil && !connection.connection.IsClosed() {
		log.Debug(fmt.Sprintf("rabbitmq connection - already connected to %s", connection.context.Server()))
		return connection.connection, nil
	}

	err := retry.Do(connection.connect, retry.Attempts(5), retry.Delay(messaging.Delay),
		retry.LastErrorOnly(true), retry.OnRetry(func(n uint, err error) {
			log.Warn(fmt.Sprintf("rabbitmq connection - failed to connect: %s", err.Error()))
		}),
	)

	if err != nil {
		log.Error(fmt.Sprintf("rabbitmq connection - failed connection to %s", connection.context.Server()))
		return nil, err
	}

	return connection.connection, nil
}

func (connection *Connection[T]) connect() error {

	var err error
	if connection.connection, err = connection.connectionDialer(connection.context.Url()); err != nil {
		return err
	}

	log.Info(fmt.Sprintf("rabbitmq connection - connected to %s", connection.context.Server()))

	return nil
}

func (connection *Connection[T]) Close() {
	time.Sleep(messaging.Delay)

	if connection.connection != nil && !connection.connection.IsClosed() {
		log.Debug("rabbitmq connection - closing connection")
		if err := connection.connection.Close(); err != nil {
			log.Error(fmt.Sprintf("rabbitmq connection - failed to close connection to %s: %s", connection.context.Server(), err.Error()))
		}
	}
	connection.connection = nil
	log.Debug(fmt.Sprintf("rabbitmq connection - closed connection to %s", connection.context.Server()))
}

func (connection *Connection[T]) Context() messaging.Context {
	return connection.context
}
