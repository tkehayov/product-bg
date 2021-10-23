package entities

type CategoryProductFilter struct {
	ProductFilter []ProductProperty `bson:"filters"`
}

type ProductProperty struct {
	Name  string   `bson:"name"`
	Value []string `bson:"values"`
}
