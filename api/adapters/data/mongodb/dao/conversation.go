package dao

import "go.mongodb.org/mongo-driver/bson/primitive"

type ConversationDAO struct {
	ID                primitive.ObjectID `bson:"_id"`
	TripID            primitive.ObjectID `bson:"tripid"`
	RequesterID       primitive.ObjectID `bson:"requesterid"`
	RequesterName     string             `bson:"requestername"`
	SupplierID        primitive.ObjectID `bson:"supplierid"`
	SupplierName      string             `bson:"suppliername"`
	RequesterApproved bool               `bson:"requesterapprove"`
	SupplierApproved  bool               `bson:"supplierapprove"`
	Messages          []MessageDAO       `bson:"messages"`
}

type MessageDAO struct {
	Direction string             `bson:"direction"`
	Date      primitive.DateTime `bson:"date"`
	Text      string             `bson:"text"`
	Read      bool               `bson:"read"`
}
