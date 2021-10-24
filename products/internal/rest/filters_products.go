package rest

import (
	"encoding/json"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
	"net/url"
	"product-bg/products/internal/dto"
	"product-bg/products/internal/repo"
	"product-bg/products/internal/services"
)

type FilterProduct struct {
}

func (FilterProduct) Get(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	//TODO YOU ARE HERE

	category := params["category"]
	filters := r.URL.Query()
	repository := repo.NewProductFilterRepository()
	productEntity := services.NewProductFilterService(repository).GetProducts(category, filters)

	dto := dto.ParseProductFilterFromEntities(productEntity)
	response, err := json.Marshal(dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func DecodeBase64(queryStr2 string) {
	params, _ := url.ParseQuery(queryStr2)
	log.Errorln(params)
}
