package entities

type ProductFilter struct {
	Id      string  `bson:"_id"`
	Name    string  `bson:"name"`
	Gallery Gallery `bson:"gallery"`
	//MinPrice     string `bson:"minPrice"`
}

type Gallery struct {
	FeatureImage string `bson:"featureImage"`
}
