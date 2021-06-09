package repo

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"time"
)

func connect() (*mongo.Client, context.Context) {
	env := os.Getenv("MONGO_URL")
	fmt.Println("emc", env)

	client, err := mongo.NewClient(options.Client().ApplyURI(env))
	if err != nil {
		log.Fatal("error Connection", err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)

	return client, ctx
}

type Product struct {
	ID              string     `bson:"_id"`
	codeId        string `json:"children"`
}

func GetOne() {
	client, ctx := connect()
	defer client.Disconnect(ctx)

	db := client.Database("products")
	prCol := db.Collection("products")

	filter := bson.D{{"codeId", "FQC-09478"}}
	var result Product

	err := prCol.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	log.Error("result",result)
}
