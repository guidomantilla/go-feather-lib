package server

import (
	"context"

	"github.com/guidomantilla/go-feather-lib/pkg/common/log"
)

type DefaultServer struct {
	ctx          context.Context
	closeChannel chan struct{}
}

func NewDefaultServer() Server {
	return &DefaultServer{
		closeChannel: make(chan struct{}),
	}
}

func (server *DefaultServer) Run(ctx context.Context) error {

	server.ctx = ctx
	log.Info("starting up - starting default server")
	<-server.closeChannel
	return nil
}

func (server *DefaultServer) Stop(_ context.Context) error {

	log.Info("shutting down - stopping default server")
	close(server.closeChannel)
	log.Debug("shutting down - default server stopped")
	return nil
}
