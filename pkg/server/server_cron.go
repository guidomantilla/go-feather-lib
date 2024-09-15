package server

import (
	"context"

	cron "github.com/robfig/cron/v3"

	"github.com/guidomantilla/go-feather-lib/pkg/common/log"
)

type CronServer struct {
	ctx      context.Context
	internal *cron.Cron
}

func BuildCronServer(cron *cron.Cron) Server {

	if cron == nil {
		log.Fatal("starting up - error setting up cron server: cron is nil")
	}

	return &CronServer{
		internal: cron,
	}
}

func (server *CronServer) Run(ctx context.Context) error {

	server.ctx = ctx
	log.Info("starting up - starting cron server")
	server.internal.Start()
	return nil
}

func (server *CronServer) Stop(_ context.Context) error {

	log.Info("shutting down - stopping cron server")
	server.internal.Stop()
	log.Debug("shutting down - cron server stopped")
	return nil
}
