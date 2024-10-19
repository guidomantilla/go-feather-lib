package messaging

import (
	"crypto/tls"
	"fmt"
	"sync"
	"time"

	retry "github.com/avast/retry-go/v4"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/rabbitmq/rabbitmq-stream-go-client/pkg/stream"

	"github.com/guidomantilla/go-feather-lib/pkg/common/log"
)

type RabbitMQConnectionOption[T ConnectionTypes] func(rabbitMQConnection *RabbitMQConnection[T])

func WithRabbitMQDialer() RabbitMQConnectionOption[*amqp.Connection] {
	return func(rabbitMQConnection *RabbitMQConnection[*amqp.Connection]) {
		rabbitMQConnection.connectionDialer = func(url string) (*amqp.Connection, error) {
			return amqp.Dial(url)
		}
	}
}

func WithRabbitMQDialerTLS(amqps *tls.Config) RabbitMQConnectionOption[*amqp.Connection] {
	return func(rabbitMQConnection *RabbitMQConnection[*amqp.Connection]) {
		rabbitMQConnection.connectionDialer = func(url string) (*amqp.Connection, error) {
			return amqp.DialTLS(url, amqps)
		}
	}
}

func WithRabbitMQStreamsDialer() RabbitMQConnectionOption[*stream.Environment] {
	return func(rabbitMQConnection *RabbitMQConnection[*stream.Environment]) {
		rabbitMQConnection.connectionDialer = func(url string) (*stream.Environment, error) {
			return stream.NewEnvironment(stream.NewEnvironmentOptions().SetUri(url))
		}
	}
}

func WithRabbitMQStreamsDialerTLS(streams *tls.Config) RabbitMQConnectionOption[*stream.Environment] {
	return func(rabbitMQConnection *RabbitMQConnection[*stream.Environment]) {
		rabbitMQConnection.connectionDialer = func(url string) (*stream.Environment, error) {
			return stream.NewEnvironment(stream.NewEnvironmentOptions().SetUri(url).SetTLSConfig(streams))
		}
	}
}

func WithMessagingConnectionDialer[T ConnectionTypes](dialer ConnectionDialer[T]) RabbitMQConnectionOption[T] {
	return func(rabbitMQConnection *RabbitMQConnection[T]) {
		rabbitMQConnection.connectionDialer = dialer
	}
}

type RabbitMQConnection[T ConnectionTypes] struct {
	context          Context
	connectionDialer ConnectionDialer[T]
	connection       T
	mu               sync.RWMutex
}

func NewRabbitMQConnection[T ConnectionTypes](context Context, options ...RabbitMQConnectionOption[T]) *RabbitMQConnection[T] {

	if context == nil {
		log.Fatal("starting up - error setting up rabbitmq connection: context is nil")
	}

	if len(options) == 0 {
		log.Fatal("starting up - error setting up rabbitmq connection: options is empty")
	}

	connection := &RabbitMQConnection[T]{
		context: context,
	}

	for _, option := range options {
		option(connection)
	}

	return connection
}

func (connection *RabbitMQConnection[T]) Connect() (T, error) {

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

func (connection *RabbitMQConnection[T]) connect() error {

	var err error
	if connection.connection, err = connection.connectionDialer(connection.context.Url()); err != nil {
		return err
	}

	log.Info(fmt.Sprintf("rabbitmq connection - connected to %s", connection.context.Server()))

	return nil
}

func (connection *RabbitMQConnection[T]) Close() {
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

func (connection *RabbitMQConnection[T]) Context() Context {
	return connection.context
}
