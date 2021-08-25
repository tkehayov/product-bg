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
	"product-bg/proto/provider"
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
	lis, err1 := net.Listen("tcp", ":50051")
	if err1 != nil {
		log.Fatalf("failed to listen: %v", err1)
	}
	s := grpc.NewServer()
	provider.RegisterProductServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

type server struct {
	provider.UnimplementedProductServiceServer
}

func (s *server) SendProducts(ctx context.Context, in *provider.Message) (*provider.Message, error) {
	UpdateProduct(in)

	return &provider.Message{
		MerchantId: in.MerchantId,
	}, nil
}

func UpdateProduct(m *provider.Message) {
	client, ctx := connect()
	defer client.Disconnect(ctx)

	db := client.Database("products")

	collection := db.Collection("products")

	//TODO loop products
	res, err := collection.UpdateOne(ctx, bson.M{"codeId": m.Products[0].CodeId}, bson.D{
		{"$set", bson.D{{"merchant",
			Merchant{ID: m.MerchantId, ProductTitle: m.Products[0].ProductTitle, Price: m.Products[0].Price, ShippingFee: m.Products[0].ShippingFee, Url: m.Products[0].Url}}}},
	})
	//TODO update instead of insert
	//_, err := collection.InsertOne(ctx, m)
	if err != nil {
		log.Error("Couldn't update products: ", err.Error())
	}
	log.Info("Updated merchant product", res.ModifiedCount)
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
