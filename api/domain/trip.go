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
