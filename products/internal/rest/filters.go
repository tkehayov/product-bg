package rest

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"product-bg/products/internal/repo"
)

type Filter struct {
}

func (Filter) GetAll(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	category := params["category"]
	categoryEntity := repo.GetFilters(category)

	response, err := json.Marshal(categoryEntity)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
