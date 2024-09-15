package messaging

import (
	"fmt"
	"time"

	retry "github.com/avast/retry-go/v4"
	amqp "github.com/rabbitmq/amqp091-go"

	"github.com/guidomantilla/go-feather-lib/pkg/common/log"
)

type DefaultRabbitMQConnection struct {
	rabbitmqContext          RabbitMQContext
	connection               *amqp.Connection
	notifyOnClosedConnection chan *amqp.Error
}

func NewDefaultRabbitMQConnection(rabbitmqContext RabbitMQContext) *DefaultRabbitMQConnection {

	if rabbitmqContext == nil {
		log.Fatal("starting up - error setting up rabbitMQConnection: rabbitmqContext is nil")
	}

	return &DefaultRabbitMQConnection{
		rabbitmqContext:          rabbitmqContext,
		notifyOnClosedConnection: make(chan *amqp.Error),
	}
}

func (connection *DefaultRabbitMQConnection) Connect() (*amqp.Connection, error) {

	if connection.connection != nil && !connection.connection.IsClosed() {
		return connection.connection, nil
	}

	err := retry.Do(connection.connect, retry.Attempts(5), retry.Delay(makeConnectionDelay),
		retry.OnRetry(func(n uint, err error) {
			log.Warn(fmt.Sprintf("rabbitmq connection - failed to connect: %s", err.Error()))
		}),
	)

	if err != nil {
		log.Error(fmt.Sprintf("rabbitmq connection - failed connection to %s", connection.rabbitmqContext.Server()))
		return nil, err
	}

	go connection.reconnect()

	return connection.connection, nil
}

func (connection *DefaultRabbitMQConnection) connect() error {

	var err error
	if connection.connection, err = amqp.Dial(connection.rabbitmqContext.Url()); err != nil {
		return err
	}

	connection.notifyOnClosedConnection = connection.connection.NotifyClose(make(chan *amqp.Error))
	log.Debug(fmt.Sprintf("rabbitmq connection - connected to %s", connection.rabbitmqContext.Server()))

	return nil
}

func (connection *DefaultRabbitMQConnection) reconnect() {

	if !connection.rabbitmqContext.FailOver() {
		return
	}

	for {
		var ok bool
		var reason *amqp.Error
		if reason, ok = <-connection.notifyOnClosedConnection; !ok {
			break
		}
		log.Warn(fmt.Sprintf("rabbitmq connection - connection closed unexpectedly: %s", reason.Reason))

		time.Sleep(makeConnectionDelay)
		connection.Close()

		for {
			time.Sleep(makeConnectionDelay)
			if err := connection.connect(); err != nil {
				log.Error(fmt.Sprintf("rabbitmq connection - failed reconnection to %s: %s", connection.rabbitmqContext.Server(), err.Error()))
				continue
			}

			log.Info(fmt.Sprintf("rabbitmq connection - reconnected to %s", connection.rabbitmqContext.Server()))
			connection.rabbitmqContext.NotifyOnFaiOverConnection() <- fmt.Sprintf("reconnected to %s", connection.rabbitmqContext.Server())
			break
		}
	}
}

func (connection *DefaultRabbitMQConnection) Close() {

	if connection.connection != nil && !connection.connection.IsClosed() {
		log.Debug("rabbitmq connection - closing connection")
		if err := connection.connection.Close(); err != nil {
			log.Error(fmt.Sprintf("rabbitmq connection - failed to close connection to %s: %s", connection.rabbitmqContext.Server(), err.Error()))
		}
	}
	connection.connection = nil
	log.Debug(fmt.Sprintf("rabbitmq connection - closed connection to %s", connection.rabbitmqContext.Server()))
}

func (connection *DefaultRabbitMQConnection) RabbitMQContext() RabbitMQContext {
	return connection.rabbitmqContext
}
