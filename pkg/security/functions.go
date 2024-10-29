package security

import (
	"github.com/guidomantilla/go-feather-lib/pkg/common/rest"
)

const (
	ApplicationCtxKey = "application"
	PrincipalCtxKey   = "principal"
)

func AddPrincipalToContext(ctx rest.Context, principal *Principal) {
	ctx.Set(PrincipalCtxKey, principal)
}

func AddApplicationToContext(ctx rest.Context, application string) {
	ctx.Set(ApplicationCtxKey, application)
}

func GetPrincipalFromContext(ctx rest.Context) (*Principal, bool) {
	var exists bool
	var value any
	if value, exists = ctx.Get(PrincipalCtxKey); !exists {
		return nil, false
	}
	return value.(*Principal), true
}

func GetApplicationFromContext(ctx rest.Context) (string, bool) {
	var exists bool
	var value any
	if value, exists = ctx.Get(ApplicationCtxKey); !exists {
		return "", false
	}
	return value.(string), true
}
