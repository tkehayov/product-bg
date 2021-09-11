package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"os"
	"product-bg/merchants/internal/database"
)

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
	router.HandleFunc("/merchants/logo/{id}", logo).Methods("GET")

	log.Error(http.ListenAndServe(":"+port, router))
}

func createNewUser(w http.ResponseWriter, r *http.Request) {
	merchant := database.Merchant{}
	body, _ := ioutil.ReadAll(r.Body)
	unmarshall(body, &merchant)

	database.Register(merchant)
}

func login(w http.ResponseWriter, r *http.Request) {
	merchant := database.Merchant{}
	body, _ := ioutil.ReadAll(r.Body)
	unmarshall(body, &merchant)

	match := database.CredentialsMatch(merchant)
	if match {
		userCookie := generateUserCookie(merchant)
		usernameCookie := generateUserNameCookie(merchant)
		http.SetCookie(w, userCookie)
		http.SetCookie(w, usernameCookie)
		return
	}
	w.WriteHeader(http.StatusUnauthorized)

}

func logo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	m := database.GetLogo(id)

	response, err := json.Marshal(m)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func generateUserNameCookie(m database.Merchant) *http.Cookie {
	return &http.Cookie{
		Name:   "username",
		Value:  m.Username,
		MaxAge: 86400,
	}
}

func generateUserCookie(m database.Merchant) *http.Cookie {
	userpass := m.Username + m.Password
	h := sha256.New()
	h.Write([]byte(userpass))
	hashString := hex.EncodeToString(h.Sum(nil))

	return &http.Cookie{
		Name:   "logged",
		Value:  hashString,
		MaxAge: 86400,
	}
}

func unmarshall(d []byte, merchant *database.Merchant) {
	err := json.Unmarshal(d, &merchant)
	if err != nil {
		log.Error("error unmarshalling: ", err.Error())
		return
	}
}
