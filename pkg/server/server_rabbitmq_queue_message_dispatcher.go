package server

import (
	"context"
	"errors"
	"fmt"
	"github.com/qmdx00/lifecycle"
	amqp "github.com/rabbitmq/amqp091-go"

	"github.com/guidomantilla/go-feather-lib/pkg/common/log"
	"github.com/guidomantilla/go-feather-lib/pkg/messaging"
)

type RabbitMQQueueMessageDispatcher struct {
	ctx                      context.Context
	connection               messaging.RabbitMQConnection
	listener                 messaging.RabbitMQQueueMessageListener
	receivedMessagesChan     <-chan amqp.Delivery
	notifyOnClosedConnection chan *amqp.Error
	notifyOnClosedChannel    chan *amqp.Error
	notifyOnClosedQueue      chan string
}

func BuildRabbitMQQueueMessageDispatcher(messagingContext messaging.MessagingContext, connection messaging.RabbitMQConnection, listener messaging.RabbitMQQueueMessageListener) Server {

	if messagingContext == nil {
		log.Fatal("starting up - error setting up rabbitmq queue dispatcher: messagingContext is nil")
	}

	if connection == nil {
		log.Fatal("starting up - error setting up rabbitmq queue dispatcher: connection is nil")
	}

	if listener == nil {
		log.Fatal("starting up - error setting up rabbitmq queue dispatcher: listener is nil")
	}

	return &RabbitMQQueueMessageDispatcher{
		connection:           connection,
		listener:             listener,
		receivedMessagesChan: make(<-chan amqp.Delivery),
	}
}

func (server *RabbitMQQueueMessageDispatcher) Run(ctx context.Context) error {

	server.ctx = ctx
	info, _ := lifecycle.FromContext(ctx)
	log.Info(fmt.Sprintf("server starting up - starting rabbitmq queue dispatcher %s, v.%s", info.Name(), info.Version()))

	if err := server.listen(); err != nil {
		log.Error(fmt.Sprintf("server starting up - rabbitmq queue dispatcher - error: %s", err.Error()))
		return err
	}

	for {
		select {
		case <-server.ctx.Done():
			return nil
		case <-server.ctx.Done():
			//log.Warn(fmt.Sprintf("rabbitmq queue dispatcher - connection closed unexpectedly: %s", reason))
			if err := server.listen(); err != nil {
				log.Error(fmt.Sprintf("rabbitmq queue dispatcher - failure reestablishing connection: %s", err.Error()))
				continue
			}
			log.Info(fmt.Sprintf("rabbitmq queue dispatcher - connection reestablish to queue %s", server.listener.Queue()))
		case reason, _ := <-server.notifyOnClosedQueue:
			return errors.New(fmt.Sprintf("rabbitmq queue dispatcher - queue %s closed unexpectedly: %s", server.listener.Queue(), reason))
		case message, ok := <-server.receivedMessagesChan:
			if ok {
				go server.Dispatch(&message)
			}
		}
	}
}

func (server *RabbitMQQueueMessageDispatcher) listen() error {

	var err error
	var connection *amqp.Connection
	if connection, err = server.connection.Connect(); err != nil {
		return err
	}

	var channel *amqp.Channel
	if channel, err = connection.Channel(); err != nil {
		return err
	}

	var queue amqp.Queue
	if queue, err = channel.QueueDeclare(server.listener.Queue(), true, false, false, false, nil); err != nil {
		return err
	}
	if server.notifyOnClosedQueue == nil {
		server.notifyOnClosedQueue = make(chan string)
	}
	channel.NotifyCancel(server.notifyOnClosedQueue)

	if server.receivedMessagesChan, err = channel.Consume(queue.Name, "xxx", false, false, false, false, nil); err != nil {
		return err
	}

	return nil
}

func (server *RabbitMQQueueMessageDispatcher) ListenAndDispatch() error {

	for {
		select {
		case <-server.ctx.Done():
			return nil
		case <-server.notifyOnClosedQueue:
			log.Warn("server starting up - rabbitmq queue dispatcher - queue closed")
			return nil
		case message, ok := <-server.receivedMessagesChan:
			if ok {
				go server.Dispatch(&message)
			}
		}
	}
}

func (server *RabbitMQQueueMessageDispatcher) Dispatch(message any) {

	var err error
	msg := message.(*amqp.Delivery)
	if err = server.listener.OnMessage(msg); err != nil {
		log.Error(fmt.Sprintf("rabbitmq queue dispatcher - error: %s, message: %s", err.Error(), msg.Body))
	}
}

func (server *RabbitMQQueueMessageDispatcher) Stop(ctx context.Context) error {

	info, _ := lifecycle.FromContext(ctx)
	log.Info(fmt.Sprintf("server shutting down - stopping rabbitmq queue dispatcher %s, v.%s", info.Name(), info.Version()))

	server.connection.Close()

	log.Info("server shutting down - rabbitmq queue dispatcher stopped")
	return nil
}
