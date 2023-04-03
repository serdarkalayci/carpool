package dao

import "go.mongodb.org/mongo-driver/bson/primitive"

type RequestDAO struct {
	ID             primitive.ObjectID   `bson:"_id"`
	RequesterID    primitive.ObjectID   `bson:"requesterid"`
	RequesterName  string               `bson:"requestername"`
	CountryID      primitive.ObjectID   `bson:"countryid"`
	Origin         string               `bson:"origin"`
	Destination    string               `bson:"destination"`
	Dates          []primitive.DateTime `bson:"dates"`
	RequestedSeats int                  `bson:"requestedseats"`
	State          int                  `bson:"state"`
}
