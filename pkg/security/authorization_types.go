package security

import (
	"context"
	"net/http"
)

var (
	_ AuthorizationFilter  = (*DefaultAuthorizationFilter)(nil)
	_ AuthorizationService = (*DefaultAuthorizationService)(nil)
	_ AuthorizationFilter  = (*MockAuthorizationFilter)(nil)
	_ AuthorizationService = (*MockAuthorizationService)(nil)
)

type ResourceCtxKey struct{}

type AuthorizationFilter interface {
	Authorize(response http.ResponseWriter, request *http.Request)
}

type AuthorizationService interface {
	Authorize(ctx context.Context, tokenString string) (*Principal, error)
}
