package services

import (
	"product-bg/products/internal/entities"
	"product-bg/products/internal/repo"
)

type ProductCategoryFilterService interface {
	GetCategoryFilters(category string) entities.CategoryProductFilter
}

type productCategoryFilter struct {
	ProductCategoryFilterRepository repo.ProductCategoryFilterRepositoryInterface
}

func NewProductCategoryService(
	productCategoryFilterRepository repo.ProductCategoryFilterRepositoryInterface,
) ProductCategoryFilterService {
	return &productCategoryFilter{ProductCategoryFilterRepository: productCategoryFilterRepository}
}

func (prodCatFilter productCategoryFilter) GetCategoryFilters(category string) entities.CategoryProductFilter {
	return prodCatFilter.ProductCategoryFilterRepository.GetFilters(category)
}
