package messaging

import (
	"fmt"
	"strings"
	"sync"

	amqp "github.com/rabbitmq/amqp091-go"

	"github.com/guidomantilla/go-feather-lib/pkg/common/log"
)

type RabbitMQQueue struct {
	messagingConnection MessagingConnection[*amqp.Connection]
	channel             *amqp.Channel
	queue               amqp.Queue
	name                string
	consumer            string
	mu                  sync.Mutex
}

func NewRabbitMQQueue(messagingConnection MessagingConnection[*amqp.Connection], queue string) *RabbitMQQueue {

	if messagingConnection == nil {
		log.Fatal("starting up - error setting up rabbitMQQueue: messagingConnection is nil")
	}

	if strings.TrimSpace(queue) == "" {
		log.Fatal("starting up - error setting up rabbitMQQueue: queue is empty")
	}

	return &RabbitMQQueue{
		messagingConnection: messagingConnection,
		name:                queue,
		consumer:            "consumer-" + queue,
	}
}

func (queue *RabbitMQQueue) Consume() (MessagingEvent, error) {

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

	var deliveries <-chan amqp.Delivery
	if deliveries, err = queue.channel.Consume(queue.name, queue.consumer, true, false, false, false, nil); err != nil {
		log.Debug(fmt.Sprintf("rabbitmq queue - failed comsuming from queue: %s", err.Error()))
		return nil, err
	}

	closeChannel := make(chan string)
	go func(closeChannel chan string) {
		for message := range deliveries {
			go log.Info(fmt.Sprintf("rabbitmq queue - message received: %s", message.Body))
		}
		close(closeChannel)
		log.Debug(fmt.Sprintf("rabbitmq queue - disconected to queue %s", queue.name))
	}(closeChannel)

	return closeChannel, nil
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
