package server

import (
	"context"

	"github.com/guidomantilla/go-feather-lib/pkg/common/log"
)

type BaseServer struct {
	ctx          context.Context
	closeChannel chan struct{}
}

func NewBaseServer() *BaseServer {
	return &BaseServer{
		closeChannel: make(chan struct{}),
	}
}

func (server *BaseServer) Run(ctx context.Context) error {

	server.ctx = ctx
	log.Info("starting up - starting base server")
	<-server.closeChannel
	return nil
}

func (server *BaseServer) Stop(_ context.Context) error {

	log.Info("shutting down - stopping base server")
	close(server.closeChannel)
	log.Debug("shutting down - default base stopped")
	return nil
}
