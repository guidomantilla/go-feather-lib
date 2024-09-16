package messaging

import (
	"fmt"
	"strings"
	"sync"

	amqp "github.com/rabbitmq/amqp091-go"

	"github.com/guidomantilla/go-feather-lib/pkg/common/log"
)

type RabbitMQConsumer struct {
	messagingConnection MessagingConnection[*amqp.Connection]
	channel             *amqp.Channel
	queue               amqp.Queue
	name                string
	consumer            string
	mu                  sync.Mutex
}

func NewRabbitMQConsumer(messagingConnection MessagingConnection[*amqp.Connection], queue string) *RabbitMQConsumer {

	if messagingConnection == nil {
		log.Fatal("starting up - error setting up rabbitmq consumer: messagingConnection is nil")
	}

	if strings.TrimSpace(queue) == "" {
		log.Fatal("starting up - error setting up rabbitmq consumer: queue is empty")
	}

	return &RabbitMQConsumer{
		messagingConnection: messagingConnection,
		name:                queue,
		consumer:            "consumer-" + queue,
	}
}

func (queue *RabbitMQConsumer) Consume() (MessagingEvent, error) {

	queue.mu.Lock()
	defer queue.mu.Unlock()

	var err error
	var connection *amqp.Connection
	if connection, err = queue.messagingConnection.Connect(); err != nil {
		log.Debug(fmt.Sprintf("rabbitmq consumer - failed connection to queue %s: %s", queue.name, err.Error()))
		return nil, err
	}

	if !(queue.channel != nil && !queue.channel.IsClosed()) {
		if queue.channel, err = connection.Channel(); err != nil {
			log.Debug(fmt.Sprintf("rabbitmq consumer - failed connection to queue %s: %s", queue.name, err.Error()))
			return nil, err
		}
	}

	if queue.queue, err = queue.channel.QueueDeclare(queue.name, true, false, false, false, nil); err != nil {
		log.Debug(fmt.Sprintf("rabbitmq consumer - failed connection to queue %s: %s", queue.name, err.Error()))
		return nil, err
	}

	log.Debug(fmt.Sprintf("rabbitmq consumer - connected to queue %s", queue.name))

	var deliveries <-chan amqp.Delivery
	if deliveries, err = queue.channel.Consume(queue.name, queue.consumer, true, false, false, false, nil); err != nil {
		log.Debug(fmt.Sprintf("rabbitmq consumer - failed comsuming from queue: %s", err.Error()))
		return nil, err
	}

	closeChannel := make(chan string)
	go func(closeChannel chan string) {
		for message := range deliveries {
			go log.Info(fmt.Sprintf("rabbitmq consumer - message received: %s", message.Body))
		}
		if err := queue.channel.Close(); err != nil {
			return
		}
		close(closeChannel)
		log.Debug(fmt.Sprintf("rabbitmq consumer - disconected from queue %s", queue.name))
	}(closeChannel)

	return closeChannel, nil
}

func (queue *RabbitMQConsumer) Close() {
	if queue.channel != nil && !queue.channel.IsClosed() {
		log.Debug("rabbitmq consumer - closing connection")
		if err := queue.channel.Close(); err != nil {
			log.Error(fmt.Sprintf("rabbitmq consumer - failed to close connection to queue %s: %s", queue.name, err.Error()))
		}
	}
	queue.channel = nil
	queue.messagingConnection.Close()
	log.Debug(fmt.Sprintf("rabbitmq consumer - closed connection to queue %s", queue.name))
}

func (queue *RabbitMQConsumer) MessagingContext() MessagingContext {
	return queue.messagingConnection.MessagingContext()
}
