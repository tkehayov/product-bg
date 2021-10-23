package repo

import (
	"context"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"product-bg/products/internal/database"
	"product-bg/products/internal/entities"
)

type ProductCategoryFilterRepositoryInterface interface {
	GetFilters(category string) entities.CategoryProductFilter
}

type productCategoryFilter struct{}

func NewProductCategoryFilterRepository() ProductCategoryFilterRepositoryInterface {
	return &productCategoryFilter{}
}

func (productCategory *productCategoryFilter) GetFilters(category string) entities.CategoryProductFilter {
	client, ctx := database.Connect()
	defer client.Disconnect(ctx)

	db := client.Database("products")
	collection := db.Collection("products")

	var result entities.CategoryProductFilter

	unwindStage := bson.D{{"$unwind", bson.D{{"path", "$properties"}}}}

	filters := getCategoryFilters(category)
	matchStage := bson.D{{"$match", bson.D{{"properties.name", bson.D{{"$in", filters}}}, {"category", "laptops"}}}}
	groupStage := bson.D{{"$group", bson.D{{"_id", "$properties.name"}, {"values", bson.D{{"$addToSet", "$properties.value"}}}}}}
	projectStage := bson.D{{"$project", bson.D{{"values", 1}, {"name", "$_id"}, {"_id", 0}}}}

	cursorAggregator, errAggregator := collection.Aggregate(context.TODO(), mongo.Pipeline{unwindStage, matchStage, groupStage, projectStage})

	if errAggregator != nil {
		log.Error("Error aggregator: ", errAggregator.Error())
	}

	if errAggregator = cursorAggregator.All(ctx, &result.ProductFilter); errAggregator != nil {
		log.Error("Error cursor aggregator: ", errAggregator.Error())
	}

	return result
}

func getCategoryFilters(category string) []string {
	var categoryEntity entities.Category
	var result []string

	client, ctx := database.Connect()
	defer client.Disconnect(ctx)

	db := client.Database("products")
	collection := db.Collection("categories")

	err := collection.FindOne(context.TODO(), bson.D{{"_id", category}}).Decode(&categoryEntity)
	if err != nil {
		log.Error("Could not fetch category: ", err.Error())
	}

	for _, filter := range categoryEntity.Filter {
		result = append(result, filter.Value)
	}

	return result
}
