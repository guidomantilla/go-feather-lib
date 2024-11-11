package boot

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"gorm.io/gorm"

	"github.com/guidomantilla/go-feather-lib/pkg/common/environment"
	"github.com/guidomantilla/go-feather-lib/pkg/common/log"
	"github.com/guidomantilla/go-feather-lib/pkg/common/rest"
	dgorm "github.com/guidomantilla/go-feather-lib/pkg/datasource/gorm"
	"github.com/guidomantilla/go-feather-lib/pkg/security"
	"github.com/guidomantilla/go-feather-lib/pkg/web"
)

type EnvironmentBuilderFunc func(appCtx *ApplicationContext) environment.Environment

type ConfigLoaderFunc func(appCtx *ApplicationContext)

type DatasourceOpenFunc func(appCtx *ApplicationContext) dgorm.OpenFn

type DatasourceContextBuilderFunc func(appCtx *ApplicationContext) dgorm.Context

type DatasourceConnectionBuilderFunc func(appCtx *ApplicationContext) dgorm.Connection

type DatasourceTransactionHandlerBuilderFunc func(appCtx *ApplicationContext) dgorm.TransactionHandler

type PasswordGeneratorBuilderFunc func(appCtx *ApplicationContext) security.PasswordGenerator

type PasswordEncoderBuilderFunc func(appCtx *ApplicationContext) security.PasswordEncoder

type PasswordManagerBuilderFunc func(appCtx *ApplicationContext) security.PasswordManager

type PrincipalManagerBuilderFunc func(appCtx *ApplicationContext) security.PrincipalManager

type TokenManagerBuilderFunc func(appCtx *ApplicationContext) security.TokenManager

type AuthenticationServiceBuilderFunc func(appCtx *ApplicationContext) security.AuthenticationService

type AuthorizationServiceBuilderFunc func(appCtx *ApplicationContext) security.AuthorizationService

type AuthenticationEndpointBuilderFunc func(appCtx *ApplicationContext) security.AuthenticationEndpoint

type AuthorizationFilterBuilderFunc func(appCtx *ApplicationContext) security.AuthorizationFilter

type HttpServerBuilderFunc func(appCtx *ApplicationContext) (*gin.Engine, *gin.RouterGroup)

type GrpcServerBuilderFunc func(appCtx *ApplicationContext) (*grpc.ServiceDesc, any)

type BeanBuilder struct {
	Environment                  EnvironmentBuilderFunc
	Config                       ConfigLoaderFunc
	DatasourceOpenFn             DatasourceOpenFunc
	DatasourceContext            DatasourceContextBuilderFunc
	DatasourceConnection         DatasourceConnectionBuilderFunc
	DatasourceTransactionHandler DatasourceTransactionHandlerBuilderFunc
	PasswordEncoder              PasswordEncoderBuilderFunc
	PasswordGenerator            PasswordGeneratorBuilderFunc
	PasswordManager              PasswordManagerBuilderFunc
	PrincipalManager             PrincipalManagerBuilderFunc
	TokenManager                 TokenManagerBuilderFunc
	AuthenticationService        AuthenticationServiceBuilderFunc
	AuthorizationService         AuthorizationServiceBuilderFunc
	AuthenticationEndpoint       AuthenticationEndpointBuilderFunc
	AuthorizationFilter          AuthorizationFilterBuilderFunc
	HttpServer                   HttpServerBuilderFunc
	GrpcServer                   GrpcServerBuilderFunc
}

func NewBeanBuilder(ctx context.Context) *BeanBuilder {

	if ctx == nil {
		log.Fatal(ctx, "starting up - error setting up builder: context is nil")
	}

	return &BeanBuilder{

		Environment: func(appCtx *ApplicationContext) environment.Environment {
			return environment.New(environment.OptionsBuilder().WithSSL().WithCmd(appCtx.CmdArgs).Build())
		},
		Config: func(appCtx *ApplicationContext) {
			log.Warn(ctx, "starting up - warning setting up configuration: config function not implemented")
		},
		DatasourceOpenFn: func(appCtx *ApplicationContext) dgorm.OpenFn {
			log.Warn(ctx, "starting up - warning setting up configuration: datasource OpenFn function not implemented")
			return nil
		},
		DatasourceContext: func(appCtx *ApplicationContext) dgorm.Context {
			if !appCtx.Enablers.DatabaseEnabled {
				return nil
			}

			if appCtx.DatabaseConfig != nil && appCtx.DatasourceOpenFn != nil {
				return dgorm.NewContext(*appCtx.DatabaseConfig.DatasourceUrl, *appCtx.DatabaseConfig.DatasourceUsername, *appCtx.DatabaseConfig.DatasourcePassword, *appCtx.DatabaseConfig.DatasourceServer, *appCtx.DatabaseConfig.DatasourceService)
			}

			log.Fatal(ctx, "starting up - error setting up configuration: database config or openFn is nil")
			return nil
		},
		DatasourceConnection: func(appCtx *ApplicationContext) dgorm.Connection {
			if !appCtx.Enablers.DatabaseEnabled {
				return nil
			}

			if appCtx.DatabaseConfig != nil && appCtx.DatasourceOpenFn != nil {
				config := &gorm.Config{
					SkipDefaultTransaction: true,
					Logger:                 dgorm.Logger(),
				}
				return dgorm.NewConnection(appCtx.DatasourceContext, appCtx.DatasourceOpenFn, config)
			}

			log.Fatal(ctx, "starting up - error setting up configuration: database config or openFn is nil")
			return nil
		},
		DatasourceTransactionHandler: func(appCtx *ApplicationContext) dgorm.TransactionHandler {
			if !appCtx.Enablers.DatabaseEnabled {
				return nil
			}

			if appCtx.DatabaseConfig != nil && appCtx.DatasourceOpenFn != nil {
				return dgorm.NewTransactionHandler(appCtx.DatasourceConnection)
			}

			log.Fatal(ctx, "starting up - error setting up configuration: database config or openFn is nil")
			return nil
		},
		PasswordEncoder: func(appCtx *ApplicationContext) security.PasswordEncoder {
			return security.NewBcryptPasswordEncoder()
		},
		PasswordGenerator: func(appCtx *ApplicationContext) security.PasswordGenerator {
			return security.NewPasswordGenerator()
		},
		PasswordManager: func(appCtx *ApplicationContext) security.PasswordManager {
			return security.NewPasswordManager(appCtx.PasswordEncoder, appCtx.PasswordGenerator)
		},
		PrincipalManager: func(appCtx *ApplicationContext) security.PrincipalManager {
			if !appCtx.Enablers.DatabaseEnabled {
				return security.NewBasePrincipalManager(appCtx.PasswordManager)
			}

			if appCtx.DatabaseConfig != nil {
				return security.NewGormPrincipalManager(appCtx.DatasourceTransactionHandler, appCtx.PasswordManager)
			}

			log.Fatal(ctx, "starting up - error setting up configuration: database config is nil")
			return nil
		},
		TokenManager: func(appCtx *ApplicationContext) security.TokenManager {
			options := security.JwtTokenManagerOptionsBuilder().WithIssuer(appCtx.AppName).
				WithSigningKey([]byte(*appCtx.SecurityConfig.TokenSignatureKey)).
				WithVerifyingKey([]byte(*appCtx.SecurityConfig.TokenVerificationKey)).Build()
			return security.NewJwtTokenManager(options)
		},
		AuthenticationService: func(appCtx *ApplicationContext) security.AuthenticationService {
			return security.NewDefaultAuthenticationService(appCtx.PasswordManager, appCtx.PrincipalManager, appCtx.TokenManager)
		},
		AuthorizationService: func(appCtx *ApplicationContext) security.AuthorizationService {
			return security.NewDefaultAuthorizationService(appCtx.TokenManager, appCtx.PrincipalManager)
		},
		AuthenticationEndpoint: func(appCtx *ApplicationContext) security.AuthenticationEndpoint {
			return security.NewDefaultAuthenticationEndpoint(appCtx.AuthenticationService)
		},
		AuthorizationFilter: func(appCtx *ApplicationContext) security.AuthorizationFilter {
			return security.NewDefaultAuthorizationFilter(appCtx.AuthorizationService)
		},
		HttpServer: func(appCtx *ApplicationContext) (*gin.Engine, *gin.RouterGroup) {
			if !appCtx.Enablers.HttpServerEnabled {
				return nil, nil
			}

			engine := gin.New()
			engine.Use(web.Logger(), gin.Recovery(), func(ctx *gin.Context) {
				security.AddApplicationToContext(ctx, appCtx.AppName)
				ctx.Next()
			})
			engine.POST("/login", func(ctx *gin.Context) {
				gin.WrapF(appCtx.AuthenticationEndpoint.Authenticate)
			})
			engine.GET("/health", func(ctx *gin.Context) {
				ctx.JSON(http.StatusOK, gin.H{"status": "alive"})
			})
			engine.NoRoute(func(c *gin.Context) {
				c.JSON(http.StatusNotFound, rest.NotFoundException("resource not found"))
			})
			engine.GET("/info", func(ctx *gin.Context) {
				ctx.JSON(http.StatusOK, gin.H{"appName": appCtx.AppName})
			})

			return engine, engine.Group("/api", func(ctx *gin.Context) {
				gin.WrapF(appCtx.AuthorizationFilter.Authorize)
			})
		},
		GrpcServer: func(appCtx *ApplicationContext) (*grpc.ServiceDesc, any) {
			if !appCtx.Enablers.GrpcServerEnabled {
				return nil, nil
			}
			log.Fatal(ctx, "starting up - error setting up grpc configuration: grpc server function not implemented")
			return nil, nil
		},
	}
}
