package repo

import (
	"context"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"product-bg/products/internal/database"
	"product-bg/products/internal/entities"
)

type CategoryRepositoryInterface interface {
	GetOne(id string) entities.Category
}

type category struct{}

func NewCategoryRepository() CategoryRepositoryInterface {
	return &category{}
}

func (cat *category) GetOne(category string) entities.Category {
	var categoryEntity entities.Category

	client, ctx := database.Connect()
	defer client.Disconnect(ctx)

	db := client.Database("products")
	collection := db.Collection("categories")

	err := collection.FindOne(context.TODO(), bson.D{{"_id", category}}).Decode(&categoryEntity)
	if err != nil {
		log.Error("Could not fetch category: ", err.Error())
	}

	return categoryEntity

}
