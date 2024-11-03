package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/mongo"
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
	User() string
	Password() string
	Server() any
	Service() string
}

type Connection interface {
	Connect(ctx context.Context) (*mongo.Client, error)
	Close(ctx context.Context)
	Context() Context
}

//

type TransactionCtxKey struct{}

type TransactionHandlerFn func(client *mongo.Client, sesctx context.Context) (any, error)

type TransactionHandler interface {
	HandleTransaction(ctx context.Context, fn TransactionHandlerFn) (any, error)
}
