package messaging

import (
	"fmt"
	"strings"
	"sync"

	amqp "github.com/rabbitmq/amqp091-go"

	"github.com/guidomantilla/go-feather-lib/pkg/common/log"
)

type RabbitMQQueue struct {
	messagingConnection   MessagingConnection[*amqp.Connection]
	channel               *amqp.Channel
	notifyOnClosedChannel chan *amqp.Error
	queue                 amqp.Queue
	name                  string
	consumer              string
	notifyOnClosedQueue   chan string
	mu                    sync.Mutex
}

func NewRabbitMQQueue(messagingConnection MessagingConnection[*amqp.Connection], queue string) *RabbitMQQueue {

	if messagingConnection == nil {
		log.Fatal("starting up - error setting up rabbitMQQueue: messagingConnection is nil")
	}

	if strings.TrimSpace(queue) == "" {
		log.Fatal("starting up - error setting up rabbitMQQueue: queue is empty")
	}

	return &RabbitMQQueue{
		messagingConnection:   messagingConnection,
		notifyOnClosedChannel: make(chan *amqp.Error),
		name:                  queue,
		consumer:              "consumer-" + queue,
		notifyOnClosedQueue:   make(chan string),
	}
}

func (queue *RabbitMQQueue) Connect() (*amqp.Channel, error) {

	queue.mu.Lock()
	defer queue.mu.Unlock()

	var err error
	var connection *amqp.Connection
	if connection, err = queue.messagingConnection.Connect(); err != nil {
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

func (queue *RabbitMQQueue) Close() {
	if queue.channel != nil && !queue.channel.IsClosed() {
		log.Debug("rabbitmq queue - closing connection")
		if err := queue.channel.Close(); err != nil {
			log.Error(fmt.Sprintf("rabbitmq queue - failed to close connection to queue %s: %s", queue.name, err.Error()))
		}
	}
	queue.channel = nil
	queue.messagingConnection.Close()
	log.Debug(fmt.Sprintf("rabbitmq queue - closed connection to queue %s", queue.name))
}

func (queue *RabbitMQQueue) MessagingContext() MessagingContext {
	return queue.messagingConnection.MessagingContext()
}

func (queue *RabbitMQQueue) Name() string {
	return queue.name
}

func (queue *RabbitMQQueue) Consumer() string {
	return queue.consumer
}
