// Package domain is the package that holds the very basic domain objects
package domain

import "time"

// Conversation is the domain object for a conversation
type Conversation struct {
	TripID            string
	RequestID         string
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

// Message is the domain object for a message
type Message struct {
	Direction string
	Date      time.Time
	Text      string
	Read      bool
}

// ContactDetails is the domain object for contact details
type ContactDetails struct {
	Email string
	Phone string
}
