package repo

import (
	"context"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
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
//{$and: [
//		  {"_id":{$gt:ObjectId("5ef8b9eb5512aa0006c60f12")}},
//       {$or: [
//               {
//                   "properties.name": "Памет",
//                   "properties.value": "8GB"
//               }
//           ]
//       },
//       {$or: [
//               {
//                   "properties.name": "Памет",
//                   "properties.value": "Intel Core i5"
//               }
//           ]
//       },
//       {
//           "category": "laptops",
//           "brand": "Acer",
//       },
//   ]
//	}.limit(10)
func (p productFilter) GetFilteredProducts(category string, filters map[string][]string) []entities.ProductFilter {
	client, ctx := database.Connect()
	defer client.Disconnect(ctx)

	db := client.Database("products")
	collection := db.Collection("products")

	matchStage, sort := filterDocument(category, filters)

	options := options.Find()
	options.SetLimit(10)
	if sort != nil {
		options.SetSort(sort)
	}
	cur, err := collection.Find(context.TODO(), matchStage, options)

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

func filterDocument(category string, filters map[string][]string) (result bson.D, sort bson.D) {
	params := bson.A{}
	var andStage bson.A
	var orStage bson.A

	categoryDocument := bson.D{{"category", category}}
	andStage = bson.A{}

	andStage, sort = paginate(filters, andStage)
	categoryDocument = addBrand(filters, categoryDocument)

	andStage = append(andStage, categoryDocument)

	for key, values := range filters {
		for _, value := range values {
			properties := bson.D{{"properties.name", key}, {"properties.value", value}}
			params = append(params, properties)
		}
		orStage = append(orStage, bson.D{{"$or", params}})
		params = bson.A{}
	}

	matchStage := bson.D{{"$and", orStage}, {"$and", andStage}}
	if orStage == nil {
		matchStage = bson.D{{"$and", andStage}}
	}

	return matchStage, sort
}

func paginate(filters map[string][]string, andStage bson.A) (result bson.A, sort bson.D) {
	document, sort := paginationDocument(filters)

	if document != nil {
		idElement := bson.D{{"_id", document}}
		andStage = append(andStage, idElement)
	}

	return andStage, sort
}

func paginationDocument(filters map[string][]string) (result bson.D, sort bson.D) {
	var beforeAfterParam string
	var document bson.D
	var sortAfterBefore bson.D
	afterParam := filters["after"]
	beforeParam := filters["before"]

	if afterParam != nil {
		beforeAfterParam = afterParam[0]
		delete(filters, "after")
		id, _ := primitive.ObjectIDFromHex(beforeAfterParam)
		document = bson.D{{"$gt", id}}
		sortAfterBefore = bson.D{{"_id", 1}}
	}

	if beforeParam != nil {
		beforeAfterParam = beforeParam[0]
		delete(filters, "before")
		id, _ := primitive.ObjectIDFromHex(beforeAfterParam)
		document = bson.D{{"$lt", id}}
		sortAfterBefore = bson.D{{"_id", -1}}
	}

	return document, sortAfterBefore
}

func addBrand(filters map[string][]string, categoryDocument bson.D) bson.D {
	//TODO delete after moving product.brand into properties.name. Example:
	//  {product.brand:"Acer"}, should look like {product.properties.name:"brand",product.properties.value:"Acer"}
	brandParam := filters["brand"]
	if brandParam != nil {
		brand := brandParam[0]
		delete(filters, "brand")

		brandElement := bson.E{"brand", brand}
		categoryDocument = append(categoryDocument, brandElement)
	}
	return categoryDocument
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
