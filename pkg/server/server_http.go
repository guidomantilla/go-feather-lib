package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/guidomantilla/go-feather-lib/pkg/common/assert"
	"github.com/guidomantilla/go-feather-lib/pkg/common/log"
)

type HttpServer struct {
	ctx      context.Context
	internal *http.Server
}

func NewHttpServer(server *http.Server) *HttpServer {
	assert.NotNil(server, "starting up - error setting up http server: server is nil")

	return &HttpServer{
		internal: server,
	}
}

func (server *HttpServer) Run(ctx context.Context) error {
	assert.NotNil(ctx, "http server - error starting: context is nil")

	server.ctx = ctx
	log.Info(fmt.Sprintf("starting up - starting http server: %s", server.internal.Addr))

	if err := server.internal.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Error(fmt.Sprintf("starting up - starting http server error: %s", err.Error()))
		return err
	}
	return nil
}

func (server *HttpServer) Stop(ctx context.Context) error {
	assert.NotNil(ctx, "http server - error shutting down: context is nil")

	log.Info("shutting down - stopping http server")
	if err := server.internal.Shutdown(ctx); err != nil {
		log.Error(fmt.Sprintf("shutting down - forced shutdown: %s", err.Error()))
		return err
	}
	log.Debug("shutting down - http server stopped")
	return nil
}
