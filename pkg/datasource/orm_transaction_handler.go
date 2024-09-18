package datasource

import (
	"context"

	"gorm.io/gorm"

	"github.com/guidomantilla/go-feather-lib/pkg/common/log"
)

type OrmTransactionHandler struct {
	connection StoreConnection[*gorm.DB]
}

func NewOrmTransactionHandler(connection StoreConnection[*gorm.DB]) *OrmTransactionHandler {
	return &OrmTransactionHandler{connection: connection}
}

func (handler *OrmTransactionHandler) HandleTransaction(ctx context.Context, fn StoreTransactionHandlerFn[*gorm.DB]) error {
	dbx, err := handler.connection.Connect()
	if err != nil {
		log.Error(err.Error())
		return err
	}
	return dbx.Transaction(func(tx *gorm.DB) error {
		return fn(ctx, tx)
	})
}
