package server

import (
	"context"

	"github.com/guidomantilla/go-feather-lib/pkg/common/assert"
	"github.com/guidomantilla/go-feather-lib/pkg/common/log"
)

type cronServer struct {
	ctx          context.Context
	internal     CronServer
	closeChannel chan struct{}
}

func NewCronServer(cron CronServer) *cronServer {
	assert.NotNil(cron, "starting up - error setting up cron server: cron is nil")

	return &cronServer{
		internal:     cron,
		closeChannel: make(chan struct{}),
	}
}

func (server *cronServer) Run(ctx context.Context) error {

	server.ctx = ctx
	log.Info("starting up - starting cron server")
	server.internal.Start()
	<-server.closeChannel
	return nil
}

func (server *cronServer) Stop(_ context.Context) error {

	log.Info("shutting down - stopping cron server")
	close(server.closeChannel)
	server.internal.Stop()
	log.Debug("shutting down - cron server stopped")
	return nil
}
