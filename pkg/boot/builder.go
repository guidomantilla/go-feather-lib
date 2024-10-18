package boot

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	slogGorm "github.com/orandin/slog-gorm"
	sloggin "github.com/samber/slog-gin"
	"google.golang.org/grpc"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/guidomantilla/go-feather-lib/pkg/common/environment"
	"github.com/guidomantilla/go-feather-lib/pkg/common/log"
	"github.com/guidomantilla/go-feather-lib/pkg/datasource"
	"github.com/guidomantilla/go-feather-lib/pkg/rest"
	"github.com/guidomantilla/go-feather-lib/pkg/security"
)

type EnvironmentBuilderFunc func(appCtx *ApplicationContext) environment.Environment

type ConfigLoaderFunc func(appCtx *ApplicationContext)

type DatasourceContextBuilderFunc func(appCtx *ApplicationContext) datasource.Context

type DatasourceConnectionBuilderFunc func(appCtx *ApplicationContext) datasource.Connection[*gorm.DB]

type DatasourceTransactionHandlerBuilderFunc func(appCtx *ApplicationContext) datasource.TransactionHandler[*gorm.DB]

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
		log.Fatal("starting up - error setting up builder: context is nil")
	}

	return &BeanBuilder{

		Environment: func(appCtx *ApplicationContext) environment.Environment {
			return environment.New(environment.WithSSL(), environment.With(appCtx.CmdArgs))
		},
		Config: func(appCtx *ApplicationContext) {
			log.Warn("starting up - warning setting up configuration: config function not implemented")
		},
		DatasourceContext: func(appCtx *ApplicationContext) datasource.Context {
			if !appCtx.Enablers.DatabaseEnabled {
				return nil
			}

			if appCtx.DatabaseConfig != nil {
				return datasource.NewContext(*appCtx.DatabaseConfig.DatasourceUrl, *appCtx.DatabaseConfig.DatasourceUsername, *appCtx.DatabaseConfig.DatasourcePassword, *appCtx.DatabaseConfig.DatasourceServer, *appCtx.DatabaseConfig.DatasourceService)
			}

			log.Fatal("starting up - error setting up configuration: database config is nil")
			return nil
		},
		DatasourceConnection: func(appCtx *ApplicationContext) datasource.Connection[*gorm.DB] {
			if !appCtx.Enablers.DatabaseEnabled {
				return nil
			}

			if appCtx.DatabaseConfig != nil {
				config := &gorm.Config{
					SkipDefaultTransaction: true,
					Logger:                 slogGorm.New(slogGorm.WithHandler(log.AsSlogLogger().Handler()), slogGorm.WithTraceAll(), slogGorm.WithRecordNotFoundError()),
				}
				//TODO: create a factory function for enabling different database types not only: mysql.Open
				return datasource.NewConnection(appCtx.DatasourceContext, mysql.Open(appCtx.DatasourceContext.Url()), config)
			}

			log.Fatal("starting up - error setting up configuration: database config is nil")
			return nil
		},
		DatasourceTransactionHandler: func(appCtx *ApplicationContext) datasource.TransactionHandler[*gorm.DB] {
			if !appCtx.Enablers.DatabaseEnabled {
				return nil
			}

			if appCtx.DatabaseConfig != nil {
				return datasource.NewOrmTransactionHandler(appCtx.DatasourceConnection)
			}

			log.Fatal("starting up - error setting up configuration: database config is nil")
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

			log.Fatal("starting up - error setting up configuration: database config is nil")
			return nil
		},
		TokenManager: func(appCtx *ApplicationContext) security.TokenManager {
			options := security.JwtTokenManagerOptionsChainBuilder().WithIssuer(appCtx.AppName).
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

			recoveryFilter := gin.Recovery()
			loggerFilter := sloggin.New(log.AsSlogLogger().WithGroup("http"))
			customFilter := func(ctx *gin.Context) {
				security.AddApplicationToContext(ctx, appCtx.AppName)
				ctx.Next()
			}

			engine := gin.New()
			engine.Use(loggerFilter, recoveryFilter, customFilter)
			engine.POST("/login", appCtx.AuthenticationEndpoint.Authenticate)
			engine.GET("/health", func(ctx *gin.Context) {
				ctx.JSON(http.StatusOK, gin.H{"status": "alive"})
			})
			engine.NoRoute(func(c *gin.Context) {
				c.JSON(http.StatusNotFound, rest.NotFoundException("resource not found"))
			})
			engine.GET("/info", func(ctx *gin.Context) {
				ctx.JSON(http.StatusOK, gin.H{"appName": appCtx.AppName})
			})

			return engine, engine.Group("/api", appCtx.AuthorizationFilter.Authorize)
		},
		GrpcServer: func(appCtx *ApplicationContext) (*grpc.ServiceDesc, any) {
			if !appCtx.Enablers.GrpcServerEnabled {
				return nil, nil
			}
			log.Fatal("starting up - error setting up grpc configuration: grpc server function not implemented")
			return nil, nil
		},
	}
}
