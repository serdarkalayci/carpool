// Package dao is the package that holds the database access objects
package dao

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CountryDAO is the data access object for the Country entity
type CountryDAO struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	Name   string             `bson:"name"`
	Cities []CityDAO          `bson:"cities"`
}

// CityDAO is the data access object for the City entity
type CityDAO struct {
	Name   string `bson:"name"`
	Ballot bool   `bson:"ballot"`
}
