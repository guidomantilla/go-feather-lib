package security

import (
	"context"

	"github.com/gin-gonic/gin"
)

var (
	_ AuthenticationEndpoint = (*DefaultAuthenticationEndpoint)(nil)
	_ AuthenticationService  = (*DefaultAuthenticationService)(nil)
	_ AuthenticationEndpoint = (*MockAuthenticationEndpoint)(nil)
	_ AuthenticationService  = (*MockAuthenticationService)(nil)
)

type AuthenticationEndpoint interface {
	Authenticate(ctx *gin.Context)
}

type AuthenticationService interface {
	Authenticate(ctx context.Context, principal *Principal) error
	Validate(principal *Principal) []error
}
