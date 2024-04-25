package datasource

import (
	"context"

	"gorm.io/gorm"
)

var (
	_ DatasourceContext  = (*DefaultDatasourceContext)(nil)
	_ Datasource         = (*DefaultDatasource)(nil)
	_ TransactionHandler = (*DefaultTransactionHandler)(nil)
)

type DatasourceContext interface {
	GetUrl() string
	GetServer() string
	GetService() string
}

//

type Datasource interface {
	GetDatabase() (*gorm.DB, error)
}

//

type TransactionCtxKey struct{}

type TransactionHandlerFunction func(ctx context.Context, tx *gorm.DB) error

type TransactionHandler interface {
	HandleTransaction(ctx context.Context, fn TransactionHandlerFunction) error
}
