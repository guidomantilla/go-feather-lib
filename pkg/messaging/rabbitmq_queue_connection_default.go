package messaging

import (
	"fmt"
	"strings"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"

	"github.com/avast/retry-go/v4"
	"github.com/guidomantilla/go-feather-lib/pkg/common/log"
)

type DefaultRabbitMQQueueConnection struct {
	url                      string
	connection               *amqp.Connection
	notifyOnClosedConnection chan *amqp.Error
	channel                  *amqp.Channel
	notifyOnClosedChannel    chan *amqp.Error
	queueName                string
	queue                    amqp.Queue
	notifyOnClosedQueue      chan string
	isReady                  bool
	notifyOnClosingWatcher   chan bool
	receivedMessagesChan     <-chan amqp.Delivery
}

func NewDefaultRabbitMQQueueConnection(url string, username string, password string, queueName string) *DefaultRabbitMQQueueConnection {

	if strings.TrimSpace(url) == "" {
		log.Fatal("starting up - error setting up rabbitmq connection: url is empty")
	}

	if strings.TrimSpace(username) == "" {
		log.Fatal("starting up - error setting up rabbitmq connection: username is empty")
	}

	if strings.TrimSpace(password) == "" {
		log.Fatal("starting up - error setting up rabbitmq connection: password is empty")
	}

	if strings.TrimSpace(queueName) == "" {
		log.Fatal("starting up - error setting up rabbitmq connection: queueName is empty")
	}

	url = strings.Replace(url, ":username", username, 1)
	url = strings.Replace(url, ":password", password, 1)

	return &DefaultRabbitMQQueueConnection{
		url:                    url,
		queueName:              queueName,
		isReady:                false,
		notifyOnClosingWatcher: make(chan bool),
		receivedMessagesChan:   make(<-chan amqp.Delivery),
	}
}

//

func (client *DefaultRabbitMQQueueConnection) Start() {
	go client.watchConnection()
	<-time.After(time.Second)
}

func (client *DefaultRabbitMQQueueConnection) Close() {
	client.notifyOnClosingWatcher <- true
}

func (client *DefaultRabbitMQQueueConnection) Connect() (*amqp.Connection, *amqp.Channel, *amqp.Queue, <-chan amqp.Delivery, error) {

	if client.isReady {
		return client.connection, client.channel, &client.queue, client.receivedMessagesChan, nil
	}

	<-time.After(makeConnectionDelay)
	if !client.isReady {
		return nil, nil, nil, nil, fmt.Errorf("unable to connect to %s", client.url)
	}

	return client.connection, client.channel, &client.queue, client.receivedMessagesChan, nil
}

//

func (client *DefaultRabbitMQQueueConnection) makeConnection() error {

	var err error
	if client.connection, err = amqp.Dial(client.url); err != nil {
		return err
	}

	if client.channel, err = client.connection.Channel(); err != nil {
		return err
	}

	if client.queue, err = client.channel.QueueDeclare(client.queueName, true, false, false, false, nil); err != nil {
		return err
	}

	if client.receivedMessagesChan, err = client.channel.Consume(client.queue.Name, "xxx", false, false, false, false, nil); err != nil {
		log.Error(fmt.Sprintf("server starting up - rabbitmq queue dispatcher - error: %s", err.Error()))
		return err
	}

	client.notifyOnClosedConnection = make(chan *amqp.Error, 1)
	client.connection.NotifyClose(client.notifyOnClosedConnection)

	client.notifyOnClosedChannel = make(chan *amqp.Error, 1)
	client.channel.NotifyClose(client.notifyOnClosedChannel)

	client.notifyOnClosedQueue = make(chan string, 1)
	client.channel.NotifyCancel(client.notifyOnClosedQueue)

	client.isReady = true
	return nil
}

func (client *DefaultRabbitMQQueueConnection) watchConnection() {

	funcToRetry := func() error {
		var err error
		if err = client.makeConnection(); err != nil {
			<-time.After(makeConnectionDelay)
			return err
		}
		return nil
	}

	log.Info("rabbitmq - connection demon started")
	defer log.Info("rabbitmq - connection demon stopped")

	for {

		if !client.isReady {
			err := retry.Do(funcToRetry, retry.Attempts(10),
				retry.OnRetry(func(n uint, err error) {
					log.Info("rabbitmq - connection demon - failed to connect. retrying...")
				}),
			)
			if err != nil {
				log.Info("rabbitmq - connection demon - failed to connect")
				continue
			}
			log.Info("rabbitmq - connection ready")
		}

		select {

		case <-client.notifyOnClosedConnection:
			client.isReady = false
			log.Info("rabbitmq - connection closed. reconnecting...")
			continue

		case <-client.notifyOnClosedChannel:
			client.isReady = false
			log.Info("rabbitmq - channel closed. recreating...")
			continue

		case <-client.notifyOnClosedQueue:
			client.isReady = false
			client.channel.Cancel("xxx", true)
			log.Info("rabbitmq - queue closed. recreating...")
			continue

		case <-time.After(5 * time.Second):
			if client.isReady {
				var err error
				if _, err = client.channel.QueueInspect(client.queue.Name); err != nil { //nolint:staticcheck
					client.isReady = false
					log.Info("rabbitmq - failed to ping the queue. notifying...")
					continue
				}
			}
			continue

		case <-client.notifyOnClosingWatcher:
			client.isReady = false
			return
		}
	}
}
