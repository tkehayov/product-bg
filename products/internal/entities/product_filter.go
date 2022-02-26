package entities

//TODO rename
type ProductFilter struct {
	Id        string     `bson:"_id"`
	Name      string     `bson:"name"`
	Gallery   Gallery    `bson:"gallery"`
	Merchants []Merchant `bson:"merchants"`
	MinPrice  float64    `bson:"minPrice"`
}

type Merchant struct {
	Price float64 `bson:"price"`
}

type Gallery struct {
	FeatureImage string `bson:"featureImage"`
}
