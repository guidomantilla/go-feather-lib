package messaging

import (
	"fmt"
	"strings"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"

	"github.com/guidomantilla/go-feather-lib/pkg/common/log"
)

type DeliveryChan <-chan amqp.Delivery

type DefaultRabbitMQQueue struct {
	rabbitMQConnection    RabbitMQConnection
	channel               *amqp.Channel
	notifyOnClosedChannel chan *amqp.Error
	queue                 amqp.Queue
	name                  string
	consumer              string
	notifyOnClosedQueue   chan string
	deliveries            chan DeliveryChan
}

func NewDefaultRabbitMQQueue(rabbitMQConnection RabbitMQConnection, queue string, consumer string) *DefaultRabbitMQQueue {

	if rabbitMQConnection == nil {
		log.Fatal("starting up - error setting up rabbitMQueue: rabbitMQConnection is nil")
	}

	if strings.TrimSpace(queue) == "" {
		log.Fatal("starting up - error setting up rabbitMQueue: queue is empty")
	}

	if strings.TrimSpace(consumer) == "" {
		log.Fatal("starting up - error setting up rabbitMQueue: consumer is empty")
	}

	return &DefaultRabbitMQQueue{
		rabbitMQConnection:    rabbitMQConnection,
		notifyOnClosedChannel: make(chan *amqp.Error),
		name:                  queue,
		consumer:              consumer,
		notifyOnClosedQueue:   make(chan string),
		deliveries:            make(chan DeliveryChan, 1),
	}
}

func (queue *DefaultRabbitMQQueue) Consume() (chan DeliveryChan, error) {

	/*
		err := retry.Do(channel.connect, retry.Attempts(5),
			retry.OnRetry(func(n uint, err error) {
				log.Warn(fmt.Sprintf("rabbitmq channel - failed to connect: %s", err.Error()))
				log.Debug(fmt.Sprintf("rabbitmq channel - trying reconnection to channel"))
			}),
		)
	*/
	if err := queue.connect(); err != nil {
		log.Error(fmt.Sprintf("rabbitmq queue - failed connection to queue"))
		return nil, err
	}

	go queue.reconnect()

	return queue.deliveries, nil
}

func (queue *DefaultRabbitMQQueue) connect() error {

	var err error
	var connection *amqp.Connection
	if connection, err = queue.rabbitMQConnection.Connect(); err != nil {
		return err
	}

	if queue.channel, err = connection.Channel(); err != nil {
		return err
	}
	queue.notifyOnClosedChannel = queue.channel.NotifyClose(make(chan *amqp.Error))

	if queue.queue, err = queue.channel.QueueDeclare(queue.name, true, false, false, false, nil); err != nil {
		return err
	}
	queue.channel.NotifyCancel(queue.notifyOnClosedQueue)

	internal := make(DeliveryChan)
	if internal, err = queue.channel.Consume(queue.name, queue.consumer, true, false, false, false, nil); err != nil {
		return err
	}

	queue.deliveries <- internal

	log.Debug(fmt.Sprintf("rabbitmq queue - connected to queue %s", queue.name))

	return nil
}

func (queue *DefaultRabbitMQQueue) reconnect() {

	if !queue.rabbitMQConnection.RabbitMQContext().FailOver() {
		return
	}

	notifyOnClosedEvent := func() (*amqp.Error, bool) {
		select {
		case reason, ok := <-queue.notifyOnClosedChannel:
			return reason, ok
			/*
				case reason, ok := <-queue.notifyOnClosedQueue:
					return &amqp.Error{Reason: reason}, ok

			*/
		}
	}

	for {
		var ok bool
		var reason *amqp.Error
		if reason, ok = notifyOnClosedEvent(); !ok {
			break
		}
		log.Debug(fmt.Sprintf("rabbitmq queue - queue %s closed unexpectedly: %s", queue.name, reason.Reason))

		<-queue.RabbitMQContext().NotifyOnFaiOverConnection()
		time.Sleep(makeConnectionDelay)
		queue.Close()

		log.Debug(fmt.Sprintf("rabbitmq queue - trying reconnection to queue"))

		for {
			time.Sleep(makeConnectionDelay)
			if err := queue.connect(); err != nil {
				log.Error(fmt.Sprintf("rabbitmq queue - failed reconnection to queue %s: %s", queue.name, err.Error()))
				continue
			}
			log.Info(fmt.Sprintf("rabbitmq queue - reconnected to queue %s", queue.name))
			break
		}
	}
}

func (queue *DefaultRabbitMQQueue) Close() {
	if queue.channel != nil && !queue.channel.IsClosed() {
		log.Debug("rabbitmq queue - closing connection")
		if err := queue.channel.Close(); err != nil {
			log.Error(fmt.Sprintf("rabbitmq queue - failed to close connection to queue %s: %s", queue.name, err.Error()))
		}
	}
	log.Debug(fmt.Sprintf("rabbitmq queue - closed connection to queue %s", queue.name))
}

func (queue *DefaultRabbitMQQueue) RabbitMQContext() RabbitMQContext {
	return queue.rabbitMQConnection.RabbitMQContext()
}
