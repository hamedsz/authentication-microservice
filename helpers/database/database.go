package database

import (
	"auth_micro/helpers/env"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
)

var Database , Ctx = getDatabase()

func getDatabase() (*mongo.Database , context.Context){

	println("new connection")

	client, err := mongo.
		NewClient(
			options.Client().
				ApplyURI(
					env.GetEnv("MONGO_URL")))

	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithCancel(context.Background())

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	return client.Database(env.GetEnv("MONGO_DB")) , ctx
}