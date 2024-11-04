package goredis

import (
	"context"

	redis "github.com/redis/go-redis/v9"
)

var (
	_ Context    = (*context_)(nil)
	_ Connection = (*connection)(nil)
	_ Context    = (*MockContext)(nil)
	_ Connection = (*MockConnection)(nil)
)

type Context interface {
	Url() string
	User() string
	Password() string
	Server() any
	Service() string
}

type Connection interface {
	Connect(ctx context.Context) (redis.UniversalClient, error)
	Close(ctx context.Context)
	Context() Context
	Set(key string, value any)
}
