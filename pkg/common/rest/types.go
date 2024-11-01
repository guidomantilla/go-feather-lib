package rest

type ContextKey struct {
}

type Context interface {
	FullPath() string
	Set(key string, value any)
	Get(key string) (value any, exists bool)
	ShouldBindJSON(obj any) error
	AbortWithStatusJSON(code int, jsonObj any)
	JSON(code int, obj any)
	Next()
}
