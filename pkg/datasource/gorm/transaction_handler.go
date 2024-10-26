package gorm

import (
	"context"

	"gorm.io/gorm"

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

func (handler *transactionHandler) HandleTransaction(ctx context.Context, fn TransactionHandlerFn) error {
	assert.NotNil(ctx, "transaction handler - error handling transaction: context is nil")
	assert.NotNil(fn, "transaction handler - error handling transaction: transaction handler function is nil")

	dbx, err := handler.connection.Connect()
	if err != nil {
		log.Error(err.Error())
		return err
	}
	return dbx.Transaction(func(tx *gorm.DB) error {
		return fn(ctx, tx)
	})
}
