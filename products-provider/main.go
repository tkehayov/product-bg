package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/tkehayov/product-bg.git/products-provider/products"
	provider "github.com/tkehayov/product-bg.git/proto/product-provider"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func init() {
	log.SetFormatter(&log.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})
	log.SetOutput(os.Stdout)
}

func init() {
	log.SetFormatter(&log.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})
	log.SetOutput(os.Stdout)
}

func main() {
	log.Info("PRODUCTS PROVIDER SERVICE STARTED")
	port := os.Getenv("PORT")

	router := mux.NewRouter()
	router.HandleFunc("/products-provider", ingest).Methods("POST")

	log.Error(http.ListenAndServe(":"+port, router))
}

func ingest(w http.ResponseWriter, r *http.Request) {

	products := products.ProductDto{}
	body, _ := ioutil.ReadAll(r.Body)
	unmarshall(body, &products)

	log.Info(products)
	runClient(products)
}

func unmarshall(d []byte, productDto *products.ProductDto) {
	err := json.Unmarshal(d, &productDto)
	if err != nil {
		log.Error("error unmarshalling: ", err.Error())
		return
	}
}

func runClient(products products.ProductDto) {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := provider.NewProductServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var chatPr []*provider.Product
	for _, pr := range products.Products {

		chatPr = append(chatPr, &provider.Product{
			CodeId:        pr.CodeId,
			Price:         pr.Price,
			DeliveryPrice: pr.DeliveryPrice,
			Url:           pr.Url,
		})
	}

	r, err := c.SendProducts(ctx, &provider.Message{
		MerchantId: products.MerchantId,
		Products:   chatPr,
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Client: %s", r.GetMerchantId())
}
