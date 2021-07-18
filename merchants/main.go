package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/tkehayov/product-bg.git/repo"
	"io/ioutil"
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

func init() {
	log.SetFormatter(&log.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})
	log.SetOutput(os.Stdout)
}

func main() {
	log.Info("MERCHANT SERVICE STARTED")
	port := os.Getenv("PORT")

	router := mux.NewRouter()
	router.HandleFunc("/merchants", createNewUser).Methods("POST")

	log.Error(http.ListenAndServe(":"+port, router))
}

func createNewUser(w http.ResponseWriter, r *http.Request) {
	merchants := repo.Merchant{}
	body, _ := ioutil.ReadAll(r.Body)
	unmarshall(body, &merchants)

	log.Info(merchants)
	repo.Register(merchants)
}

func unmarshall(d []byte, merchant *repo.Merchant) {
	err := json.Unmarshal(d, &merchant)
	if err != nil {
		log.Error("error unmarshalling: ", err.Error())
		return
	}
}
