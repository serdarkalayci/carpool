package dao

import "go.mongodb.org/mongo-driver/bson/primitive"

type TripDAO struct {
	ID             primitive.ObjectID `bson:"_id"`
	SupplierID     primitive.ObjectID `bson:"supplierid"`
	CountryID      primitive.ObjectID `bson:"countryid"`
	Origin         string             `bson:"origin"`
	Destination    string             `bson:"destination"`
	Stops          []string           `bson:"stops"`
	TripDate       primitive.DateTime `bson:"tripdate"`
	AvailableSeats int                `bson:"availableseats"`
	Note           string             `bson:"note"`
}
