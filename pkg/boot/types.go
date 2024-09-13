package boot

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/xorcare/pointer"
	"google.golang.org/grpc"
	"gorm.io/gorm"

	"github.com/guidomantilla/go-feather-lib/pkg/common/datasource"
	"github.com/guidomantilla/go-feather-lib/pkg/common/environment"
	log2 "github.com/guidomantilla/go-feather-lib/pkg/common/log"
	"github.com/guidomantilla/go-feather-lib/pkg/security"
)

type Enablers struct {
	HttpServerEnabled bool
	GrpcServerEnabled bool
	DatabaseEnabled   bool
}

type HttpConfig struct {
	Host            *string
	Port            *string
	SwaggerPort     *string
	CorsAllowOrigin *string
}

type GrpcConfig struct {
	Host *string
	Port *string
}

type SecurityConfig struct {
	TokenSignatureKey       *string
	TokenVerificationKey    *string
	PasswordMinSpecialChars *string
	PasswordMinNumber       *string
	PasswordMinUpperCase    *string
	PasswordLength          *string
}

type DatabaseConfig struct {
	DatasourceUrl      *string
	DatasourceUsername *string
	DatasourcePassword *string
	DatasourceServer   *string
	DatasourceService  *string
}

type ApplicationContext struct {
	AppName                string
	AppVersion             string
	LogLevel               string
	CmdArgs                []string
	Enablers               *Enablers
	HttpConfig             *HttpConfig
	GrpcConfig             *GrpcConfig
	SecurityConfig         *SecurityConfig
	DatabaseConfig         *DatabaseConfig
	Logger                 log2.Logger
	Environment            environment.Environment
	DatasourceContext      datasource.DatasourceContext
	Datasource             datasource.Datasource
	TransactionHandler     datasource.TransactionHandler
	PasswordEncoder        security.PasswordEncoder
	PasswordGenerator      security.PasswordGenerator
	PasswordManager        security.PasswordManager
	PrincipalManager       security.PrincipalManager
	TokenManager           security.TokenManager
	AuthenticationService  security.AuthenticationService
	AuthenticationEndpoint security.AuthenticationEndpoint
	AuthorizationService   security.AuthorizationService
	AuthorizationFilter    security.AuthorizationFilter
	PublicRouter           *gin.Engine
	PrivateRouter          *gin.RouterGroup
	GrpcServiceDesc        *grpc.ServiceDesc
	GrpcServiceServer      any
}

func NewApplicationContext(appName string, version string, args []string, logger log2.Logger, enablers *Enablers, builder *BeanBuilder) *ApplicationContext {

	if appName == "" {
		log2.Fatal("starting up - error setting up the ApplicationContext: appName is empty")
	}

	if version == "" {
		log2.Fatal("starting up - error setting up the ApplicationContext: version is empty")
	}

	if args == nil {
		log2.Fatal("starting up - error setting up the ApplicationContext: args is nil")
	}

	if logger == nil {
		log2.Fatal("starting up - error setting up the application: logger is nil")
	}

	if enablers == nil {
		log2.Warn("starting up - warning setting up the application: http server, grpc server and database connectivity are disabled")
		enablers = &Enablers{}
	}

	if builder == nil { //nolint:staticcheck
		log2.Fatal("starting up - error setting up the ApplicationContext: builder is nil")
	}

	ctx := &ApplicationContext{
		AppName:    appName,
		AppVersion: version,
		CmdArgs:    args,
		Logger:     logger,
		Enablers:   enablers,
		SecurityConfig: &SecurityConfig{
			TokenSignatureKey:    pointer.Of("SecretYouShouldHide"),
			TokenVerificationKey: pointer.Of("SecretYouShouldHide"),
		},
		HttpConfig: &HttpConfig{
			Host: pointer.Of("localhost"),
			Port: pointer.Of("8080"),
		},
		GrpcConfig: &GrpcConfig{
			Host: pointer.Of("localhost"),
			Port: pointer.Of("50051"),
		},
	}

	log2.Debug("starting up - setting up environment variables")
	ctx.Environment = builder.Environment(ctx) //nolint:staticcheck

	log2.Debug("starting up - setting up configuration")
	builder.Config(ctx) //nolint:staticcheck

	if ctx.Enablers.DatabaseEnabled {
		log2.Debug("starting up - setting up db connectivity")
		ctx.DatasourceContext = builder.DatasourceContext(ctx)   //nolint:staticcheck
		ctx.Datasource = builder.Datasource(ctx)                 //nolint:staticcheck
		ctx.TransactionHandler = builder.TransactionHandler(ctx) //nolint:staticcheck
	} else {
		log2.Warn("starting up - warning setting up database configuration. database connectivity is disabled")
	}

	log2.Debug("starting up - setting up security")
	ctx.PasswordEncoder = builder.PasswordEncoder(ctx)                                                                          //nolint:staticcheck
	ctx.PasswordGenerator = builder.PasswordGenerator(ctx)                                                                      //nolint:staticcheck
	ctx.PasswordManager = builder.PasswordManager(ctx)                                                                          //nolint:staticcheck
	ctx.PrincipalManager, ctx.TokenManager = builder.PrincipalManager(ctx), builder.TokenManager(ctx)                           //nolint:staticcheck
	ctx.AuthenticationService, ctx.AuthorizationService = builder.AuthenticationService(ctx), builder.AuthorizationService(ctx) //nolint:staticcheck
	ctx.AuthenticationEndpoint, ctx.AuthorizationFilter = builder.AuthenticationEndpoint(ctx), builder.AuthorizationFilter(ctx) //nolint:staticcheck

	if ctx.Enablers.HttpServerEnabled {
		log2.Debug("starting up - setting up http server")
		ctx.PublicRouter, ctx.PrivateRouter = builder.HttpServer(ctx) //nolint:staticcheck
	} else {
		log2.Warn("starting up - warning setting up http configuration. http server is disabled")
	}

	if ctx.Enablers.GrpcServerEnabled {
		log2.Debug("starting up - setting up grpc server")
		ctx.GrpcServiceDesc, ctx.GrpcServiceServer = builder.GrpcServer(ctx) //nolint:staticcheck
	} else {
		log2.Warn("starting up - warning setting up grpc configuration. grpc server is disabled")
	}

	return ctx
}

func (ctx *ApplicationContext) Stop() {

	var err error

	if ctx.Datasource != nil && ctx.DatasourceContext != nil {

		var database *gorm.DB
		log2.Debug("shutting down - closing up db connection")

		if database, err = ctx.Datasource.GetDatabase(); err != nil {
			log2.Error(fmt.Sprintf("shutting down - error db connection: %s", err.Error()))
			return
		}

		var db *sql.DB
		if db, err = database.DB(); err != nil {
			log2.Error(fmt.Sprintf("shutting down - error db connection: %s", err.Error()))
			return
		}

		if err = db.Close(); err != nil {
			log2.Error(fmt.Sprintf("shutting down - error closing db connection: %s", err.Error()))
			return
		}

		log2.Debug("shutting down - db connection closed")
	}

	log2.Info(fmt.Sprintf("Application %s stopped", strings.Join([]string{ctx.AppName, ctx.AppVersion}, " - ")))
}
