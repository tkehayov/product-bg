package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"os"
	"product-bg/merchants/repo"
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
	router.HandleFunc("/merchants/session", login).Methods("POST")

	log.Error(http.ListenAndServe(":"+port, router))
}

func createNewUser(w http.ResponseWriter, r *http.Request) {
	merchant := repo.Merchant{}
	body, _ := ioutil.ReadAll(r.Body)
	unmarshall(body, &merchant)

	repo.Register(merchant)
}

func login(w http.ResponseWriter, r *http.Request) {
	merchant := repo.Merchant{}
	body, _ := ioutil.ReadAll(r.Body)
	unmarshall(body, &merchant)

	match := repo.CredentialsMatch(merchant)
	//TODO CONTINUE WITH MATCH
	log.Error(match)
}
func unmarshall(d []byte, merchant *repo.Merchant) {
	err := json.Unmarshal(d, &merchant)
	if err != nil {
		log.Error("error unmarshalling: ", err.Error())
		return
	}
}
