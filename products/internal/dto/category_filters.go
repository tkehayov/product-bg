package dto

import "product-bg/products/internal/entities"

type Filter struct {
	ID    string   `json:"id"`
	Name  string   `json:"name"`
	Value []string `json:"values"`
}

func ParseCategoryFilterFromEntities(filter entities.CategoryProductFilter, category entities.Category) []Filter {
	var filters []Filter

	for _, filter := range filter.ProductFilter {
		for _, categoryFilter := range category.Filter {
			if filter.Name == categoryFilter.Value {
				filterDto := Filter{ID: categoryFilter.Name, Name: filter.Name, Value: filter.Value}
				filters = append(filters, filterDto)
			}
		}
	}

	return filters
}
