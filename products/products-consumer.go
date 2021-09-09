package main

import (
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"net"
	"os"
	"product-bg/proto/products"
	"time"
)

func init() {
	log.SetFormatter(&log.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})
	log.SetOutput(os.Stdout)
}

type Merchant struct {
	ID           string  `bson:"_id"`
	ProductTitle string  `bson:"productTitle,omitempty" json:"productTitle"`
	Price        float64 `bson:"price,omitempty" json:"price"`
	ShippingFee  float64 `bson:"shippingFee,omitempty" json:"shippingFee"`
	Url          string  `bson:"url,omitempty" json:"url"`
}

func main() {
	log.Info("PRODUCTS CONSUMER SERVICE STARTED")
	env := os.Getenv("CONSUMER_PORT")
	lis, err1 := net.Listen("tcp", ":"+env)
	if err1 != nil {
		log.Fatalf("failed to listen: %v", err1)
	}
	s := grpc.NewServer()
	products.RegisterProductServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

type server struct {
	products.UnimplementedProductServiceServer
}

func (s *server) SendProducts(ctx context.Context, in *products.Message) (*products.Message, error) {
	UpdateProduct(in)

	return &products.Message{
		MerchantId: in.MerchantId,
	}, nil
}

//TODO move into repository
func UpdateProduct(m *products.Message) {
	client, ctx := connect()
	defer client.Disconnect(ctx)

	db := client.Database("products")
	collection := db.Collection("products")

	for _, product := range m.Products {
		merchant := Merchant{
			ID:           m.MerchantId,
			ProductTitle: product.ProductTitle,
			Price:        product.Price,
			ShippingFee:  product.ShippingFee,
			Url:          product.Url,
		}

		res, err := updateMerchant(collection, ctx, product, merchant)

		if err != nil {
			resultNewMerchant, errNewMerchant := addMerchant(collection, ctx, product, merchant)
			if errNewMerchant != nil {
				log.Error("Cannot add new merchant: ", errNewMerchant.Error())
			}

			log.Info("Inserted new merchant to a product", resultNewMerchant.ModifiedCount)
			continue
		}

		log.Info("Updated merchant product", res.ModifiedCount)
	}
}

func updateMerchant(collection *mongo.Collection, ctx context.Context, product *products.Product, merchant Merchant) (*mongo.UpdateResult, error) {
	opts := options.Update().SetUpsert(true)
	res, err := collection.UpdateOne(ctx, bson.M{"codeId": product.CodeId, "merchants._id": merchant.ID}, bson.D{
		{"$set",
			bson.D{{"merchants.$", merchant}},
		},
	}, opts)

	return res, err
}

func addMerchant(collection *mongo.Collection, ctx context.Context, product *products.Product, merchant Merchant) (*mongo.UpdateResult, error) {
	resultNewMerchant, err := collection.UpdateOne(ctx, bson.M{"codeId": product.CodeId}, bson.D{
		{"$push", bson.D{{"merchants", merchant}}},
	})

	return resultNewMerchant, err
}

func connect() (*mongo.Client, context.Context) {
	env := os.Getenv("MONGO_URL")
	client, err := mongo.NewClient(options.Client().ApplyURI(env))
	if err != nil {
		log.Fatal("error Connection", err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)

	return client, ctx
}
