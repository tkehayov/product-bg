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

// GetFilteredProducts Filtering products by criteria.
// Example:
//
// {$and: [
//        {$or: [
//                {
//                    "properties.name": "Памет",
//                    "properties.value": "8GB"
//                }
//            ]
//        },
//        {$or: [
//                {
//                    "properties.name": "Памет",
//                    "properties.value": "Intel Core i5"
//                }
//            ]
//        },
//        {
//            "category": "laptops",
//            "brand": "Acer",
//        },
//    ]
//	}
func (p productFilter) GetFilteredProducts(category string, filters map[string][]string) []entities.ProductFilter {
	client, ctx := database.Connect()
	defer client.Disconnect(ctx)

	db := client.Database("products")
	collection := db.Collection("products")

	matchStage := filterDocument(category, filters)

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
		minPrice := getMinMerchantPrice(product.Merchants)
		product.MinPrice = minPrice

		if err != nil {
			log.Error("Error decoding products: ", err)
		}

		products = append(products, product)
	}

	return products
}

func filterDocument(category string, filters map[string][]string) bson.D {
	params := bson.A{}
	var brand string
	var categoryStage bson.A
	var orParams bson.A

	categoryStage = bson.A{bson.D{{"category", category}}}
	//TODO delete after moving product.brand into properties.name. Example:
	//  {product.brand:"Acer"} ,should look like {product.properties.name:"brand",product.properties.value:"Acer"}
	brandParam := filters["brand"]
	if brandParam != nil {
		brand = brandParam[0]
		delete(filters, "brand")

		categoryStage = bson.A{bson.D{{"category", category}, {"brand", brand}}}
	}

	for key, values := range filters {
		for _, value := range values {
			properties := bson.D{{"properties.name", key}, {"properties.value", value}}
			params = append(params, properties)
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

func getMinMerchantPrice(merchants []entities.Merchant) float64 {
	var prices []float64
	for _, merchant := range merchants {
		prices = append(prices, merchant.Price)
	}
	return min(prices)
}

func min(array []float64) float64 {
	var min = array[0]
	for _, value := range array {
		if min > value {
			min = value
		}
	}

	return min
}
