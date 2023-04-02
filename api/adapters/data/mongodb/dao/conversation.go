package dao

import "go.mongodb.org/mongo-driver/bson/primitive"

type ConversationDAO struct {
	ID                primitive.ObjectID `bson:"_id"`
	TripID            primitive.ObjectID `bson:"tripid"`
	RequesterID       primitive.ObjectID `bson:"requesterid"`
	RequesterName     string             `bson:"requestername"`
	SupplierID        primitive.ObjectID `bson:"supplierid"`
	SupplierName      string             `bson:"suppliername"`
	RequestedCapacity int                `bson:"requestedcapacity"`
	RequesterApproved bool               `bson:"requesterapprove"`
	SupplierApproved  bool               `bson:"supplierapprove"`
	RequesterContact  ContactDetails     `bson:"requestercontact"`
	SupplierContact   ContactDetails     `bson:"suppliercontact"`
	Messages          []MessageDAO       `bson:"messages"`
}

type MessageDAO struct {
	Direction string             `bson:"direction"`
	Date      primitive.DateTime `bson:"date"`
	Text      string             `bson:"text"`
	Read      bool               `bson:"read"`
}

type ContactDetails struct {
	Email string `bson:"email"`
	Phone string `bson:"phone"`
}
