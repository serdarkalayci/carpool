package domain

import "time"

type Trip struct {
	ID             string
	SupplierID     string
	CountryID      string
	Origin         string
	Destination    string
	Stops          []string
	TripDate       time.Time
	AvailableSeats int
	Note           string
}

type TripDetail struct {
	ID             string
	SupplierID     string
	SupplierName   string
	Country        string
	Origin         string
	Destination    string
	Stops          []string
	TripDate       time.Time
	AvailableSeats int
	Note           string
	Conversations  []Conversation
}

type Conversation struct {
	ConversationID string
	RequesterName  string
	Messages       []Message
}

type Message struct {
	Direction string
	Date      time.Time
	Text      string
	Read      bool
}

type ErrNotTheOwner struct{}
type ErrTheOwner struct{}

func (e ErrNotTheOwner) Error() string {
	return "this user is not the supplier of this trip"
}

func (e ErrTheOwner) Error() string {
	return "this user is the supplier of this trip, therefore cannot inititate conversation"
}
