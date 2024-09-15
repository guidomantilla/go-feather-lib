package messaging

import (
	"fmt"
	"strings"
	"sync"

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
	mu                    sync.Mutex
}

func NewDefaultRabbitMQQueue(rabbitMQConnection RabbitMQConnection, queue string) *DefaultRabbitMQQueue {

	if rabbitMQConnection == nil {
		log.Fatal("starting up - error setting up rabbitMQueue: rabbitMQConnection is nil")
	}

	if strings.TrimSpace(queue) == "" {
		log.Fatal("starting up - error setting up rabbitMQueue: queue is empty")
	}

	return &DefaultRabbitMQQueue{
		rabbitMQConnection:    rabbitMQConnection,
		notifyOnClosedChannel: make(chan *amqp.Error),
		name:                  queue,
		consumer:              "consumer-" + queue,
		notifyOnClosedQueue:   make(chan string),
	}
}

func (queue *DefaultRabbitMQQueue) Connect() (*amqp.Channel, error) {

	queue.mu.Lock()
	defer queue.mu.Unlock()

	var err error
	var connection *amqp.Connection
	if connection, err = queue.rabbitMQConnection.Connect(); err != nil {
		log.Debug(fmt.Sprintf("rabbitmq queue - failed connection to queue %s: %s", queue.name, err.Error()))
		return nil, err
	}

	if !(queue.channel != nil && !queue.channel.IsClosed()) {
		if queue.channel, err = connection.Channel(); err != nil {
			log.Debug(fmt.Sprintf("rabbitmq queue - failed connection to queue %s: %s", queue.name, err.Error()))
			return nil, err
		}
	}

	if queue.queue, err = queue.channel.QueueDeclare(queue.name, true, false, false, false, nil); err != nil {
		log.Debug(fmt.Sprintf("rabbitmq queue - failed connection to queue %s: %s", queue.name, err.Error()))
		return nil, err
	}

	log.Debug(fmt.Sprintf("rabbitmq queue - connected to queue %s", queue.name))

	return queue.channel, nil
}

func (queue *DefaultRabbitMQQueue) Close() {
	if queue.channel != nil && !queue.channel.IsClosed() {
		log.Debug("rabbitmq queue - closing connection")
		if err := queue.channel.Close(); err != nil {
			log.Error(fmt.Sprintf("rabbitmq queue - failed to close connection to queue %s: %s", queue.name, err.Error()))
		}
	}
	queue.channel = nil
	log.Debug(fmt.Sprintf("rabbitmq queue - closed connection to queue %s", queue.name))
}

func (queue *DefaultRabbitMQQueue) RabbitMQContext() RabbitMQContext {
	return queue.rabbitMQConnection.RabbitMQContext()
}

func (queue *DefaultRabbitMQQueue) Name() string {
	return queue.name
}

func (queue *DefaultRabbitMQQueue) Consumer() string {
	return queue.consumer
}
