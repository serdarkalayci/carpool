// Package domain is the package that holds the very basic domain objects
package domain

import "time"

// Trip is the domain object for a trip
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

// TripDetail is the domain object for a trip with conversations
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
