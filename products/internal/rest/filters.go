package rest

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"product-bg/products/internal/dto"
	"product-bg/products/internal/repo"
	"product-bg/products/internal/services"
)

type Filter struct {
}

func (Filter) GetAll(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	category := params["category"]
	repository := repo.NewProductCategoryFilterRepository()
	categoryEntity := services.NewCategoryService(repository).GetCategory(category)
	entities := dto.ParseFromEntities(categoryEntity)
	response, err := json.Marshal(entities)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
