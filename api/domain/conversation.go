package domain

import "time"

type Conversation struct {
	TripID            string
	ConversationID    string
	RequesterID       string
	RequesterName     string
	SupplierID        string
	SupplierName      string
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

type ErrNotTheOwner struct{}
type ErrTheOwner struct{}

func (e ErrNotTheOwner) Error() string {
	return "this user is not the supplier of this trip"
}

func (e ErrTheOwner) Error() string {
	return "this user is the supplier of this trip, therefore cannot inititate conversation"
}
