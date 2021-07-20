package main

import (
	log "github.com/sirupsen/logrus"
	provider "github.com/tkehayov/product-bg.git/proto/product-provider"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"net"
)

func main() {
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
	log.Print("Receivedd: ", in)
	return &provider.Message{
		MerchantId: in.MerchantId,
	}, nil
}
