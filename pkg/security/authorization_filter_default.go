package security

import (
	"context"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/guidomantilla/go-feather-lib/pkg/common/log"
	"github.com/guidomantilla/go-feather-lib/pkg/rest"
)

type DefaultAuthorizationFilter struct {
	authorizationService AuthorizationService
}

func NewDefaultAuthorizationFilter(authorizationService AuthorizationService) *DefaultAuthorizationFilter {

	if authorizationService == nil {
		log.Fatal("starting up - error setting up authorizationFilter: authorizationService is nil")
	}

	return &DefaultAuthorizationFilter{
		authorizationService: authorizationService,
	}
}

func (filter *DefaultAuthorizationFilter) Authorize(ctx *gin.Context) {

	header := ctx.Request.Header.Get("Authorization")
	if !strings.HasPrefix(header, "Bearer ") {
		ex := rest.UnauthorizedException("invalid authorization header")
		ctx.AbortWithStatusJSON(ex.Code, ex)
		return
	}

	splits := strings.Split(header, " ")
	if len(splits) != 2 {
		ex := rest.UnauthorizedException("invalid authorization header")
		ctx.AbortWithStatusJSON(ex.Code, ex)
		return
	}
	token := splits[1]

	application, exists := GetApplicationFromContext(ctx)
	if !exists {
		ex := rest.NotFoundException("application name not found in context")
		ctx.AbortWithStatusJSON(ex.Code, ex)
		return
	}
	resource := []string{application, ctx.Request.Method, ctx.FullPath()}

	var err error
	var principal *Principal
	ctxWithResource := context.WithValue(ctx.Request.Context(), ResourceCtxKey{}, strings.Join(resource, " "))
	if principal, err = filter.authorizationService.Authorize(ctxWithResource, token); err != nil {
		ex := rest.UnauthorizedException(err.Error())
		ctx.AbortWithStatusJSON(ex.Code, ex)
		return
	}

	AddPrincipalToContext(ctx, principal)
	ctx.Next()
}
