package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/glebarez/sqlite"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"gorm.io/gorm"

	cserver "github.com/guidomantilla/go-feather-lib/pkg/common/server"
	dgorm "github.com/guidomantilla/go-feather-lib/pkg/datasource/gorm"
	dmongo "github.com/guidomantilla/go-feather-lib/pkg/datasource/mongo"
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

		return case_gorm()
	})
}

func case_gorm() error {

	ctx := context.TODO()
	gormCtx := dgorm.NewContext("../resources/db/chinook.db?cache=shared&_pragma=foreign_keys(1)", "user", "pass", "localhost", "sqlite")
	connection := dgorm.NewConnection(gormCtx, sqlite.Open)
	defer connection.Close(ctx)
	transactionHandler := dgorm.NewTransactionHandler(connection)

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
}

func case_mongo() error {

	ctx := context.TODO()
	mongoCtx := dmongo.NewContext("mongodb://:username::password@:server", "root", "Raven123qweasd*+", "170.187.157.212:27017")
	opts := options.Client()
	connection := dmongo.NewConnection(mongoCtx, opts)
	defer connection.Close(ctx)
	transactionHandler := dmongo.NewTransactionHandler(connection)

	result, err := transactionHandler.HandleTransaction(ctx, func(client *mongo.Client, sesctx context.Context) (any, error) {

		collection := client.Database("sample").Collection("sample")
		res, err := collection.InsertOne(ctx, bson.D{{Key: "name", Value: "pi"}, {Key: "value", Value: 3.14159}})
		if err != nil {
			return nil, err
		}
		id := res.InsertedID
		log.Println(id)

		cur, err := collection.Find(ctx, bson.D{})
		if err != nil {
			return nil, err
		}
		defer cur.Close(ctx)
		for cur.Next(ctx) {
			var result bson.D
			if err := cur.Decode(&result); err != nil {
				log.Fatal(err)
			}

			// do something with result....
		}

		if err := cur.Err(); err != nil {
			log.Fatal(err)
		}

		var result struct {
			Value float64
		}

		filter := bson.D{{Key: "name", Value: "pi"}}
		err = collection.FindOne(ctx, filter).Decode(&result)
		if errors.Is(err, mongo.ErrNoDocuments) {
			log.Println("no documents found")
		} else if err != nil {
			log.Fatal(err)
		}
		log.Println(result)

		return result, nil
	})
	if err != nil {
		return err
	}

	log.Println(result)

	return nil
}
