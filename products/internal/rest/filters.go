package rest

import (
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type Filter struct {
}

func (Filter) GetAll(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	category := params["category"]

	log.Error(category)
}
