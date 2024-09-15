package messaging

import (
	"fmt"
	"github.com/guidomantilla/go-feather-lib/pkg/common/log"
	amqp "github.com/rabbitmq/amqp091-go"
)

type DefaultRabbitMQChannel struct {
	rabbitMQConnection    RabbitMQConnection
	channel               *amqp.Channel
	notifyOnClosedChannel chan *amqp.Error
}

func NewDefaultRabbitMQChannel(rabbitMQConnection RabbitMQConnection) *DefaultRabbitMQChannel {

	if rabbitMQConnection == nil {
		log.Fatal("starting up - error setting up rabbitMQChannel: rabbitMQConnection is nil")
	}

	return &DefaultRabbitMQChannel{
		rabbitMQConnection:    rabbitMQConnection,
		notifyOnClosedChannel: make(chan *amqp.Error),
	}
}

func (channel *DefaultRabbitMQChannel) Connect() (*amqp.Channel, error) {

	if channel.channel != nil && !channel.channel.IsClosed() {
		log.Debug(fmt.Sprintf("rabbitmq channel - already connected to channel"))
		return channel.channel, nil
	}
	/*
		err := retry.Do(channel.connect, retry.Attempts(5),
			retry.OnRetry(func(n uint, err error) {
				log.Warn(fmt.Sprintf("rabbitmq channel - failed to connect: %s", err.Error()))
				log.Debug(fmt.Sprintf("rabbitmq channel - trying reconnection to channel"))
			}),
		)
	*/
	if err := channel.connect(); err != nil {
		log.Error(fmt.Sprintf("rabbitmq channel - failed connection to channel"))
		return nil, err
	}

	return channel.channel, nil
}

func (channel *DefaultRabbitMQChannel) connect() error {

	var err error
	var connection *amqp.Connection
	if connection, err = channel.rabbitMQConnection.Connect(); err != nil {
		return err
	}

	if channel.channel, err = connection.Channel(); err != nil {
		return err
	}

	//channel.notifyOnClosedChannel = channel.channel.NotifyClose(make(chan *amqp.Error))
	log.Debug(fmt.Sprintf("rabbitmq channel - connected to channel"))

	return nil
}

func (channel *DefaultRabbitMQChannel) Close() {

	if channel.channel != nil && !channel.channel.IsClosed() {
		log.Debug("rabbitmq channel - closing connection")
		if err := channel.channel.Close(); err != nil {
			log.Error(fmt.Sprintf("rabbitmq channel - failed to close connection to channel: %s", err.Error()))
		}
	}
	channel.channel = nil
	log.Debug(fmt.Sprintf("rabbitmq channel - closed connection to channel"))
}

func (channel *DefaultRabbitMQChannel) RabbitMQContext() RabbitMQContext {
	return channel.rabbitMQConnection.RabbitMQContext()
}
