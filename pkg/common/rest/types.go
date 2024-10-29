package rest

import "net/http"

type Context interface {
	Request() *http.Request
	FullPath() string
	Set(key string, value any)
	Get(key string) (value any, exists bool)
	ShouldBindJSON(obj any) error
	AbortWithStatusJSON(code int, jsonObj any)
	JSON(code int, obj any)
	Next()
}
