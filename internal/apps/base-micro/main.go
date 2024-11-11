package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/glebarez/sqlite"
	redis "github.com/redis/go-redis/v9"
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

type BikeInfo struct {
	Model string `redis:"model"`
	Brand string `redis:"brand"`
	Type  string `redis:"type"`
	Price int    `redis:"price"`
}

func (Artist) TableName() string {
	return "artists"
}

func main() {

	_ = os.Setenv("LOG_LEVEL", "TRACE")
	cserver.Run("base-micro", "1.0.0", func(ctx context.Context, application cserver.Application) error {

		return case_mongo()
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
	loggerOptions := options.Logger().SetSink(dmongo.Logger()).
		SetComponentLevel(options.LogComponentCommand, options.LogLevelInfo).
		SetComponentLevel(options.LogComponentConnection, options.LogLevelDebug)
	opts := options.Client().SetLoggerOptions(loggerOptions)
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
	gocqlCtx := dgocql.NewContext("root", "Raven123qweasd*", "170.187.157.212")
	connection := dgocql.NewConnection(gocqlCtx)
	defer connection.Close(ctx)

	session, err := connection.Connect(ctx)
	if err != nil {
		log.Println(err)
		return err
	}
	defer session.Close()

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

func case_goredis() error {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "Raven123qweasd*+-789", // No password set
		DB:       0,                      // Use default DB
		//UnstableResp3: true,                   // Habilita RESP3
		Protocol: 2,
	})
	/*
		client := redis.NewClusterClient(&redis.ClusterOptions{
			Addrs:    []string{"localhost:6379"},
			Password: "Raven123qweasd*+-789",

			// To route commands by latency or randomly, enable one of the following.
			//RouteByLatency: true,
			//RouteRandomly: true,
		})
	*/

	ctx := context.Background()

	err := client.Set(ctx, "foo", "bar", 0).Err()
	if err != nil {
		log.Println(err)
		return err
	}

	val, err := client.Get(ctx, "foo").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("foo", val)

	hashFields := []string{
		"model", "Deimos",
		"brand", "Ergonom",
		"type", "Enduro bikes",
		"price", "4972",
	}

	res1, err := client.HSet(ctx, "bike:1", hashFields).Result()

	if err != nil {
		panic(err)
	}

	fmt.Println(res1) // >>> 4

	res2, err := client.HGet(ctx, "bike:1", "model").Result()

	if err != nil {
		panic(err)
	}

	fmt.Println(res2) // >>> Deimos

	res3, err := client.HGet(ctx, "bike:1", "price").Result()

	if err != nil {
		panic(err)
	}

	fmt.Println(res3) // >>> 4972

	res4, err := client.HGetAll(ctx, "bike:1").Result()

	if err != nil {
		panic(err)
	}

	fmt.Println(res4)

	var res4a BikeInfo
	err = client.HGetAll(ctx, "bike:1").Scan(&res4a)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Model: %v, Brand: %v, Type: %v, Price: $%v\n",
		res4a.Model, res4a.Brand, res4a.Type, res4a.Price)

	//JSON
	{
		_, err = client.FTCreate(
			ctx,
			"idx:users-json",
			// Options:
			&redis.FTCreateOptions{
				OnJSON: true,
				Prefix: []interface{}{"user-json:"},
			},
			// Index schema fields:
			&redis.FieldSchema{
				FieldName: "$.name",
				As:        "name",
				FieldType: redis.SearchFieldTypeText,
			},
			&redis.FieldSchema{
				FieldName: "$.city",
				As:        "city",
				FieldType: redis.SearchFieldTypeTag,
			},
			&redis.FieldSchema{
				FieldName: "$.age",
				As:        "age",
				FieldType: redis.SearchFieldTypeNumeric,
			},
		).Result()

		if err != nil {
			log.Println(err)
		}

		userJson1 := map[string]interface{}{
			"name":  "Paul John",
			"email": "paul.john@example.com",
			"age":   42,
			"city":  "London",
		}

		userJson2 := map[string]interface{}{
			"name":  "Eden Zamir",
			"email": "eden.zamir@example.com",
			"age":   29,
			"city":  "Tel Aviv",
		}

		userJson3 := map[string]interface{}{
			"name":  "Paul Zamir",
			"email": "paul.zamir@example.com",
			"age":   35,
			"city":  "Tel Aviv",
		}

		_, err = client.JSONSet(ctx, "user-json:1", "$", userJson1).Result()

		if err != nil {
			panic(err)
		}

		_, err = client.JSONSet(ctx, "user-json:2", "$", userJson2).Result()

		if err != nil {
			panic(err)
		}

		_, err = client.JSONSet(ctx, "user-json:3", "$", userJson3).Result()
		if err != nil {
			panic(err)
		}

		searchResult, err := client.FTSearch(
			ctx,
			"idx:users-json",
			"Paul @age:[30 40]",
		).Result()

		if err != nil {
			panic(err)
		}

		fmt.Println(searchResult)
	}

	//HASH
	{

		userHash1 := []string{
			"name", "Paul John",
			"email", "paul.john@example.com",
			"age", "42",
			"city", "London",
		}

		userHash2 := []string{
			"name", "Eden Zamir",
			"email", "eden.zamir@example.com",
			"age", "29",
			"city", "Tel Aviv",
		}

		userHash3 := []string{
			"name", "Paul Zamir",
			"email", "paul.zamir@example.com",
			"age", "35",
			"city", "Tel Aviv",
		}

		_, err = client.FTCreate(
			ctx,
			"idx:users-hash",
			// Options:
			&redis.FTCreateOptions{
				OnHash: true,
				Prefix: []interface{}{"user-hash:"},
			},
			// Index schema fields:
			&redis.FieldSchema{
				FieldName: "$.name",
				As:        "name",
				FieldType: redis.SearchFieldTypeText,
			},
			&redis.FieldSchema{
				FieldName: "$.city",
				As:        "city",
				FieldType: redis.SearchFieldTypeTag,
			},
			&redis.FieldSchema{
				FieldName: "$.age",
				As:        "age",
				FieldType: redis.SearchFieldTypeNumeric,
			},
		).Result()

		if err != nil {
			log.Println(err)
		}

		_, err = client.HSet(ctx, "user-hash:1", userHash1).Result()

		if err != nil {
			panic(err)
		}

		_, err = client.HSet(ctx, "user-hash:2", userHash2).Result()

		if err != nil {
			panic(err)
		}

		_, err = client.HSet(ctx, "user-hash:3", userHash3).Result()

		if err != nil {
			panic(err)
		}

		searchResult, err := client.FTSearch(
			ctx,
			"idx:users-hash",
			"Paul @age:[30 40]",
		).Result()

		if err != nil {
			panic(err)
		}

		fmt.Println(searchResult)
	}

	/*
		err := client.Watch(func(tx *redis.Tx) error {
		    n, err := tx.Get(key).Int64()
		    if err != nil && err != redis.Nil {
		        return err
		    }

		    _, err = tx.TxPipelined(func(pipe *redis.Pipeline) error {
		        pipe.Set(key, strconv.FormatInt(n+1, 10), 0)
		        return nil
		    })
		    return err
		}, key)
	*/

	/*
		func (r *Repository) BuyShares(ctx context.Context, userId, companyId string, numShares int, wg *sync.WaitGroup) error {

		 defer wg.Done()

		 companySharesKey := BuildCompanySharesKey(companyId)

		 err := r.client.Watch(ctx, func(tx *goredislib.Tx) error {
		  // --- (1) ----
		  // Get current number of shares
		  currentShares, err := tx.Get(ctx, companySharesKey).Int()
		  if err != nil {
		   fmt.Print(fmt.Errorf("error getting value %v", err.Error()))
		   return err
		  }

		  // --- (2) ----
		  // Validate if the shares remaining are enough to be bought
		  if currentShares < numShares {
		   fmt.Print("error: company does not have enough shares \n")
		   return errors.New("error: company does not have enough shares")
		  }
		  currentShares -= numShares

		  // --- (3) ----
		  // Update the current shares of the company and log who has bought shares
		  _, err = tx.TxPipelined(ctx, func(pipe goredislib.Pipeliner) error {
		   // pipe handles the error case
		   pipe.Pipeline().Set(ctx, companySharesKey, currentShares, 0)
		   return nil
		  })
		  if err != nil {
		   fmt.Println(fmt.Errorf("error in pipeline %v", err.Error()))
		   return err
		  }
		  return nil
		 }, companySharesKey)
		 return err
		}
	*/
	return nil
}
