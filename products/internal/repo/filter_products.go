package repo

import (
	"context"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"product-bg/products/internal/database"
	"product-bg/products/internal/entities"
)

type ProductFilterRepositoryInterface interface {
	GetFilteredProducts(category string, filter map[string][]string) []entities.ProductFilter
}

type productFilter struct{}

func NewProductFilterRepository() ProductFilterRepositoryInterface {
	return &productFilter{}
}

func (p productFilter) GetFilteredProducts(category string, filters map[string][]string) []entities.ProductFilter {
	client, ctx := database.Connect()
	defer client.Disconnect(ctx)

	db := client.Database("products")
	collection := db.Collection("products")

	params := bson.A{}

	matchStage := filterDocument(category, filters, params)

	cur, err := collection.Find(context.TODO(), matchStage)

	var products []entities.ProductFilter
	if err != nil {
		log.Error("Error finding products: ", err)
		return products
	}

	for cur.Next(context.TODO()) {
		// create a value into which the single document can be decoded
		var product entities.ProductFilter
		err := cur.Decode(&product)
		if err != nil {
			log.Error("Error decoding products: ", err)
		}

		products = append(products, product)
	}

	return products
}

func filterDocument(category string, filters map[string][]string, params bson.A) bson.D {
	categoryStage := bson.A{bson.D{{"category", category}}}
	var orParams bson.A

	for key, values := range filters {
		for _, value := range values {
			processors := bson.D{{"properties.name", key}, {"properties.value", value}}
			params = append(params, processors)
		}
		orParams = append(orParams, bson.D{{"$or", params}})
		params = bson.A{}
	}

	matchStage := bson.D{{"$and", orParams}, {"$and", categoryStage}}
	if orParams == nil {
		matchStage = bson.D{{"$and", categoryStage}}
	}

	return matchStage
}
