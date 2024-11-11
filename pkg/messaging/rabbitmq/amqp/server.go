package amqp

import (
	"context"
	"fmt"

	"github.com/qmdx00/lifecycle"

	"github.com/guidomantilla/go-feather-lib/pkg/common/assert"
	"github.com/guidomantilla/go-feather-lib/pkg/common/log"
)

func BuildConsumerServer(consumers ...Consumer) (string, lifecycle.Server) {
	return "rabbitmq-amqp-server", NewConsumerServer(consumers...)
}

//

type ConsumerServer struct {
	ctx          context.Context
	consumers    []Consumer
	closeChannel chan struct{}
}

func NewConsumerServer(consumers ...Consumer) *ConsumerServer {
	assert.NotEmpty(consumers, "starting up - error setting up rabbitmq server: consumers is empty")

	return &ConsumerServer{
		consumers:    consumers,
		closeChannel: make(chan struct{}),
	}
}

func (server *ConsumerServer) Run(ctx context.Context) error {
	assert.NotNil(ctx, "rabbitmq server - error starting up: context is nil")

	server.ctx = ctx
	log.Info(ctx, fmt.Sprintf("starting up - starting rabbitmq server: %s", server.consumers[0].Context().Server()))

	for _, consumer := range server.consumers {
		go func(ctx context.Context, consumer Consumer, closeChannel chan struct{}) {
			for {
				select {
				case <-closeChannel:
					return

				default:
					var err error
					var closeChannel chan string
					if closeChannel, err = consumer.Consume(ctx); err != nil {
						log.Error(ctx, fmt.Sprintf("rabbitmq server - error: %s", err.Error()))
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

func (server *ConsumerServer) Stop(ctx context.Context) error {
	assert.NotNil(server.ctx, "rabbitmq server - error shutting down: context is nil")

	log.Debug(ctx, "server shutting down - stopping rabbitmq server")
	close(server.closeChannel)
	for _, consumer := range server.consumers {
		consumer.Close(ctx)
	}
	log.Debug(ctx, "server shutting down - rabbitmq server stopped")
	return nil
}
