package security

import (
	"github.com/gin-gonic/gin"
)

const (
	ApplicationCtxKey = "application"
	PrincipalCtxKey   = "principal"
)

func AddPrincipalToContext(ctx *gin.Context, principal *Principal) {
	ctx.Set(PrincipalCtxKey, principal)
}

func AddApplicationToContext(ctx *gin.Context, application string) {
	ctx.Set(ApplicationCtxKey, application)
}

func GetPrincipalFromContext(ctx *gin.Context) (*Principal, bool) {
	var exists bool
	var value any
	if value, exists = ctx.Get(PrincipalCtxKey); !exists {
		return nil, false
	}
	return value.(*Principal), true
}

func GetApplicationFromContext(ctx *gin.Context) (string, bool) {
	var exists bool
	var value any
	if value, exists = ctx.Get(ApplicationCtxKey); !exists {
		return "", false
	}
	return value.(string), true
}
