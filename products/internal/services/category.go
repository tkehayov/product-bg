package services

import (
	"product-bg/products/internal/entities"
	"product-bg/products/internal/repo"
)

type CategoryService interface {
	GetOne(category string) entities.Category
}

type category struct {
	CategoryRepository repo.CategoryRepositoryInterface
}

func NewCategoryService(
	productCategoryFilterRepository repo.CategoryRepositoryInterface,
) CategoryService {
	return &category{CategoryRepository: productCategoryFilterRepository}
}

func (category category) GetOne(id string) entities.Category {
	return category.CategoryRepository.GetOne(id)
}
