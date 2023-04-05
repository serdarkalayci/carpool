// Package dao is the package that holds the database access objects
package dao

import "go.mongodb.org/mongo-driver/bson/primitive"

// ConversationDAO is the data access object for the conversation object.
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

// MessageDAO is the data access object for the message object.
type MessageDAO struct {
	Direction string             `bson:"direction"`
	Date      primitive.DateTime `bson:"date"`
	Text      string             `bson:"text"`
	Read      bool               `bson:"read"`
}

// ContactDetails is the data access object for the contact details object which is used for both supplier and requester.
type ContactDetails struct {
	Email string `bson:"email"`
	Phone string `bson:"phone"`
}
