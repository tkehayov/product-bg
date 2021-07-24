package main

import (
	"encoding/json"
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
	router.HandleFunc("/products/{id}", GetOne)

	log.Error(http.ListenAndServe(":"+port, router))
}

func GetOne(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	product := repo.GetOne(id)

	response, err := json.Marshal(product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
