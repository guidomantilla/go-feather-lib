package server

import (
	"context"
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"

	"github.com/guidomantilla/go-feather-lib/pkg/common/log"
	"github.com/guidomantilla/go-feather-lib/pkg/messaging"
)

type RabbitMQMessageDispatcher struct {
	ctx           context.Context
	listener      messaging.RabbitMQMessageListener
	rabbitMQQueue []messaging.RabbitMQQueue
	deliveries    <-chan amqp.Delivery
	stopCh        chan struct{}
}

func BuildRabbitMQMessageDispatcher(listener messaging.RabbitMQMessageListener, rabbitMQQueue ...messaging.RabbitMQQueue) Server {

	if listener == nil {
		log.Fatal("starting up - error setting up rabbitmq dispatcher: listener is nil")
	}

	if len(rabbitMQQueue) == 0 {
		log.Fatal("starting up - error setting up rabbitmq dispatcher: rabbitMQQueue is empty")
	}

	return &RabbitMQMessageDispatcher{
		listener:      listener,
		rabbitMQQueue: rabbitMQQueue,
		deliveries:    make(<-chan amqp.Delivery),
		stopCh:        make(chan struct{}),
	}
}

func (server *RabbitMQMessageDispatcher) Run(ctx context.Context) error {

	server.ctx = ctx
	log.Info(fmt.Sprintf("starting up - starting rabbitmq dispatcher: %s", server.rabbitMQQueue[0].RabbitMQContext().Server()))

	for _, queue := range server.rabbitMQQueue {
		go func(queue messaging.RabbitMQQueue) {
			for {
				select {
				case <-server.stopCh:
					return

				default:
					var err error

					var rabbitChannel *amqp.Channel
					if rabbitChannel, err = queue.Connect(); err != nil {
						log.Error(fmt.Sprintf("rabbitmq dispatcher - error: %s", err.Error()))
						continue
					}

					var deliveries <-chan amqp.Delivery
					if deliveries, err = rabbitChannel.Consume(queue.Name(), queue.Consumer(), true, false, false, false, nil); err != nil {
						log.Error(fmt.Sprintf("rabbitmq dispatcher - error: %s", err.Error()))
						continue
					}

					for message := range deliveries {
						go server.Dispatch(&message)
					}
				}
			}
		}(queue)
	}

	<-server.ctx.Done()
	return nil
}

func (server *RabbitMQMessageDispatcher) Dispatch(message any) {

	var err error
	msg := message.(*amqp.Delivery)
	if err = server.listener.OnMessage(msg); err != nil {
		log.Error(fmt.Sprintf("rabbitmq dispatcher - error: %s, message: %s", err.Error(), msg.Body))
	}
}

func (server *RabbitMQMessageDispatcher) Stop(ctx context.Context) error {

	log.Debug("server shutting down - stopping rabbitmq dispatcher")
	close(server.stopCh)
	for _, queue := range server.rabbitMQQueue {
		queue.Close()
	}
	log.Debug("server shutting down - rabbitmq dispatcher stopped")
	return nil
}
