package dto

import "product-bg/products/internal/entities"

type Filter struct {
	Name  string   `json:"name"`
	Value []string `json:"values"`
}

func ParseFromEntities(data entities.ProductCategoryFilter) []Filter {
	var filterDto []Filter
	for _, filter := range data.ProductFilter {
		filt := Filter{Name: filter.Name, Value: filter.Value}
		filterDto = append(filterDto, filt)
	}

	return filterDto

}
