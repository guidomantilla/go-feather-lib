package server

import (
	"context"
	"fmt"

	"github.com/qmdx00/lifecycle"
	amqp "github.com/rabbitmq/amqp091-go"

	"github.com/guidomantilla/go-feather-lib/pkg/common/log"
	"github.com/guidomantilla/go-feather-lib/pkg/messaging"
)

type RabbitMQQueueMessageDispatcher struct {
	ctx                  context.Context
	connection           messaging.RabbitMQQueueConnection
	listener             messaging.RabbitMQQueueMessageListener
	receivedMessagesChan <-chan amqp.Delivery
}

func BuildRabbitMQQueueMessageDispatcher(connection messaging.RabbitMQQueueConnection, listener messaging.RabbitMQQueueMessageListener) Server {

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

	server.connection.Start()

	var err error
	var channel *amqp.Channel
	var queue *amqp.Queue

	if _, channel, queue, err = server.connection.Connect(); err != nil {
		log.Error(fmt.Sprintf("server starting up - rabbitmq queue dispatcher - error: %s", err.Error()))
		return err
	}

	if server.receivedMessagesChan, err = channel.Consume(queue.Name, "", true, false, false, false, nil); err != nil {
		log.Error(fmt.Sprintf("server starting up - rabbitmq queue dispatcher - error: %s", err.Error()))
		return err
	}

	if err = server.ListenAndDispatch(); err != nil {
		log.Error(fmt.Sprintf("server starting up - rabbitmq queue dispatcher - error: %s", err.Error()))
		return err
	}

	return nil
}

func (server *RabbitMQQueueMessageDispatcher) ListenAndDispatch() error {

	for {
		select {
		case <-server.ctx.Done():
			return nil
		case message := <-server.receivedMessagesChan:
			go server.Dispatch(&message)
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
