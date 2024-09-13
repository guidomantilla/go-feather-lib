package server

import (
	"context"
	"fmt"

	"github.com/qmdx00/lifecycle"
	amqp "github.com/rabbitmq/amqp091-go"

	"github.com/guidomantilla/go-feather-lib/pkg/common/log"
	"github.com/guidomantilla/go-feather-lib/pkg/messaging"
)

var _ lifecycle.Server = (*RabbitMQMessageDispatcher)(nil)

type RabbitMQMessageDispatcher struct {
	ctx                  context.Context
	connection           messaging.RabbitMQQueueConnection
	listener             messaging.RabbitMQMessageListener
	receivedMessagesChan <-chan amqp.Delivery
}

func BuildRabbitMQMessageDispatcher(connection messaging.RabbitMQQueueConnection, listener messaging.RabbitMQMessageListener) lifecycle.Server {

	if connection == nil {
		log.Fatal("starting up - error setting up rabbitmq dispatcher: connection is nil")
	}

	if listener == nil {
		log.Fatal("starting up - error setting up rabbitmq dispatcher: listener is nil")
	}

	return &RabbitMQMessageDispatcher{
		connection:           connection,
		listener:             listener,
		receivedMessagesChan: make(<-chan amqp.Delivery),
	}
}

func (server *RabbitMQMessageDispatcher) Run(ctx context.Context) error {

	server.ctx = ctx
	info, _ := lifecycle.FromContext(ctx)
	log.Info(fmt.Sprintf("server starting up - starting rabbitmq dispatcher %s, v.%s", info.Name(), info.Version()))

	server.connection.Start()

	var err error
	var channel *amqp.Channel
	var queue *amqp.Queue

	if _, channel, queue, err = server.connection.Connect(); err != nil {
		log.Error(fmt.Sprintf("server starting up - rabbitmq dispatcher - error: %s", err.Error()))
		return err
	}

	if server.receivedMessagesChan, err = channel.Consume(queue.Name, "", true, false, false, false, nil); err != nil {
		log.Error(fmt.Sprintf("server starting up - rabbitmq dispatcher - error: %s", err.Error()))
		return err
	}

	if err = server.ListenAndDispatch(); err != nil {
		log.Error(fmt.Sprintf("server starting up - rabbitmq dispatcher - error: %s", err.Error()))
		return err
	}

	return nil
}

func (server *RabbitMQMessageDispatcher) ListenAndDispatch() error {

	for {
		select {
		case <-server.ctx.Done():
			return nil
		case message := <-server.receivedMessagesChan:
			go server.Dispatch(&message)
		}
	}
}

func (server *RabbitMQMessageDispatcher) Dispatch(message *amqp.Delivery) {

	var err error
	if err = server.listener.OnMessage(message); err != nil {
		log.Error(fmt.Sprintf("rabbitmq listener - error: %s, message: %s", err.Error(), message.Body))
	}
}

func (server *RabbitMQMessageDispatcher) Stop(ctx context.Context) error {

	info, _ := lifecycle.FromContext(ctx)
	log.Info(fmt.Sprintf("server shutting down - stopping rabbitmq dispatcher %s, v.%s", info.Name(), info.Version()))

	server.connection.Close()

	log.Info("server shutting down - rabbitmq dispatcher stopped")
	return nil
}
