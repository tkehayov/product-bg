package repo

import (
	"context"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"product-bg/merchants/internal/database"
)

type Merchant struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Website  string `json:"website"`
}

type MerchantEntity struct {
	ID       primitive.ObjectID `bson:"_id, omitempty"`
	Username string             `json:"username"`
	Password string             `json:"password"`
	Website  string             `json:"website"`
}

type MerchantLogo struct {
	Logo string `bson:"logo" json:"logo"`
}

func Register(m Merchant) {
	client, ctx := database.Connect()
	defer client.Disconnect(ctx)

	db := client.Database("merchants")

	collection := db.Collection("merchants")

	_, err := collection.InsertOne(ctx, m)
	if err != nil {
		log.Error("Couldn't insert merchant: ", err.Error())
	}
}

func GetLogo(id string) MerchantLogo {
	var merchantLogo MerchantLogo
	mId, errMerchant := primitive.ObjectIDFromHex(id)
	if errMerchant != nil {
		log.Error("error parsing merchant: ", errMerchant.Error())
	}
	client, ctx := database.Connect()
	defer client.Disconnect(ctx)
	db := client.Database("merchants")
	collection := db.Collection("merchants")

	criteria := bson.D{{"_id", mId}}
	err := collection.FindOne(context.TODO(), criteria).Decode(&merchantLogo)
	if err != nil {
		log.Error("Error finding merchant logo: ", err.Error())
	}

	return merchantLogo
}

func CredentialsMatch(m Merchant) bool {
	client, ctx := database.Connect()
	defer client.Disconnect(ctx)

	db := client.Database("merchants")
	collection := db.Collection("merchants")
	criteria := bson.D{{"username", m.Username}, {"password", m.Password}}

	var result Merchant
	err := collection.FindOne(context.TODO(), criteria).Decode(&result)

	return err == nil
}
