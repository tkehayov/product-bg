package database

import (
	"context"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"time"
)

func Connect() (*mongo.Client, context.Context) {
	env := os.Getenv("MONGO_URL")
	client, err := mongo.NewClient(options.Client().ApplyURI(env))
	if err != nil {
		log.Error("Cannot connect to db", err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	errConn := client.Connect(ctx)
	if errConn != nil {
		log.Error("Error connection: ", errConn)
	}
	return client, ctx
}
