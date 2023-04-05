// Package dao is the package that holds the database access objects
package dao

import "go.mongodb.org/mongo-driver/bson/primitive"

// TripDAO is the slim data access object for the Trip entity. This object just contains with foreign keys, without the relations
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

// TripDetailDAO is the data access object for the Trip entity. This object contains the relations
type TripDetailDAO struct {
	ID             primitive.ObjectID `bson:"_id"`
	SupplierID     primitive.ObjectID `bson:"supplierid"`
	SupplierName   string             `bson:"username"`
	Country        string             `bson:"countryname"`
	Origin         string             `bson:"origin"`
	Destination    string             `bson:"destination"`
	Stops          []string           `bson:"stops"`
	TripDate       primitive.DateTime `bson:"tripdate"`
	AvailableSeats int                `bson:"availableseats"`
	Note           string             `bson:"note"`
}
