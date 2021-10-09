package dto

type Category struct {
	Filter []Filter `bson:"filters" json:"filters"`
}

type Filter struct {
	Name  string   `bson:"name" json:"name"`
	Value []string `bson:"values" json:"values"`
}
