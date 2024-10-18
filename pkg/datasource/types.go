package datasource

import (
	"context"

	"gorm.io/gorm"
)

var (
	_ Context                      = (*context_)(nil)
	_ Connection[*gorm.DB]         = (*connection)(nil)
	_ TransactionHandler[*gorm.DB] = (*transactionHandler)(nil)
	_ Context                      = (*MockContext)(nil)
	_ Connection[*gorm.DB]         = (*MockConnection[*gorm.DB])(nil)
	_ TransactionHandler[*gorm.DB] = (*MockTransactionHandler[*gorm.DB])(nil)
)

type Context interface {
	Url() string
	Server() string
	Service() string
}

//

type StoreConnectionTypes interface {
	*gorm.DB | struct{}
}

//type StoreConnectionTypes = *gorm.DB

type Connection[T StoreConnectionTypes] interface {
	Connect() (T, error)
	Close()
	Context() Context
}

//

type TransactionCtxKey struct{}

type TransactionHandlerFn[T StoreConnectionTypes] func(ctx context.Context, tx T) error

type TransactionHandler[T StoreConnectionTypes] interface {
	HandleTransaction(ctx context.Context, fn TransactionHandlerFn[T]) error
}
