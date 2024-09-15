package messaging

import (
	"fmt"
	"sync"

	retry "github.com/avast/retry-go/v4"
	"github.com/rabbitmq/rabbitmq-stream-go-client/pkg/stream"

	"github.com/guidomantilla/go-feather-lib/pkg/common/log"
)

type DefaultRabbitMQStreamsConnection struct {
	rabbitmqContext RabbitMQContext
	environment     *stream.Environment
	mu              sync.Mutex
}

func NewDefaultRabbitMQStreamsConnection(rabbitmqContext RabbitMQContext) *DefaultRabbitMQStreamsConnection {

	if rabbitmqContext == nil {
		log.Fatal("starting up - error setting up rabbitMQStreamsConnection: rabbitmqContext is nil")
	}

	return &DefaultRabbitMQStreamsConnection{
		rabbitmqContext: rabbitmqContext,
	}
}

func (connection *DefaultRabbitMQStreamsConnection) Connect() (*stream.Environment, error) {

	connection.mu.Lock()
	defer connection.mu.Unlock()

	if connection.environment != nil && !connection.environment.IsClosed() {
		log.Debug(fmt.Sprintf("rabbitmq streams connection - already connected to %s", connection.rabbitmqContext.Server()))
		return connection.environment, nil
	}

	err := retry.Do(connection.connect, retry.Attempts(5), retry.Delay(makeConnectionDelay),
		retry.LastErrorOnly(true), retry.OnRetry(func(n uint, err error) {
			log.Warn(fmt.Sprintf("rabbitmq streams connection - failed to connect: %s", err.Error()))
		}),
	)

	if err != nil {
		log.Error(fmt.Sprintf("rabbitmq streams connection - failed connection to %s", connection.rabbitmqContext.Server()))
		return nil, err
	}

	return connection.environment, nil
}

func (connection *DefaultRabbitMQStreamsConnection) connect() error {

	var err error
	if connection.environment, err = stream.NewEnvironment(stream.NewEnvironmentOptions().SetUri(connection.rabbitmqContext.Url())); err != nil {
		return err
	}

	log.Info(fmt.Sprintf("rabbitmq streams connection - connected to %s", connection.rabbitmqContext.Server()))

	return nil
}

func (connection *DefaultRabbitMQStreamsConnection) Close() {

	if connection.environment != nil && !connection.environment.IsClosed() {
		log.Debug("rabbitmq streams connection - closing connection")
		if err := connection.environment.Close(); err != nil {
			log.Error(fmt.Sprintf("rabbitmq streams connection - failed to close connection to %s: %s", connection.rabbitmqContext.Server(), err.Error()))
		}
	}
	connection.environment = nil
	log.Debug(fmt.Sprintf("rabbitmq streams connection - closed connection to %s", connection.rabbitmqContext.Server()))
}

func (connection *DefaultRabbitMQStreamsConnection) RabbitMQContext() RabbitMQContext {
	return connection.rabbitmqContext
}
