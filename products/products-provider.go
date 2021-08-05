package main

import (
	log "github.com/sirupsen/logrus"
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

func main() {
	log.Info("PRODUCTS PROVIDER SERVICE STARTED")
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
	Persist(in)

	return &provider.Message{
		MerchantId: in.MerchantId,
	}, nil
}

func Persist(m *provider.Message) {
	client, ctx := connect()
	defer client.Disconnect(ctx)

	db := client.Database("products-merchants")

	collection := db.Collection("products")

	_, err := collection.InsertOne(ctx, m)
	if err != nil {
		log.Error("Couldn't insert products: ", err.Error())
	}
}

func connect() (*mongo.Client, context.Context) {
	env := os.Getenv("MONGO_URL")
	log.Error("herrreee: ", env)
	client, err := mongo.NewClient(options.Client().ApplyURI(env))
	if err != nil {
		log.Fatal("error Connection", err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)

	return client, ctx
}
