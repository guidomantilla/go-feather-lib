package server

import (
	"context"
	"github.com/guidomantilla/go-feather-lib/pkg/common/log"
)

type DummyServer struct {
	ctx     context.Context
	channel chan string
}

func BuildDummyServer() Server {
	return &DummyServer{
		channel: make(chan string),
	}
}

func (server *DummyServer) Run(ctx context.Context) error {
	log.Info("starting up - starting dummy server")
	<-server.channel
	return nil
}

func (server *DummyServer) Stop(_ context.Context) error {

	log.Info("shutting down - stopping dummy server")
	server.channel <- "stop"
	log.Info("shutting down - dummy server stopped")
	return nil
}
