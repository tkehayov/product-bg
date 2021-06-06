package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("port habe: ")
	io.WriteString(w, "Hello World71!")

}
func hi(w http.ResponseWriter, r *http.Request) {
	client, ctx := connect()
	defer client.Disconnect(ctx)

	db := client.Database("products")
	prCol := db.Collection("productInfo")

	_, err := prCol.InsertOne(ctx, bson.D{
		{Key: "title", Value: "The Polyglot Developer Podcast"},
		{Key: "author", Value: "Nic Raboy"},
	})

	if err != nil {
		fmt.Println("Error insert many: ", err)
	}

}

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

func main() {

	port := os.Getenv("PORT")

	http.HandleFunc("/", hello)
	http.HandleFunc("/hi", hi)
	http.ListenAndServe(":"+port, nil)
}
