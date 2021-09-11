package main

import (
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"net"
	"os"
	"product-bg/products/internal/database"
	"product-bg/proto/products"
)

func init() {
	log.SetFormatter(&log.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})
	log.SetOutput(os.Stdout)
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
	database.UpdateProduct(in)

	return &products.Message{
		MerchantId: in.MerchantId,
	}, nil
}
