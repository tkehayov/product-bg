package main

import (
	"context"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"net"
	"os"
	"product-bg/merchants/internal/database"
	"product-bg/proto/merchants"
	"time"
)

func main() {
	log.Info("MERCHANT CONSUMER SERVICE STARTED")
	env := os.Getenv("CONSUMER_PORT")
	lis, err1 := net.Listen("tcp", ":"+env)
	if err1 != nil {
		log.Fatalf("failed to listen: %v", err1)
	}
	s := grpc.NewServer()
	merchants.RegisterMerchantServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

type server struct {
	merchants.UnimplementedMerchantServiceServer
}

func (s *server) SendMerchant(ctx context.Context, merchant *merchants.Merchant) (*merchants.Logo, error) {
	logo := getLogo(merchant.Id)

	return &merchants.Logo{
		Logo: logo,
	}, nil
}

func getLogo(id string) string {
	merchantID, errMerchantID := primitive.ObjectIDFromHex(id)
	var logo database.MerchantLogo

	client, ctx := connect()
	defer client.Disconnect(ctx)

	db := client.Database("merchants")
	collection := db.Collection("merchants")

	if errMerchantID != nil {
		log.Error("error parsing id: ", errMerchantID.Error())
	}

	criteria := bson.D{{"_id", merchantID}}

	err := collection.FindOne(context.TODO(), criteria).Decode(&logo)
	if err != nil {
		log.Error("cannot get logo: ", err)
	}
	return logo.Logo
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
