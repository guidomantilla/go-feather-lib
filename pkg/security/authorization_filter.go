package security

import (
	"context"
	"net/http"
	"strings"

	"github.com/guidomantilla/go-feather-lib/pkg/common/assert"
	"github.com/guidomantilla/go-feather-lib/pkg/common/rest"
)

type DefaultAuthorizationFilter struct {
	authorizationService AuthorizationService
}

func NewDefaultAuthorizationFilter(authorizationService AuthorizationService) *DefaultAuthorizationFilter {
	assert.NotNil(authorizationService, "starting up - error setting up authorizationFilter: authorizationService is nil")

	return &DefaultAuthorizationFilter{
		authorizationService: authorizationService,
	}
}

func (filter *DefaultAuthorizationFilter) Authorize(response http.ResponseWriter, request *http.Request) {
	assert.NotNil(response, "authorization filter - error authorizing: response is nil")
	assert.NotNil(request, "authorization filter - error authorizing: request is nil")

	ctx := request.Context()
	value := ctx.Value(rest.ContextKey{})
	assert.NotNil(value, "authentication endpoint - error authenticating: rest context is nil")
	restCtx := value.(rest.Context)

	header := request.Header.Get("Authorization")
	if !strings.HasPrefix(header, "Bearer ") {
		ex := rest.UnauthorizedException("invalid authorization header")
		restCtx.AbortWithStatusJSON(ex.Code, ex)
		return
	}

	splits := strings.Split(header, " ")
	if len(splits) != 2 {
		ex := rest.UnauthorizedException("invalid authorization header")
		restCtx.AbortWithStatusJSON(ex.Code, ex)
		return
	}
	token := splits[1]

	application, exists := GetApplicationFromContext(restCtx)
	if !exists {
		ex := rest.NotFoundException("application name not found in context")
		restCtx.AbortWithStatusJSON(ex.Code, ex)
		return
	}
	resource := []string{application, request.Method, restCtx.FullPath()}

	var err error
	var principal *Principal
	ctxWithResource := context.WithValue(request.Context(), ResourceCtxKey{}, strings.Join(resource, " "))
	if principal, err = filter.authorizationService.Authorize(ctxWithResource, token); err != nil {
		ex := rest.UnauthorizedException(err.Error())
		restCtx.AbortWithStatusJSON(ex.Code, ex)
		return
	}

	AddPrincipalToContext(restCtx, principal)
	restCtx.Next()
}
