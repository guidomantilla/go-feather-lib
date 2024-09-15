package messaging

import (
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"

	"github.com/guidomantilla/go-feather-lib/pkg/common/log"
)

type DefaultRabbitMQChannel struct {
	rabbitMQConnection    RabbitMQConnection[*amqp.Connection]
	channel               *amqp.Channel
	notifyOnClosedChannel chan *amqp.Error
}

func NewDefaultRabbitMQChannel(rabbitMQConnection RabbitMQConnection[*amqp.Connection]) *DefaultRabbitMQChannel {

	if rabbitMQConnection == nil {
		log.Fatal("starting up - error setting up rabbitMQChannel: rabbitMQConnection is nil")
	}

	return &DefaultRabbitMQChannel{
		rabbitMQConnection:    rabbitMQConnection,
		notifyOnClosedChannel: make(chan *amqp.Error),
	}
}

func (channel *DefaultRabbitMQChannel) Connect() (*amqp.Channel, error) {

	var err error
	var connection *amqp.Connection
	if connection, err = channel.rabbitMQConnection.Connect(); err != nil {
		log.Debug(fmt.Sprintf("rabbitmq channel - failed connection to channel: %s", err.Error()))
		return nil, err
	}

	if !(channel.channel != nil && !channel.channel.IsClosed()) {
		if channel.channel, err = connection.Channel(); err != nil {
			log.Debug(fmt.Sprintf("rabbitmq channel - failed connection channel: %s", err.Error()))
			return nil, err
		}
	}

	log.Debug("rabbitmq channel - connected to channel")

	return channel.channel, nil
}

func (channel *DefaultRabbitMQChannel) Close() {

	if channel.channel != nil && !channel.channel.IsClosed() {
		log.Debug("rabbitmq channel - closing connection")
		if err := channel.channel.Close(); err != nil {
			log.Error(fmt.Sprintf("rabbitmq channel - failed to close connection to channel: %s", err.Error()))
		}
	}
	channel.channel = nil
	log.Debug("rabbitmq channel - closed connection to channel")
}

func (channel *DefaultRabbitMQChannel) MessagingContext() MessagingContext {
	return channel.rabbitMQConnection.MessagingContext()
}
