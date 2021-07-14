package dbUtil

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type DBWrapper struct {
	Ctx    context.Context
	Client *mongo.Client
	DB     *mongo.Database
}

func Init() (*DBWrapper, error) {
	// connection mongodb
	uri := "mongodb://localhost:27017"
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	db := client.Database("example")

	return &DBWrapper{Ctx: ctx, Client: client, DB: db}, nil
}

func (dwb *DBWrapper) Close() {
	err := dwb.Client.Disconnect(dwb.Ctx)
	if err != nil {
		log.Panic(err)
	}
}
