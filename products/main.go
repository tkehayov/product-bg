package main

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net/http"
	"os"
	"product-bg/products/repo"
	"product-bg/proto/merchants"
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
	log.Info("PRODUCT SERVICE STARTED")
	port := os.Getenv("PORT")
	router := mux.NewRouter()
	router.HandleFunc("/products/{id}", GetOne)

	log.Error(http.ListenAndServe(":"+port, router))
}

func GetOne(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	product := repo.GetOne(id)

	for i, merchant := range product.Merchants {
		logo := getLogo(merchant.ID)

		merchant.Logo = logo.Logo
		product.Merchants[i] = merchant
	}
	response, err := json.Marshal(product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func getLogo(Id string) *merchants.Logo {
	env := os.Getenv("LOGO_PROVIDER")
	conn, err := grpc.Dial(env, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := merchants.NewMerchantServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	logo, err := client.SendMerchant(ctx, &merchants.Merchant{
		Id: Id,
	})
	if err != nil {
		log.Fatalf("could not send logo: %v", err)
	}
	return logo
}
