package boot

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/xorcare/pointer"
	"google.golang.org/grpc"
	"gorm.io/gorm"

	"github.com/guidomantilla/go-feather-lib/pkg/common/environment"
	log "github.com/guidomantilla/go-feather-lib/pkg/common/log"
	"github.com/guidomantilla/go-feather-lib/pkg/datasource"
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
	AppName                      string
	AppVersion                   string
	LogLevel                     string
	CmdArgs                      []string
	Enablers                     *Enablers
	HttpConfig                   *HttpConfig
	GrpcConfig                   *GrpcConfig
	SecurityConfig               *SecurityConfig
	DatabaseConfig               *DatabaseConfig
	Environment                  environment.Environment
	DatasourceContext            datasource.Context
	DatasourceConnection         datasource.Connection[*gorm.DB]
	DatasourceTransactionHandler datasource.TransactionHandler[*gorm.DB]
	PasswordEncoder              security.PasswordEncoder
	PasswordGenerator            security.PasswordGenerator
	PasswordManager              security.PasswordManager
	PrincipalManager             security.PrincipalManager
	TokenManager                 security.TokenManager
	AuthenticationService        security.AuthenticationService
	AuthenticationEndpoint       security.AuthenticationEndpoint
	AuthorizationService         security.AuthorizationService
	AuthorizationFilter          security.AuthorizationFilter
	PublicRouter                 *gin.Engine
	PrivateRouter                *gin.RouterGroup
	GrpcServiceDesc              *grpc.ServiceDesc
	GrpcServiceServer            any
}

func NewApplicationContext(appName string, version string, args []string, enablers *Enablers, builder *BeanBuilder) *ApplicationContext {

	if appName == "" {
		log.Fatal("starting up - error setting up the ApplicationContext: appName is empty")
	}

	if version == "" {
		log.Fatal("starting up - error setting up the ApplicationContext: version is empty")
	}

	if args == nil {
		log.Fatal("starting up - error setting up the ApplicationContext: args is nil")
	}

	if enablers == nil {
		log.Warn("starting up - warning setting up the application: http server, grpc server and database connectivity are disabled")
		enablers = &Enablers{}
	}

	if builder == nil { //nolint:staticcheck
		log.Fatal("starting up - error setting up the ApplicationContext: builder is nil")
	}

	ctx := &ApplicationContext{
		AppName:    appName,
		AppVersion: version,
		CmdArgs:    args,
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

	log.Debug("starting up - setting up environment variables")
	ctx.Environment = builder.Environment(ctx) //nolint:staticcheck

	log.Debug("starting up - setting up configuration")
	builder.Config(ctx) //nolint:staticcheck

	if ctx.Enablers.DatabaseEnabled {
		log.Debug("starting up - setting up db connectivity")
		ctx.DatasourceContext = builder.DatasourceContext(ctx)                       //nolint:staticcheck
		ctx.DatasourceConnection = builder.DatasourceConnection(ctx)                 //nolint:staticcheck
		ctx.DatasourceTransactionHandler = builder.DatasourceTransactionHandler(ctx) //nolint:staticcheck
	} else {
		log.Warn("starting up - warning setting up database configuration. database connectivity is disabled")
	}

	log.Debug("starting up - setting up security")
	ctx.PasswordEncoder = builder.PasswordEncoder(ctx)                                                                          //nolint:staticcheck
	ctx.PasswordGenerator = builder.PasswordGenerator(ctx)                                                                      //nolint:staticcheck
	ctx.PasswordManager = builder.PasswordManager(ctx)                                                                          //nolint:staticcheck
	ctx.PrincipalManager, ctx.TokenManager = builder.PrincipalManager(ctx), builder.TokenManager(ctx)                           //nolint:staticcheck
	ctx.AuthenticationService, ctx.AuthorizationService = builder.AuthenticationService(ctx), builder.AuthorizationService(ctx) //nolint:staticcheck
	ctx.AuthenticationEndpoint, ctx.AuthorizationFilter = builder.AuthenticationEndpoint(ctx), builder.AuthorizationFilter(ctx) //nolint:staticcheck

	if ctx.Enablers.HttpServerEnabled {
		log.Debug("starting up - setting up http server")
		ctx.PublicRouter, ctx.PrivateRouter = builder.HttpServer(ctx) //nolint:staticcheck
	} else {
		log.Warn("starting up - warning setting up http configuration. http server is disabled")
	}

	if ctx.Enablers.GrpcServerEnabled {
		log.Debug("starting up - setting up grpc server")
		ctx.GrpcServiceDesc, ctx.GrpcServiceServer = builder.GrpcServer(ctx) //nolint:staticcheck
	} else {
		log.Warn("starting up - warning setting up grpc configuration. grpc server is disabled")
	}

	return ctx
}

func (ctx *ApplicationContext) Stop() {

	var err error

	if ctx.DatasourceConnection != nil && ctx.DatasourceContext != nil {

		var database *gorm.DB
		log.Debug("shutting down - closing up db connection")

		if database, err = ctx.DatasourceConnection.Connect(); err != nil {
			log.Error(fmt.Sprintf("shutting down - error db connection: %s", err.Error()))
			return
		}

		var db *sql.DB
		if db, err = database.DB(); err != nil {
			log.Error(fmt.Sprintf("shutting down - error db connection: %s", err.Error()))
			return
		}

		if err = db.Close(); err != nil {
			log.Error(fmt.Sprintf("shutting down - error closing db connection: %s", err.Error()))
			return
		}

		log.Debug("shutting down - db connection closed")
	}

	log.Info(fmt.Sprintf("Application %s stopped", strings.Join([]string{ctx.AppName, ctx.AppVersion}, " - ")))
}
