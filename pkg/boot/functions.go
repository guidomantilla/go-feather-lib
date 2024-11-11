package boot

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"strings"
	"syscall"

	"github.com/qmdx00/lifecycle"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	log "github.com/guidomantilla/go-feather-lib/pkg/common/log"
	"github.com/guidomantilla/go-feather-lib/pkg/server"
)

type InitDelegateFunc func(ctx ApplicationContext) error

func Init(ctx context.Context, appName string, version string, args []string, enablers *Enablers, builder *BeanBuilder, fn InitDelegateFunc) error {

	log.Info(ctx, fmt.Sprintf("Application %s", strings.Join([]string{appName, version}, " - ")))

	if appName == "" {
		log.Fatal(ctx, "starting up - error setting up the application: appName is empty")
	}

	if args == nil {
		log.Warn(ctx, "starting up - warning setting up the application: args is nil")
		args = make([]string, 0)
	}

	if enablers == nil {
		log.Warn(ctx, "starting up - warning setting up the application: http server, grpc server and database connectivity are disabled")
		enablers = &Enablers{}
	}

	if builder == nil {
		log.Fatal(ctx, "starting up - error setting up the application: builder is nil")
	}

	if fn == nil {
		log.Fatal(ctx, "starting up - error setting up the application: fn is nil")
	}

	app := lifecycle.NewApp(
		lifecycle.WithName(appName),
		lifecycle.WithVersion(version),
		lifecycle.WithSignal(syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGKILL),
	)

	actx := NewApplicationContext(appName, version, args, enablers, builder)
	defer actx.Stop(ctx)

	if err := fn(*actx); err != nil {
		log.Fatal(ctx, fmt.Sprintf("starting up - error setting up the application: %s", err.Error()))
	}

	if actx.Enablers.HttpServerEnabled {
		if actx.PublicRouter == nil || actx.HttpConfig == nil || actx.HttpConfig.Host == nil || actx.HttpConfig.Port == nil {
			log.Fatal(ctx, "starting up - error setting up the application: http server is enabled but no public router or http config is provided")
		}
		httpServer := &http.Server{
			Addr:              net.JoinHostPort(*actx.HttpConfig.Host, *actx.HttpConfig.Port),
			Handler:           actx.PublicRouter,
			ReadHeaderTimeout: 60000,
		}
		app.Attach(server.BuildHttpServer(httpServer))
	}

	if actx.Enablers.GrpcServerEnabled {
		if actx.GrpcServiceDesc == nil || actx.GrpcServiceServer == nil || actx.GrpcConfig == nil || actx.GrpcConfig.Host == nil || actx.GrpcConfig.Port == nil {
			log.Fatal(ctx, "starting up - error setting up the application: grpc server is enabled but no grpc service descriptor, grpc service server or grpc config is provided")
		}
		srv := grpc.NewServer()
		srv.RegisterService(actx.GrpcServiceDesc, actx.GrpcServiceServer)
		reflection.Register(srv)
		app.Attach(server.BuildGrpcServer(net.JoinHostPort(*actx.GrpcConfig.Host, *actx.GrpcConfig.Port), srv))
	}

	log.Info(ctx, fmt.Sprintf("Application %s started", strings.Join([]string{appName, version}, " - ")))
	return app.Run()
}
