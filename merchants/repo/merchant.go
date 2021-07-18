package repo

import (
	"context"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"time"
)

type Merchant struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Website  string `json:"website"`
}

func Register(m Merchant) {
	client, ctx := connect()
	defer client.Disconnect(ctx)

	db := client.Database("merchants")

	collection := db.Collection("merchants")

	_, err := collection.InsertOne(ctx, m)
	if err != nil {
		log.Error("Couldn't insert merchant: ", err.Error())
	}
}

func connect() (*mongo.Client, context.Context) {
	env := os.Getenv("MONGO_URL")

	client, err := mongo.NewClient(options.Client().ApplyURI(env))
	if err != nil {
		log.Fatal("error Connection", err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)

	return client, ctx
}
