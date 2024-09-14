package messaging

import (
	"fmt"
	"time"

	retry "github.com/avast/retry-go/v4"
	amqp "github.com/rabbitmq/amqp091-go"

	"github.com/guidomantilla/go-feather-lib/pkg/common/log"
)

type DefaultRabbitMQConnection struct {
	url                      string
	server                   string
	connection               *amqp.Connection
	notifyOnClosedConnection chan *amqp.Error
	channel                  *amqp.Channel
	notifyOnClosedChannel    chan *amqp.Error
}

func NewDefaultRabbitMQConnection(messagingContext MessagingContext) *DefaultRabbitMQConnection {

	if messagingContext == nil {
		log.Fatal("starting up - error setting up rabbitMQConnection: messagingContext is nil")
	}

	return &DefaultRabbitMQConnection{
		url:                      messagingContext.GetUrl(),
		server:                   messagingContext.GetServer(),
		notifyOnClosedConnection: make(chan *amqp.Error),
		notifyOnClosedChannel:    make(chan *amqp.Error),
	}
}

func (connection *DefaultRabbitMQConnection) Close() {

	if connection.channel != nil && !connection.channel.IsClosed() {
		if err := connection.channel.Close(); err != nil {
			log.Error(fmt.Sprintf("rabbitmq connection - failed to close channel to %s: %s", connection.server, err.Error()))
		}
	}

	if connection.connection != nil && !connection.connection.IsClosed() {
		if err := connection.connection.Close(); err != nil {
			log.Error(fmt.Sprintf("rabbitmq connection - failed to close connection to %s: %s", connection.server, err.Error()))
		}
	}
}

func (connection *DefaultRabbitMQConnection) Connect() (*amqp.Connection, error) {

	if connection.connection != nil {
		return connection.connection, nil
	}

	err := retry.Do(connection.connect, retry.Attempts(5),
		retry.OnRetry(func(n uint, err error) {
			log.Warn(fmt.Sprintf("rabbitmq connection - failed to connect: %s", err.Error()))
			log.Debug(fmt.Sprintf("rabbitmq connection - trying reconnection to %s", connection.server))
		}),
	)

	if err != nil {
		log.Error(fmt.Sprintf("rabbitmq connection - failed connection to %s", connection.server))
		return nil, err
	}

	go connection.reconnect()

	return connection.connection, nil
}

func (connection *DefaultRabbitMQConnection) connect() error {

	var err error
	if connection.connection, err = amqp.Dial(connection.url); err != nil {
		return err
	}

	if connection.channel, err = connection.connection.Channel(); err != nil {
		return err
	}

	connection.notifyOnClosedConnection = connection.connection.NotifyClose(make(chan *amqp.Error))
	connection.notifyOnClosedChannel = connection.channel.NotifyClose(make(chan *amqp.Error))
	log.Debug(fmt.Sprintf("rabbitmq connection - connected to %s", connection.server))

	return nil
}

func (connection *DefaultRabbitMQConnection) reconnect() {

	checkClosedNotificationChannel := func() (*amqp.Error, bool) {
		select {
		case reason, ok := <-connection.notifyOnClosedConnection:
			return reason, ok
		case reason, ok := <-connection.notifyOnClosedChannel:
			return reason, ok
		}
	}

	for {
		var ok bool
		var reason *amqp.Error
		if reason, ok = checkClosedNotificationChannel(); !ok {
			log.Info(fmt.Sprintf("rabbitmq connection - connection closed: %s", connection.server))
			continue
		}
		log.Info(fmt.Sprintf("rabbitmq connection - connection closed unexpectedly: %s", reason.Reason))
		connection.Close()

		for {
			time.Sleep(time.Duration(1) * time.Second)
			if err := connection.connect(); err != nil {
				log.Error(fmt.Sprintf("connection - failed reconnection to %s: %s", connection.server, err.Error()))
				continue
			}

			log.Info(fmt.Sprintf("rabbitmq connection - reconnected to %s", connection.server))
			break
		}
	}
}
