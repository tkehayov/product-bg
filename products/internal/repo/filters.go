package repo

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"product-bg/products/internal/database"
	"product-bg/products/internal/dto"
	"product-bg/products/internal/entities"
)

type ProductCategoryRepositoryInterface interface {
	//GetFilters(category string) entities.Category
}
type productCategoryRepostiory struct{}

func NewProductCategoryRepostiory() ProductCategoryRepositoryInterface {
	return &productCategoryRepostiory{}
}

func GetFilters(category string) dto.Category {
	client, ctx := database.Connect()
	defer client.Disconnect(ctx)

	db := client.Database("products")
	collection := db.Collection("products")

	var result dto.Category

	unwindStage := bson.D{{"$unwind", bson.D{{"path", "$properties"}}}}
	groupStage := bson.D{{"$group", bson.D{{"_id", "$properties.name"}, {"values", bson.D{{"$addToSet", "$properties.value"}}}}}}

	filters := getCategoryFilters(category)

	matchStage := bson.D{{"$match", bson.D{{"$or", filters}}}}
	projectStage := bson.D{{"$project", bson.D{{"values", 1}, {"name", "$_id"}, {"_id", 0}}}}

	cursorAggregator, errAggregator := collection.Aggregate(context.TODO(), mongo.Pipeline{unwindStage, groupStage, matchStage, projectStage})

	if errAggregator != nil {
		log.Error("Error aggregator: ", errAggregator.Error())
	}

	if errAggregator = cursorAggregator.All(ctx, &result.Filter); errAggregator != nil {
		log.Error("Error cursor aggregator: ", errAggregator.Error())
	}

	return result
}

func getCategoryFilters(category string) []interface{} {
	var categoryEntity entities.Category
	var result []interface{}

	client, ctx := database.Connect()
	defer client.Disconnect(ctx)

	db := client.Database("products")
	collection := db.Collection("categories")

	err := collection.FindOne(context.TODO(), bson.D{{"_id", category}}).Decode(&categoryEntity)
	if err != nil {
		log.Error("Could not fetch category: ", err.Error())
	}

	for _, filter := range categoryEntity.Filter {
		categoryFormated := fmt.Sprintf("^%s$", filter.Value)
		filter1 := bson.D{{"_id", bson.D{{"$regex", categoryFormated}, {"$options", "i"}}}}
		result = append(result, filter1)
	}

	return result
}
