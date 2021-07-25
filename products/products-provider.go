package main

import (
	log "github.com/sirupsen/logrus"
	"os"
)

func init() {
	log.SetFormatter(&log.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})
	log.SetOutput(os.Stdout)
}

//func main() {
//	log.Info("PRODUCTS PROVIDER SERVICE STARTED")
//	lis, err1 := net.Listen("tcp", ":50051")
//	if err1 != nil {
//		log.Fatalf("failed to listen: %v", err1)
//	}
//	s := grpc.NewServer()
//	provider.RegisterProductServiceServer(s, &server{})
//	if err := s.Serve(lis); err != nil {
//		log.Fatalf("failed to serve: %v", err)
//	}
//}

//type server struct {
//	provider.UnimplementedProductServiceServer
//}
//
//func (s *server) SendProducts(ctx context.Context, in *provider.Message) (*provider.Message, error) {
//	log.Print("Received: ", in)
//	return &provider.Message{
//		MerchantId: in.MerchantId,
//	}, nil
//}
