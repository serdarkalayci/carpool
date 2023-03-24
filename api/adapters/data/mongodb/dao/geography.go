package dao

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CountryDAO struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	Name   string             `bson:"name"`
	Cities []CityDAO          `bson:"cities"`
}

type CityDAO struct {
	Name   string `bson:"name"`
	Ballot bool   `bson:"ballot"`
}
