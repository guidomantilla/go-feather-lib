package datasource

import (
	"context"

	"gorm.io/gorm"
)

var (
	_ StoreContext                      = (*OrmContext)(nil)
	_ StoreConnection[*gorm.DB]         = (*OrmConnection)(nil)
	_ StoreTransactionHandler[*gorm.DB] = (*OrmTransactionHandler)(nil)
	_ StoreContext                      = (*MockStoreContext)(nil)
	_ StoreConnection[*gorm.DB]         = (*MockStoreConnection[*gorm.DB])(nil)
	_ StoreTransactionHandler[*gorm.DB] = (*MockStoreTransactionHandler[*gorm.DB])(nil)
)

type StoreContext interface {
	Url() string
	Server() string
	Service() string
}

//

type StoreConnectionTypes interface {
	*gorm.DB | struct{}
}

//type StoreConnectionTypes = *gorm.DB

type StoreConnection[T StoreConnectionTypes] interface {
	Connect() (T, error)
	Close()
	Context() StoreContext
}

//

type StoreTransactionCtxKey struct{}

type StoreTransactionHandlerFn[T StoreConnectionTypes] func(ctx context.Context, tx T) error

type StoreTransactionHandler[T StoreConnectionTypes] interface {
	HandleTransaction(ctx context.Context, fn StoreTransactionHandlerFn[T]) error
}
