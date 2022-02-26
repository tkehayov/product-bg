package rest

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"product-bg/products/internal/dto"
	"product-bg/products/internal/repo"
	"product-bg/products/internal/services"
)

type FilterProduct struct {
}

func (FilterProduct) Get(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	category := params["category"]
	filters := r.URL.Query()
	repository := repo.NewProductFilterRepository()
	productEntity := services.NewProductFilterService(repository).GetProducts(category, filters)

	products := dto.ParseProductFilterFromEntities(productEntity)

	if products == nil {
		products = []dto.Product{}
	}

	response, err := json.Marshal(products)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
