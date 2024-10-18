package datasource

import (
	"context"

	"gorm.io/gorm"

	"github.com/guidomantilla/go-feather-lib/pkg/common/log"
)

type transactionHandler struct {
	connection Connection[*gorm.DB]
}

func NewOrmTransactionHandler(connection Connection[*gorm.DB]) TransactionHandler[*gorm.DB] {
	return &transactionHandler{connection: connection}
}

func (handler *transactionHandler) HandleTransaction(ctx context.Context, fn TransactionHandlerFn[*gorm.DB]) error {
	dbx, err := handler.connection.Connect()
	if err != nil {
		log.Error(err.Error())
		return err
	}
	return dbx.Transaction(func(tx *gorm.DB) error {
		return fn(ctx, tx)
	})
}
