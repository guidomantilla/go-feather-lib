package messaging

import (
	"fmt"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"

	"github.com/guidomantilla/go-feather-lib/pkg/common/log"
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

	go channel.reconnect()

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

	channel.notifyOnClosedChannel = channel.channel.NotifyClose(make(chan *amqp.Error))
	log.Debug(fmt.Sprintf("rabbitmq channel - connected to channel"))

	return nil
}

func (channel *DefaultRabbitMQChannel) reconnect() {

	if !channel.rabbitMQConnection.RabbitMQContext().FailOver() {
		return
	}

	for {
		var ok bool
		var reason *amqp.Error
		if reason, ok = <-channel.notifyOnClosedChannel; !ok {
			break
		}
		log.Debug(fmt.Sprintf("rabbitmq channel - channel closed unexpectedly: %s", reason.Reason))

		<-channel.RabbitMQContext().NotifyOnFaiOverConnection()
		time.Sleep(makeConnectionDelay)
		channel.Close()

		log.Debug(fmt.Sprintf("rabbitmq channel - trying reconnection to channel"))

		for {
			time.Sleep(makeConnectionDelay)
			if err := channel.connect(); err != nil {
				log.Error(fmt.Sprintf("rabbitmq channel - failed reconnection to channel: %s", err.Error()))
				continue
			}
			log.Info(fmt.Sprintf("rabbitmq channel - reconnected to channel"))
			break
		}
	}
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
