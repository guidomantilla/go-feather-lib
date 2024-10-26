package rabbitmq

import (
	"fmt"
	"sync"
	"time"

	retry "github.com/avast/retry-go/v4"

	"github.com/guidomantilla/go-feather-lib/pkg/common/assert"
	"github.com/guidomantilla/go-feather-lib/pkg/common/log"
)

type connection[T ConnectionTypes] struct {
	context          Context
	connectionDialer ConnectionDialer[T]
	connection       T
	mu               sync.RWMutex
}

func NewConnection[T ConnectionTypes](context Context, connectionDialer ConnectionDialer[T]) *connection[T] {
	assert.NotNil(context, "starting up - error setting up rabbitmq connection: context is nil")
	assert.NotNil(connectionDialer, "starting up - error setting up rabbitmq connection: connectionDialer is nil")

	connection := &connection[T]{
		context:          context,
		connectionDialer: connectionDialer,
	}

	return connection
}

func (connection *connection[T]) Connect() (T, error) {

	connection.mu.Lock()
	defer connection.mu.Unlock()

	if connection.connection != nil && !connection.connection.IsClosed() {
		log.Debug(fmt.Sprintf("rabbitmq connection - already connected to %s", connection.context.Server()))
		return connection.connection, nil
	}

	err := retry.Do(connection.connect, retry.Attempts(5), retry.Delay(Delay),
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

func (connection *connection[T]) connect() error {

	var err error
	if connection.connection, err = connection.connectionDialer(connection.context.Url()); err != nil {
		return err
	}

	log.Info(fmt.Sprintf("rabbitmq connection - connected to %s", connection.context.Server()))

	return nil
}

func (connection *connection[T]) Close() {
	time.Sleep(Delay)

	if connection.connection != nil && !connection.connection.IsClosed() {
		log.Debug("rabbitmq connection - closing connection")
		if err := connection.connection.Close(); err != nil {
			log.Error(fmt.Sprintf("rabbitmq connection - failed to close connection to %s: %s", connection.context.Server(), err.Error()))
		}
	}
	connection.connection = nil
	log.Debug(fmt.Sprintf("rabbitmq connection - closed connection to %s", connection.context.Server()))
}

func (connection *connection[T]) Context() Context {
	return connection.context
}
