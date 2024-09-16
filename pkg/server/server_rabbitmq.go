package server

import (
	"context"
	"fmt"

	"github.com/guidomantilla/go-feather-lib/pkg/common/log"
	"github.com/guidomantilla/go-feather-lib/pkg/messaging"
)

type RabbitMQServer struct {
	ctx     context.Context
	targets []messaging.MessagingTarget
	stopCh  chan struct{}
}

func BuildRabbitMQServer(targets ...messaging.MessagingTarget) Server {

	if len(targets) == 0 {
		log.Fatal("starting up - error setting up rabbitmq server: targets is empty")
	}

	return &RabbitMQServer{
		targets: targets,
		stopCh:  make(chan struct{}),
	}
}

func (server *RabbitMQServer) Run(ctx context.Context) error {

	server.ctx = ctx
	log.Info(fmt.Sprintf("starting up - starting rabbitmq server: %s", server.targets[0].MessagingContext().Server()))

	for _, target := range server.targets {
		go func(target messaging.MessagingTarget) {
			for {
				select {
				case <-server.stopCh:
					return

				default:
					var err error
					var closeChannel chan string
					if closeChannel, err = target.Consume(); err != nil {
						log.Error(fmt.Sprintf("rabbitmq server - error: %s", err.Error()))
						continue
					}
					<-closeChannel
				}
			}
		}(target)
	}
	return nil
}

func (server *RabbitMQServer) Stop(_ context.Context) error {

	log.Debug("server shutting down - stopping rabbitmq server")
	close(server.stopCh)
	for _, target := range server.targets {
		target.Close()
	}
	log.Debug("server shutting down - rabbitmq server stopped")
	return nil
}
