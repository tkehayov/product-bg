package services

import (
	"product-bg/products/internal/entities"
	"product-bg/products/internal/repo"
)

type ProductCategoryFilterService interface {
	GetCategory(category string) entities.ProductCategoryFilter
}

type productCategoryFilter struct {
	ProductCategoryFilterRepository repo.ProductCategoryFilterRepositoryInterface
}

func NewCategoryService(
	productCategoryFilterRepository repo.ProductCategoryFilterRepositoryInterface,
) ProductCategoryFilterService {
	return &productCategoryFilter{ProductCategoryFilterRepository: productCategoryFilterRepository}
}

func (prodCatFilter productCategoryFilter) GetCategory(category string) entities.ProductCategoryFilter {
	return prodCatFilter.ProductCategoryFilterRepository.GetFilters(category)
}
