package server

import (
	"context"
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"

	"github.com/guidomantilla/go-feather-lib/pkg/common/log"
	"github.com/guidomantilla/go-feather-lib/pkg/messaging"
)

type RabbitMQMessageDispatcher struct {
	ctx        context.Context
	listener   messaging.MessagingListener[*amqp.Delivery]
	targets    []messaging.MessagingTarget[*amqp.Channel]
	deliveries <-chan amqp.Delivery
	stopCh     chan struct{}
}

func BuildRabbitMQMessageDispatcher(listener messaging.MessagingListener[*amqp.Delivery], targets ...messaging.MessagingTarget[*amqp.Channel]) Server {

	if listener == nil {
		log.Fatal("starting up - error setting up rabbitmq dispatcher: listener is nil")
	}

	if len(targets) == 0 {
		log.Fatal("starting up - error setting up rabbitmq dispatcher: targets is empty")
	}

	return &RabbitMQMessageDispatcher{
		listener:   listener,
		targets:    targets,
		deliveries: make(<-chan amqp.Delivery),
		stopCh:     make(chan struct{}),
	}
}

func (server *RabbitMQMessageDispatcher) Run(ctx context.Context) error {

	server.ctx = ctx
	log.Info(fmt.Sprintf("starting up - starting rabbitmq dispatcher: %s", server.targets[0].MessagingContext().Server()))

	for _, target := range server.targets {
		go func(target messaging.MessagingTarget[*amqp.Channel]) {
			for {
				select {
				case <-server.stopCh:
					return

				default:
					var err error

					var channel *amqp.Channel
					if channel, err = target.Connect(); err != nil {
						log.Error(fmt.Sprintf("rabbitmq dispatcher - error: %s", err.Error()))
						continue
					}

					var deliveries <-chan amqp.Delivery
					if deliveries, err = channel.Consume(target.Name(), target.Consumer(), true, false, false, false, nil); err != nil {
						log.Error(fmt.Sprintf("rabbitmq dispatcher - error: %s", err.Error()))
						continue
					}

					for message := range deliveries {
						go server.Dispatch(&message)
					}
				}
			}
		}(target)
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
	for _, queue := range server.targets {
		queue.Close()
	}
	log.Debug("server shutting down - rabbitmq dispatcher stopped")
	return nil
}
