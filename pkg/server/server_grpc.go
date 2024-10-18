package server

import (
	"context"
	"errors"
	"fmt"
	"github.com/guidomantilla/go-feather-lib/pkg/common/assert"
	"google.golang.org/grpc"
	"net"
	"net/http"

	"github.com/guidomantilla/go-feather-lib/pkg/common/log"
)

type GrpcServer struct {
	ctx      context.Context
	address  string
	internal *grpc.Server
}

func NewGrpcServer(address string, server *grpc.Server) *GrpcServer {
	assert.NotEmpty(address, "starting up - error setting up grpc server: address is empty")
	assert.NotNil(server, "starting up - error setting up grpc server: server is nil")

	return &GrpcServer{
		address:  address,
		internal: server,
	}
}

func (server *GrpcServer) Run(ctx context.Context) error {
	assert.NotNil(ctx, "grpc server - error starting up: context is nil")

	server.ctx = ctx
	log.Info(fmt.Sprintf("starting up - starting grpc server: %s", server.address))

	var err error
	var listener net.Listener
	if listener, err = net.Listen("tcp", server.address); err != nil {
		log.Error(fmt.Sprintf("starting up - starting grpc server error: %s", err.Error()))
		return err
	}

	if err = server.internal.Serve(listener); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Error(fmt.Sprintf("starting up - starting grpc server error: %s", err.Error()))
		return err
	}
	return nil
}

func (server *GrpcServer) Stop(ctx context.Context) error {
	assert.NotNil(ctx, "grpc server - error shutting down: context is nil")

	log.Info("shutting down - stopping grpc server")
	server.internal.GracefulStop()
	log.Debug("shutting down - grpc server stopped")
	return nil
}
