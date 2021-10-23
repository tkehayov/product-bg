package dto

import "product-bg/products/internal/entities"

type Product struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	FeatureImage string `json:"featureImage"`
	MinPrice     string `json:"minPrice"`
}

func ParseProductFilterFromEntities(data entities.ProductFilter) []Product {
	//TODO implement

	return []Product{}

}
