package messaging

import (
	"fmt"
	"github.com/avast/retry-go/v4"
	"github.com/guidomantilla/go-feather-lib/pkg/common/log"
	amqp "github.com/rabbitmq/amqp091-go"
	"time"
)

type DefaultRabbitMQConnection struct {
	url                      string
	server                   string
	connection               *amqp.Connection
	notifyOnClosedConnection chan *amqp.Error
}

func NewDefaultRabbitMQConnection(messagingContext MessagingContext) *DefaultRabbitMQConnection {

	if messagingContext == nil {
		log.Fatal("starting up - error setting up rabbitMQConnection: messagingContext is nil")
	}

	return &DefaultRabbitMQConnection{
		url:                      messagingContext.GetUrl(),
		server:                   messagingContext.GetServer(),
		notifyOnClosedConnection: make(chan *amqp.Error),
	}
}

func (connection *DefaultRabbitMQConnection) Connect() (*amqp.Connection, error) {

	if connection.connection == nil {

		err := retry.Do(connection.connect, retry.Attempts(5),
			retry.OnRetry(func(n uint, err error) {
				log.Info("rabbitmq connection - failed to connect")
				log.Info(fmt.Sprintf("rabbitmq connection - trying reconnection to %s", connection.server))
			}),
		)

		if err != nil {
			log.Error(fmt.Sprintf("connection - failed connection to %s: %s", connection.server, err.Error()))
			return nil, err
		}

		go connection.reconnect()
	}

	return connection.connection, nil
}

func (connection *DefaultRabbitMQConnection) connect() error {

	var err error
	if connection.connection, err = amqp.Dial(connection.url); err != nil {
		log.Error(err.Error())
		return err
	}

	connection.connection.NotifyClose(connection.notifyOnClosedConnection)
	log.Debug(fmt.Sprintf("rabbitmq connection - connected to %s", connection.server))

	return nil
}

func (connection *DefaultRabbitMQConnection) reconnect() {

	for {
		var ok bool
		var reason *amqp.Error
		if reason, ok = <-connection.notifyOnClosedConnection; !ok {
			log.Info("connection - connection closed")
			break
		}
		log.Info(fmt.Sprintf("rabbitmq connection - connection closed unexpectedly: %s", reason.Reason))

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
