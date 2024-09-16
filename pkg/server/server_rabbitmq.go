package server

import (
	"context"
	"fmt"

	"github.com/guidomantilla/go-feather-lib/pkg/common/log"
	"github.com/guidomantilla/go-feather-lib/pkg/messaging"
)

type RabbitMQServer struct {
	ctx          context.Context
	consumers    []messaging.MessagingConsumer
	closeChannel chan struct{}
}

func BuildRabbitMQServer(consumers ...messaging.MessagingConsumer) Server {

	if len(consumers) == 0 {
		log.Fatal("starting up - error setting up rabbitmq server: consumers is empty")
	}

	return &RabbitMQServer{
		consumers:    consumers,
		closeChannel: make(chan struct{}),
	}
}

func (server *RabbitMQServer) Run(ctx context.Context) error {

	server.ctx = ctx
	log.Info(fmt.Sprintf("starting up - starting rabbitmq server: %s", server.consumers[0].MessagingContext().Server()))

	for _, consumer := range server.consumers {
		go func(ctx context.Context, consumer messaging.MessagingConsumer, closeChannel chan struct{}) {
			for {
				select {
				case <-closeChannel:
					return

				default:
					var err error
					var closeChannel chan string
					if closeChannel, err = consumer.Consume(ctx); err != nil {
						log.Error(fmt.Sprintf("rabbitmq server - error: %s", err.Error()))
						continue
					}
					<-closeChannel
				}
			}
		}(ctx, consumer, server.closeChannel)
	}
	<-server.closeChannel
	return nil
}

func (server *RabbitMQServer) Stop(_ context.Context) error {

	log.Debug("server shutting down - stopping rabbitmq server")
	close(server.closeChannel)
	for _, consumer := range server.consumers {
		consumer.Close()
	}
	log.Debug("server shutting down - rabbitmq server stopped")
	return nil
}
