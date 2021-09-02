package repo

import (
	"context"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"time"
)

type Product struct {
	ID         string     `bson:"_id" json:"id"`
	Name       string     `bson:"name" json:"name"`
	CodeId     string     `bson:"codeId" json:"codeId"`
	Gallery    Gallery    `bson:"gallery" json:"gallery"`
	Properties []Property `bson:"properties" json:"properties"`
	Merchants  []Merchant `bson:"merchants" json:"merchants"`
}

type Property struct {
	Name  string `bson:"name" json:"name"`
	Value string `bson:"value" json:"value"`
}
type Gallery struct {
	FeatureImage string   `json:"featureImage"`
	Images       []string `json:"images"`
}

type Merchant struct {
	ID           string  `bson:"_id" json:"id"`
	ProductTitle string  `bson:"productTitle" json:"productTitle"`
	Price        float64 `bson:"price" json:"price"`
	ShippingFee  float64 `bson:"shippingFee" json:"shippingFee"`
	Url          string  `bson:"url" json:"url"`
	Logo         string  `json:"logo"`
}

func GetOne(id string) Product {
	prodId, errProd := primitive.ObjectIDFromHex(id)
	if errProd != nil {
		log.Error("error parsing id: ", errProd.Error())
	}
	client, ctx := connect()
	defer client.Disconnect(ctx)

	db := client.Database("products")
	collection := db.Collection("products")

	product := bson.D{{"_id", prodId}}
	var result Product

	err := collection.FindOne(context.TODO(), product).Decode(&result)
	if err != nil {
		log.Error(err)
	}

	return result
}

func connect() (*mongo.Client, context.Context) {
	env := os.Getenv("MONGO_URL")
	client, err := mongo.NewClient(options.Client().ApplyURI(env))
	if err != nil {
		log.Error("Cannot connect to db", err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	errConn := client.Connect(ctx)
	if errConn != nil {
		log.Error("Error connection: ", errConn)
	}
	return client, ctx
}
