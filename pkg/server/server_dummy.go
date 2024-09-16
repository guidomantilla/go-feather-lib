package server

import (
	"context"

	"github.com/guidomantilla/go-feather-lib/pkg/common/log"
)

type DummyServer struct {
	ctx          context.Context
	closeChannel chan struct{}
}

func BuildDummyServer() Server {
	return &DummyServer{
		closeChannel: make(chan struct{}),
	}
}

func (server *DummyServer) Run(ctx context.Context) error {

	server.ctx = ctx
	log.Info("starting up - starting dummy server")
	<-server.closeChannel
	return nil
}

func (server *DummyServer) Stop(_ context.Context) error {

	log.Info("shutting down - stopping dummy server")
	close(server.closeChannel)
	log.Debug("shutting down - dummy server stopped")
	return nil
}
