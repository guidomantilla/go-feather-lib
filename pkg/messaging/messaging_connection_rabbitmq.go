package messaging

import (
	"fmt"
	"sync"

	retry "github.com/avast/retry-go/v4"

	"github.com/guidomantilla/go-feather-lib/pkg/common/log"
)

type RabbitMQConnection[T MessagingConnectionTypes] struct {
	messagingContext          MessagingContext
	messagingConnectionDialer MessagingConnectionDialer[T]
	connection                T
	mu                        sync.Mutex
}

func NewRabbitMQConnection[T MessagingConnectionTypes](messagingContext MessagingContext, messagingConnectionDialer MessagingConnectionDialer[T]) *RabbitMQConnection[T] {

	if messagingContext == nil {
		log.Fatal("starting up - error setting up rabbitMQConnection: messagingContext is nil")
	}

	if messagingConnectionDialer == nil {
		log.Fatal("starting up - error setting up rabbitMQConnection: messagingConnectionDialer is nil")
	}

	return &RabbitMQConnection[T]{
		messagingContext:          messagingContext,
		messagingConnectionDialer: messagingConnectionDialer,
	}
}

func (connection *RabbitMQConnection[T]) Connect() (T, error) {

	connection.mu.Lock()
	defer connection.mu.Unlock()

	if connection.connection != nil && !connection.connection.IsClosed() {
		log.Debug(fmt.Sprintf("rabbitmq connection - already connected to %s", connection.messagingContext.Server()))
		return connection.connection, nil
	}

	err := retry.Do(connection.connect, retry.Attempts(5), retry.Delay(makeConnectionDelay),
		retry.LastErrorOnly(true), retry.OnRetry(func(n uint, err error) {
			log.Warn(fmt.Sprintf("rabbitmq connection - failed to connect: %s", err.Error()))
		}),
	)

	if err != nil {
		log.Error(fmt.Sprintf("rabbitmq connection - failed connection to %s", connection.messagingContext.Server()))
		return nil, err
	}

	return connection.connection, nil
}

func (connection *RabbitMQConnection[T]) connect() error {

	var err error
	if connection.connection, err = connection.messagingConnectionDialer(connection.messagingContext.Url()); err != nil {
		return err
	}

	log.Info(fmt.Sprintf("rabbitmq connection - connected to %s", connection.messagingContext.Server()))

	return nil
}

func (connection *RabbitMQConnection[T]) Close() {

	if connection.connection != nil && !connection.connection.IsClosed() {
		log.Debug("rabbitmq connection - closing connection")
		if err := connection.connection.Close(); err != nil {
			log.Error(fmt.Sprintf("rabbitmq connection - failed to close connection to %s: %s", connection.messagingContext.Server(), err.Error()))
		}
	}
	connection.connection = nil
	log.Debug(fmt.Sprintf("rabbitmq connection - closed connection to %s", connection.messagingContext.Server()))
}

func (connection *RabbitMQConnection[T]) MessagingContext() MessagingContext {
	return connection.messagingContext
}
