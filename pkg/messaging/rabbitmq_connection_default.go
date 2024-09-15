package messaging

import (
	"fmt"
	"sync"

	retry "github.com/avast/retry-go/v4"
	amqp "github.com/rabbitmq/amqp091-go"

	"github.com/guidomantilla/go-feather-lib/pkg/common/log"
)

type DefaultRabbitMQConnection struct {
	messagingContext MessagingContext
	connection       *amqp.Connection
	mu               sync.Mutex
}

func NewDefaultRabbitMQConnection(messagingContext MessagingContext) *DefaultRabbitMQConnection {

	if messagingContext == nil {
		log.Fatal("starting up - error setting up rabbitMQConnection: messagingContext is nil")
	}

	return &DefaultRabbitMQConnection{
		messagingContext: messagingContext,
	}
}

func (connection *DefaultRabbitMQConnection) Connect() (*amqp.Connection, error) {

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

func (connection *DefaultRabbitMQConnection) connect() error {

	var err error
	if connection.connection, err = amqp.Dial(connection.messagingContext.Url()); err != nil {
		return err
	}

	log.Info(fmt.Sprintf("rabbitmq connection - connected to %s", connection.messagingContext.Server()))

	return nil
}

func (connection *DefaultRabbitMQConnection) Close() {

	if connection.connection != nil && !connection.connection.IsClosed() {
		log.Debug("rabbitmq connection - closing connection")
		if err := connection.connection.Close(); err != nil {
			log.Error(fmt.Sprintf("rabbitmq connection - failed to close connection to %s: %s", connection.messagingContext.Server(), err.Error()))
		}
	}
	connection.connection = nil
	log.Debug(fmt.Sprintf("rabbitmq connection - closed connection to %s", connection.messagingContext.Server()))
}

func (connection *DefaultRabbitMQConnection) MessagingContext() MessagingContext {
	return connection.messagingContext
}
