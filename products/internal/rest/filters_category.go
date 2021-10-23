package rest

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"product-bg/products/internal/dto"
	"product-bg/products/internal/repo"
	"product-bg/products/internal/services"
)

type FilterCategory struct {
}

func (FilterCategory) GetAll(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	category := params["category"]
	productCategoryFilterRepository := repo.NewProductCategoryFilterRepository()
	categoryRepository := repo.NewCategoryRepository()
	categoryFilterEntity := services.NewProductCategoryService(productCategoryFilterRepository).GetCategoryFilters(category)
	categoryService := services.NewCategoryService(categoryRepository)
	categoryEntity := categoryService.GetOne(category)
	dto := dto.ParseCategoryFilterFromEntities(categoryFilterEntity, categoryEntity)
	response, err := json.Marshal(dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
