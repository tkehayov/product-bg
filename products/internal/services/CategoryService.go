package services

import "product-bg/products/internal/entities"

type CategoryService interface {
	GetCategory() []entities.Category
}

type categoryService struct {
}
