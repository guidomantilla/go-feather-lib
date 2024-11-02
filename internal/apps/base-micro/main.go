package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	gocqlastra "github.com/datastax/gocql-astra"
	"github.com/glebarez/sqlite"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"gorm.io/gorm"

	cserver "github.com/guidomantilla/go-feather-lib/pkg/common/server"
	dgocql "github.com/guidomantilla/go-feather-lib/pkg/datasource/gocql"
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

		return case_gocql()
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

func case_gocql() error {

	ctx := context.TODO()
	gocqlCtx := dgocql.NewContext("PublicIP", "Username", "PublicIP;PublicIP;PublicIP")
	dialer, _ := gocqlastra.NewDialerFromURL(gocqlastra.AstraAPIURL, "<astra-database-id>", "<astra-token", 10*time.Second)
	connection := dgocql.NewConnection(gocqlCtx, dgocql.ConnectionOptionsBuilder().WithDialer(dialer).Build())
	defer connection.Close(ctx)

	//cluster, err := gocqlastra.NewClusterFromURL(gocqlastra.AstraAPIURL, "<astra-database-id>", "<astra-token>", 10*time.Second)
	//cluster, err = gocqlastra.NewClusterFromBundle("/path/to/your/bundle.zip", "<username>", "<password>", 10*time.Second)
	//cluster := gocql.NewCluster("PublicIP", "PublicIP", "PublicIP") //replace PublicIP with the IP addresses used by your cluster.
	//cluster.Consistency = gocql.Quorum
	//cluster.ProtoVersion = 4
	//cluster.ConnectTimeout = time.Second * 10
	//cluster.Authenticator = gocql.PasswordAuthenticator{Username: "Username", Password: "Password", AllowedAuthenticators: []string{"com.instaclustr.cassandra.auth.InstaclustrPasswordAuthenticator"}} //replace the username and password fields with their real settings, you will need to allow the use of the Instaclustr Password Authenticator.

	session, err := connection.Connect(ctx)
	if err != nil {
		log.Println(err)
		return err
	}
	defer session.Close()

	// create keyspaces
	err = session.Query("CREATE KEYSPACE IF NOT EXISTS sleep_centre WITH REPLICATION = {'class' : 'NetworkTopologyStrategy', 'AWS_VPC_US_WEST_2' : 3};").Exec() //Replace AWS_VPC_US_WEST_2 with the name of the DataCentre you are connecting to.
	if err != nil {
		log.Println(err)
		return err
	}

	// create table
	err = session.Query("CREATE TABLE IF NOT EXISTS sleep_centre.sleep_study (name text, study_date date, sleep_time_hours float, PRIMARY KEY (name, study_date));").Exec()
	if err != nil {
		log.Println(err)
		return err
	}

	// insert some practice data
	err = session.Query("INSERT INTO sleep_centre.sleep_study (name, study_date, sleep_time_hours) VALUES ('James', '2018-01-07', 8.2);").Exec()
	err = session.Query("INSERT INTO sleep_centre.sleep_study (name, study_date, sleep_time_hours) VALUES ('James', '2018-01-08', 6.4);").Exec()
	err = session.Query("INSERT INTO sleep_centre.sleep_study (name, study_date, sleep_time_hours) VALUES ('James', '2018-01-09', 7.5);").Exec()
	err = session.Query("INSERT INTO sleep_centre.sleep_study (name, study_date, sleep_time_hours) VALUES ('Bob', '2018-01-07', 6.6);").Exec()
	err = session.Query("INSERT INTO sleep_centre.sleep_study (name, study_date, sleep_time_hours) VALUES ('Bob', '2018-01-08', 6.3);").Exec()
	err = session.Query("INSERT INTO sleep_centre.sleep_study (name, study_date, sleep_time_hours) VALUES ('Bob', '2018-01-09', 6.7);").Exec()
	err = session.Query("INSERT INTO sleep_centre.sleep_study (name, study_date, sleep_time_hours) VALUES ('Emily', '2018-01-07', 7.2);").Exec()
	err = session.Query("INSERT INTO sleep_centre.sleep_study (name, study_date, sleep_time_hours) VALUES ('Emily', '2018-01-09', 7.5);").Exec()
	if err != nil {
		log.Println(err)
		return err
	}

	// Return average sleep time for James
	var sleep_time_hours float32

	sleep_time_output := session.Query("SELECT avg(sleep_time_hours) FROM sleep_centre.sleep_study WHERE name = 'James';").Iter()
	sleep_time_output.Scan(&sleep_time_hours)
	fmt.Println("Average sleep time for James was: ", sleep_time_hours, "h")

	// return average sleep time for group
	sleep_time_output = session.Query("SELECT avg(sleep_time_hours) FROM sleep_centre.sleep_study;").Iter()
	sleep_time_output.Scan(&sleep_time_hours)
	fmt.Println("Average sleep time for the group was: ", sleep_time_hours, "h")

	return nil
}
