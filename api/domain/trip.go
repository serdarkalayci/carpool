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
	RequesterID string
	Messages    []Message
}

type Message struct {
	Direction string
	Date      time.Time
	Text      string
	Read      bool
}
