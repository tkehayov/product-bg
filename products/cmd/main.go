package main

import (
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"product-bg/products/internal/rest"
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
	router.HandleFunc("/products/{id}", rest.Product{}.GetOne)
	router.HandleFunc("/filters/categories/{category}", rest.FilterCategory{}.GetAll)
	//TODO implement
	router.HandleFunc("/filters/products/{category}", rest.FilterProduct{}.Get)
	log.Error(http.ListenAndServe(":"+port, router))
}
