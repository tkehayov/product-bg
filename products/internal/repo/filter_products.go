package repo

import (
	log "github.com/sirupsen/logrus"
	"product-bg/products/internal/entities"
)

type ProductFilterRepositoryInterface interface {
	GetFilters(category string) entities.ProductFilter
}

type productFilter struct{}

func NewProductFilterRepository() ProductFilterRepositoryInterface {
	return &productFilter{}
}

func (p productFilter) GetFilters(category string) entities.ProductFilter {
	log.Errorln("TODO REPO", category)

	return entities.ProductFilter{}
}
