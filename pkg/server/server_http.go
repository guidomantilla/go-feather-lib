package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/guidomantilla/go-feather-lib/pkg/common/assert"
	"github.com/guidomantilla/go-feather-lib/pkg/common/log"
)

type httpServer struct {
	ctx      context.Context
	internal *http.Server
}

func NewHttpServer(server *http.Server) *httpServer {
	assert.NotNil(server, "starting up - error setting up http server: server is nil")

	return &httpServer{
		internal: server,
	}
}

func (server *httpServer) Run(ctx context.Context) error {
	assert.NotNil(ctx, "http server - error starting: context is nil")

	server.ctx = ctx
	log.Info(ctx, fmt.Sprintf("starting up - starting http server: %s", server.internal.Addr))

	if err := server.internal.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Error(ctx, fmt.Sprintf("starting up - starting http server error: %s", err.Error()))
		return err
	}
	return nil
}

func (server *httpServer) Stop(ctx context.Context) error {
	assert.NotNil(ctx, "http server - error shutting down: context is nil")

	log.Info(ctx, "shutting down - stopping http server")
	if err := server.internal.Shutdown(ctx); err != nil {
		log.Error(ctx, fmt.Sprintf("shutting down - forced shutdown: %s", err.Error()))
		return err
	}
	log.Debug(ctx, "shutting down - http server stopped")
	return nil
}
