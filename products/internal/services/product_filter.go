package services

import (
	log "github.com/sirupsen/logrus"
	"product-bg/products/internal/entities"
	"product-bg/products/internal/repo"
)

type ProductFilterService interface {
	GetProducts(category string, filter string) entities.ProductFilter
}

type productFilter struct {
	ProductFilterRepository repo.ProductFilterRepositoryInterface
}

func NewProductFilterService(productFilterRepositoryInterface repo.ProductFilterRepositoryInterface) ProductFilterService {
	return &productFilter{ProductFilterRepository: productFilterRepositoryInterface}
}

//TODO implement
func (p productFilter) GetProducts(category string, filter string) entities.ProductFilter {
	log.Error(category)
	log.Error(filter)
	log.Error("implement me")
	return entities.ProductFilter{}
}
