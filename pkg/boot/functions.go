package boot

import (
	"fmt"
	"net"
	"net/http"
	"strings"
	"syscall"

	"github.com/qmdx00/lifecycle"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	log2 "github.com/guidomantilla/go-feather-lib/pkg/common/log"
	"github.com/guidomantilla/go-feather-lib/pkg/server"
)

type InitDelegateFunc func(ctx ApplicationContext) error

func Init(appName string, version string, args []string, logger log2.Logger, enablers *Enablers, builder *BeanBuilder, fn InitDelegateFunc) error {

	if appName == "" {
		log2.Fatal("starting up - error setting up the application: appName is empty")
	}

	if args == nil {
		log2.Fatal("starting up - error setting up the application: args is nil")
	}

	//TODO: Check if it is actually needed this logger reference
	if logger == nil {
		log2.Fatal("starting up - error setting up the application: logger is nil")
	}
	if enablers == nil {
		log2.Warn("starting up - warning setting up the application: http server, grpc server and database connectivity are disabled")
		enablers = &Enablers{}
	}

	if builder == nil {
		log2.Fatal("starting up - error setting up the application: builder is nil")
	}

	if fn == nil {
		log2.Fatal("starting up - error setting up the application: fn is nil")
	}

	app := lifecycle.NewApp(
		lifecycle.WithName(appName),
		lifecycle.WithVersion(version),
		lifecycle.WithSignal(syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGKILL),
	)

	ctx := NewApplicationContext(appName, version, args, logger, enablers, builder)
	defer ctx.Stop()

	if err := fn(*ctx); err != nil {
		log2.Fatal(fmt.Sprintf("starting up - error setting up the application: %s", err.Error()))
	}

	if ctx.Enablers.HttpServerEnabled {
		if ctx.PublicRouter == nil || ctx.HttpConfig == nil || ctx.HttpConfig.Host == nil || ctx.HttpConfig.Port == nil {
			log2.Fatal("starting up - error setting up the application: http server is enabled but no public router or http config is provided")
		}
		httpServer := &http.Server{
			Addr:              net.JoinHostPort(*ctx.HttpConfig.Host, *ctx.HttpConfig.Port),
			Handler:           ctx.PublicRouter,
			ReadHeaderTimeout: 60000,
		}
		app.Attach("HttpServer", server.BuildHttpServer(httpServer))
	}

	if ctx.Enablers.GrpcServerEnabled {
		if ctx.GrpcServiceDesc == nil || ctx.GrpcServiceServer == nil || ctx.GrpcConfig == nil || ctx.GrpcConfig.Host == nil || ctx.GrpcConfig.Port == nil {
			log2.Fatal("starting up - error setting up the application: grpc server is enabled but no grpc service descriptor, grpc service server or grpc config is provided")
		}
		srv := grpc.NewServer()
		srv.RegisterService(ctx.GrpcServiceDesc, ctx.GrpcServiceServer)
		reflection.Register(srv)
		app.Attach("GrpcServer", server.BuildGrpcServer(net.JoinHostPort(*ctx.GrpcConfig.Host, *ctx.GrpcConfig.Port), srv))
	}

	log2.Info(fmt.Sprintf("Application %s started", strings.Join([]string{appName, version}, " - ")))
	return app.Run()
}
