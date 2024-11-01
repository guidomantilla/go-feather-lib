package security

import (
	"net/http"

	"github.com/guidomantilla/go-feather-lib/pkg/common/assert"
	"github.com/guidomantilla/go-feather-lib/pkg/common/rest"
)

type DefaultAuthenticationEndpoint struct {
	authenticationService AuthenticationService
}

func NewDefaultAuthenticationEndpoint(authenticationService AuthenticationService) *DefaultAuthenticationEndpoint {
	assert.NotNil(authenticationService, "starting up - error setting up authenticationEndpoint: authenticationService is nil")

	return &DefaultAuthenticationEndpoint{
		authenticationService: authenticationService,
	}
}

func (endpoint *DefaultAuthenticationEndpoint) Authenticate(response http.ResponseWriter, request *http.Request) {
	assert.NotNil(response, "authentication endpoint - error authenticating: response is nil")
	assert.NotNil(request, "authentication endpoint - error authenticating: request is nil")

	ctx := request.Context()
	value := ctx.Value(rest.ContextKey{})
	assert.NotNil(value, "authentication endpoint - error authenticating: rest context is nil")
	restCtx := value.(rest.Context)

	var err error
	var principal *Principal
	if err = restCtx.ShouldBindJSON(&principal); err != nil {
		ex := rest.BadRequestException("error unmarshalling request json to object")
		restCtx.AbortWithStatusJSON(ex.Code, ex)
		return
	}

	if errs := endpoint.authenticationService.Validate(principal); errs != nil {
		ex := rest.BadRequestException("error validating the principal", errs...)
		restCtx.AbortWithStatusJSON(ex.Code, ex)
		return
	}

	if err = endpoint.authenticationService.Authenticate(ctx, principal); err != nil {
		ex := rest.UnauthorizedException(err.Error())
		restCtx.AbortWithStatusJSON(ex.Code, ex)
		return
	}

	restCtx.JSON(http.StatusOK, principal)
}
