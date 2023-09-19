package infrastructure

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Database struct {
	DB *mongo.Client
}

func NewDatabase(env Env) Database {
	fmt.Print("database connection")
	uri := env.MongoURI
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environmental variable")
	}
	ctx, _ := context.WithTimeout(context.Background(),
		30*time.Second)

	// mongo.Connect return mongo.Client method
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(env.MongoURI))
	if err != nil {
		fmt.Printf("Url:%s", env.MongoURI)
		panic(err)
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		panic(err)
	}
	return Database{
		DB: client,
	}
}
