package mycontext

import (
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Context struct {
	Db *mongo.Database
}

func connectWithDatabase() (*mongo.Database, error) {
	uri := os.Getenv("MONGODB_URI")

	client, err := mongo.NewClient(options.Client().ApplyURI(uri))

	if err != nil {
		return nil, err
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	err = client.Connect(ctx)

	if err != nil {
		return nil, err
	}

	// Ping the primary
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		return nil, err
	}

	database := client.Database("salom-imagepool-db", nil)

	return database, nil
}

func NewContext() (*Context, error) {
	db, err := connectWithDatabase()

	if err != nil {
		return nil, err
	}

	return &Context{Db: db}, nil
}
