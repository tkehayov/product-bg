package entities

type ProductCategoryFilter struct {
	ProductFilter []ProductFilter `bson:"filters"`
}

type ProductFilter struct {
	Name  string   `bson:"name"`
	Value []string `bson:"values"`
}
