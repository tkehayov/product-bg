package services

import (
	log "github.com/sirupsen/logrus"
	"product-bg/products/internal/entities"
	"product-bg/products/internal/repo"
)

type ProductFilterService interface {
	GetProducts(category string, filter map[string][]string) entities.ProductFilter
}

type productFilter struct {
	ProductFilterRepository repo.ProductFilterRepositoryInterface
}

func NewProductFilterService(productFilterRepositoryInterface repo.ProductFilterRepositoryInterface) ProductFilterService {
	return &productFilter{ProductFilterRepository: productFilterRepositoryInterface}
}

//TODO implement
func (p productFilter) GetProducts(category string, filter map[string][]string) entities.ProductFilter {
	log.Error("filterss: ", filter)
	return entities.ProductFilter{}
}
