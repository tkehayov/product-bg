package dto

import "product-bg/products/internal/entities"

type Product struct {
	Id           string  `json:"id"`
	Name         string  `json:"name"`
	FeatureImage string  `json:"featureImage"`
	MinPrice     float64 `json:"minPrice"`
}

func ParseProductFilterFromEntities(filters []entities.ProductFilter) []Product {
	var products []Product
	for _, productFilter := range filters {
		product := Product{
			Id:           productFilter.Id,
			Name:         productFilter.Name,
			FeatureImage: productFilter.Gallery.FeatureImage,
			MinPrice:     productFilter.MinPrice,
		}
		products = append(products, product)
	}
	return products

}
