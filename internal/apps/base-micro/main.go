package main

import (
	"context"
	"fmt"
	"github.com/glebarez/sqlite"
	cserver "github.com/guidomantilla/go-feather-lib/pkg/common/server"
	dgorm "github.com/guidomantilla/go-feather-lib/pkg/datasource/gorm"
	"gorm.io/gorm"
	"os"
)

type Artist struct {
	ArtistId uint   `gorm:"primaryKey;column:ArtistId"`
	Name     string `gorm:"column:Name"`
}

func (Artist) TableName() string {
	return "artists"
}

func main() {

	_ = os.Setenv("LOG_LEVEL", "TRACE")
	cserver.Run("base-micro", "1.0.0", func(application cserver.Application) error {

		datasourceCtx := dgorm.NewContext("../resources/db/chinook.db?cache=shared&_pragma=foreign_keys(1)", "user", "pass", "localhost", "sqlite")
		connection := dgorm.NewConnection(datasourceCtx, sqlite.Open)
		transactionHandler := dgorm.NewTransactionHandler(connection)

		ctx := context.TODO()
		err := transactionHandler.HandleTransaction(ctx, func(ctx context.Context, tx *gorm.DB) error {

			insert := &Artist{Name: "D42"}
			tx.Create(insert)
			fmt.Printf("insert ID: %d, Name: %s\n", insert.ArtistId, insert.Name)

			read := &Artist{}
			tx.First(&read, "Name = ?", "D42")
			fmt.Printf("read ID: %d, Name: %s\n", read.ArtistId, read.Name)

			return nil
		})
		if err != nil {
			return err
		}

		return nil
	})
}
