package services

import (
	"product-bg/products/internal/entities"
	"product-bg/products/internal/repo"
)

type ProductFilterService interface {
	GetProducts(category string, filter map[string][]string) []entities.ProductFilter
}

type productFilter struct {
	ProductFilterRepository repo.ProductFilterRepositoryInterface
}

func NewProductFilterService(productFilterRepositoryInterface repo.ProductFilterRepositoryInterface) ProductFilterService {
	return &productFilter{ProductFilterRepository: productFilterRepositoryInterface}
}

func (productFilters productFilter) GetProducts(category string, filter map[string][]string) []entities.ProductFilter {

	categoryRepository := repo.NewCategoryRepository()
	categoryService := NewCategoryService(categoryRepository)
	categoryEntity := categoryService.GetOne(category)

	mapFilters := mapFilters(categoryEntity, filter)
	return productFilters.ProductFilterRepository.GetFilteredProducts(category, mapFilters)
}

func mapFilters(category entities.Category, filters map[string][]string) map[string][]string {
	results := make(map[string][]string)

	for name, value := range filters {
		for _, categoryFilter := range category.Filter {
			if name == categoryFilter.Name {
				results[categoryFilter.Value] = value
				break
			}

			//TODO dirty hack - should find another solution
			if name == "after" {
				results["after"] = value
			}

			if name == "before" {
				results["before"] = value
			}
		}
		//results[name] = value
	}

	return results
}
