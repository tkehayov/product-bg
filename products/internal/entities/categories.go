package entities

type Category struct {
	ID     string   `bson:"_id"`
	Name   string   `bson:"name"`
	Filter []Filter `bson:"filters"`
}

type Filter struct {
	Name  string `bson:"name"`
	Value string `bson:"value"`
}
