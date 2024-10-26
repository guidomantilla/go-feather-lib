package gorm

import (
	"context"

	"gorm.io/gorm"
)

var (
	_ Context            = (*context_)(nil)
	_ Connection         = (*connection)(nil)
	_ TransactionHandler = (*transactionHandler)(nil)
	_ Context            = (*MockContext)(nil)
	_ Connection         = (*MockConnection)(nil)
	_ TransactionHandler = (*MockTransactionHandler)(nil)
)

type OpenFn func(dsn string) gorm.Dialector

type Context interface {
	Url() string
	Server() string
	Service() string
}

type Connection interface {
	Connect() (*gorm.DB, error)
	Close()
	Context() Context
}

//

type TransactionCtxKey struct{}

type TransactionHandlerFn func(ctx context.Context, tx *gorm.DB) error

type TransactionHandler interface {
	HandleTransaction(ctx context.Context, fn TransactionHandlerFn) error
}
