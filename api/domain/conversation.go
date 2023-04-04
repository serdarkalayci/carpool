package domain

import "time"

type Conversation struct {
	TripID            string
	ConversationID    string
	RequesterID       string
	RequesterName     string
	SupplierID        string
	SupplierName      string
	RequestedCapacity int
	RequesterApproved bool
	SupplierApproved  bool
	RequesterContact  ContactDetails
	SupplierContact   ContactDetails
	Messages          []Message
}

type Message struct {
	Direction string
	Date      time.Time
	Text      string
	Read      bool
}

type ContactDetails struct {
	Email string
	Phone string
}
