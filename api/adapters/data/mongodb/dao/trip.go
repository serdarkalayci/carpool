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

type ConversationDAO struct {
	ID            primitive.ObjectID `bson:"_id"`
	TripID        primitive.ObjectID `bson:"tripid"`
	RequesterID   primitive.ObjectID `bson:"requesterid"`
	RequesterName string             `bson:"requestername"`
	Messages      []MessageDAO       `bson:"messages"`
}

type MessageDAO struct {
	Direction string             `bson:"direction"`
	Date      primitive.DateTime `bson:"date"`
	Text      string             `bson:"text"`
	Read      bool               `bson:"read"`
}
