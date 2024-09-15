package datasource

import (
	"context"

	"gorm.io/gorm"
)

var (
	_ DatasourceContext  = (*DefaultDatasourceContext)(nil)
	_ Datasource         = (*DefaultDatasource)(nil)
	_ TransactionHandler = (*DefaultTransactionHandler)(nil)
	_ DatasourceContext  = (*MockDatasourceContext)(nil)
	_ Datasource         = (*MockDatasource)(nil)
	_ TransactionHandler = (*MockTransactionHandler)(nil)
)

type DatasourceContext interface {
	Url() string
	Server() string
	Service() string
}

//

type Datasource interface {
	Connect() (*gorm.DB, error)
	Close()
	DatasourceContext() DatasourceContext
}

//

type TransactionCtxKey struct{}

type TransactionHandlerFunction func(ctx context.Context, tx *gorm.DB) error

type TransactionHandler interface {
	HandleTransaction(ctx context.Context, fn TransactionHandlerFunction) error
}
