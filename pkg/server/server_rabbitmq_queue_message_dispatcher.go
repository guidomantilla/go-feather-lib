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
	ctx           context.Context
	listener      messaging.RabbitMQMessageListener
	rabbitMQQueue []messaging.RabbitMQQueue
	deliveries    <-chan amqp.Delivery
	stopCh        chan struct{}
}

func BuildRabbitMQQueueMessageDispatcher(listener messaging.RabbitMQMessageListener, rabbitMQQueue ...messaging.RabbitMQQueue) Server {

	if listener == nil {
		log.Fatal("starting up - error setting up rabbitmq queue dispatcher: listener is nil")
	}

	if len(rabbitMQQueue) == 0 {
		log.Fatal("starting up - error setting up rabbitmq queue dispatcher: rabbitMQQueue is empty")
	}

	return &RabbitMQQueueMessageDispatcher{
		listener:      listener,
		rabbitMQQueue: rabbitMQQueue,
		deliveries:    make(<-chan amqp.Delivery),
		stopCh:        make(chan struct{}),
	}
}

func (server *RabbitMQQueueMessageDispatcher) Run(ctx context.Context) error {

	server.ctx = ctx
	info, _ := lifecycle.FromContext(ctx)
	log.Info(fmt.Sprintf("server starting up - starting rabbitmq queue dispatcher %s, v.%s", info.Name(), info.Version()))

	for _, queue := range server.rabbitMQQueue {
		go func(queue messaging.RabbitMQQueue) {
			for {
				select {
				case <-server.stopCh:
					return
				default:
					rabbitChannel, _ := queue.Connect()
					deliveries, _ := rabbitChannel.Consume(queue.Name(), queue.Consumer(), true, false, false, false, nil)
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
	close(server.stopCh)
	for _, queue := range server.rabbitMQQueue {
		queue.Close()
	}
	log.Debug("server shutting down - rabbitmq queue dispatcher stopped")
	return nil
}
