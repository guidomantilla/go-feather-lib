package server

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"strings"

	"google.golang.org/grpc"

	"github.com/guidomantilla/go-feather-lib/pkg/common/log"
)

type GrpcServer struct {
	ctx      context.Context
	address  string
	internal *grpc.Server
}

func BuildGrpcServer(address string, server *grpc.Server) Server {

	if strings.TrimSpace(address) == "" {
		log.Fatal("starting up - error setting up grpc server: address is empty")
	}

	if server == nil {
		log.Fatal("starting up - error setting up grpc server: server is nil")
	}

	return &GrpcServer{
		address:  address,
		internal: server,
	}
}

func (server *GrpcServer) Run(ctx context.Context) error {

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

func (server *GrpcServer) Stop(_ context.Context) error {

	log.Info("shutting down - stopping grpc server")
	server.internal.GracefulStop()
	log.Debug("shutting down - grpc server stopped")
	return nil
}
