package mongo

import (
	"context"

	"github.com/guidomantilla/go-feather-lib/pkg/common/assert"
	"github.com/guidomantilla/go-feather-lib/pkg/common/log"
)

type transactionHandler struct {
	connection Connection
}

func NewTransactionHandler(connection Connection) TransactionHandler {
	assert.NotNil(connection, "starting up - error setting up orm transaction handler: connection is nil")

	return &transactionHandler{connection: connection}
}

func (handler *transactionHandler) HandleTransaction(ctx context.Context, fn TransactionHandlerFn) (any, error) {
	assert.NotNil(ctx, "transaction handler - error handling transaction: context is nil")
	assert.NotNil(fn, "transaction handler - error handling transaction: transaction handler function is nil")

	dbx, err := handler.connection.Connect(ctx)
	if err != nil {
		log.Error(ctx, err.Error())
		return nil, err
	}

	session, err := dbx.StartSession()
	if err != nil {
		log.Error(ctx, err.Error())
		return nil, err
	}
	defer session.EndSession(ctx)

	callback := func(sesctx context.Context) (any, error) {
		return fn(dbx, sesctx)
	}

	return session.WithTransaction(ctx, callback)
}
