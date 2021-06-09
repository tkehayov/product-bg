package main

import (
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/tkehayov/product-bg.git/repo"
	"net/http"
	"os"
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
	router.HandleFunc("/product", GetOne)

	log.Fatal(http.ListenAndServe(":"+port, router))
}

func GetOne(w http.ResponseWriter, r *http.Request) {
	repo.GetOne()
}