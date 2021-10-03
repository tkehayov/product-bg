package repo

import (
	"context"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"product-bg/products/internal/database"
)

type Category struct {
	Filter []Filter `bson:"filters"  json:"filters"`
}

type Filter struct {
	Name  string `bson:"name" json:"name"`
	Value string `bson:"value" json:"value"`
}

func GetFilters(category string) Category {
	client, ctx := database.Connect()
	defer client.Disconnect(ctx)

	db := client.Database("products")
	collection := db.Collection("categories")

	var result Category

	criteria := bson.D{{"_id", category}}
	err := collection.FindOne(context.TODO(), criteria).Decode(&result)
	if err != nil {
		log.Error("Error on finding filter: ", err.Error())
	}

	return result
}
