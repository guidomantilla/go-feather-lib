package security

import (
	"context"

	"github.com/gin-gonic/gin"
)

var (
	_ AuthorizationFilter  = (*DefaultAuthorizationFilter)(nil)
	_ AuthorizationService = (*DefaultAuthorizationService)(nil)
	_ AuthorizationFilter  = (*MockAuthorizationFilter)(nil)
	_ AuthorizationService = (*MockAuthorizationService)(nil)
)

type ResourceCtxKey struct{}

type AuthorizationFilter interface {
	Authorize(ctx *gin.Context)
}

type AuthorizationService interface {
	Authorize(ctx context.Context, tokenString string) (*Principal, error)
}
