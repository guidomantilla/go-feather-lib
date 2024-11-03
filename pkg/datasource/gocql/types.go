package gocql

import (
	"context"

	"github.com/gocql/gocql"
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
	Connect(ctx context.Context) (*gocql.Session, error)
	Close(ctx context.Context)
	Context() Context
	Set(key string, value any)
}
