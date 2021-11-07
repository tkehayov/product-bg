package dto

import "product-bg/products/internal/entities"

type Product struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	FeatureImage string `json:"featureImage"`
	MinPrice     string `json:"minPrice"`
}

func ParseProductFilterFromEntities(filters []entities.ProductFilter) []Product {
	var products []Product
	for _, productFilter := range filters {
		product := Product{
			Id:           productFilter.Id,
			Name:         productFilter.Name,
			FeatureImage: productFilter.Gallery.FeatureImage,
			MinPrice:     "todo",
		}
		products = append(products, product)
	}
	return products

}
