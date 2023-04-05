// Package domain is the package that holds the very basic domain objects
package domain

import (
	"time"

	"github.com/spf13/viper"
)

// Request is the domain object for a request
type Request struct {
	ID             string
	RequesterID    string
	RequesterName  string
	CountryID      string
	Origin         string
	Destination    string
	RequestedSeats int
	Dates          []time.Time
	State          RequestState
}

// RequestState is the state of a request
type RequestState int

const (
	// Requested is the initial state of a request, nothing has been done yet
	Requested RequestState = iota
	// BeingServed is the state of a request when one or more conversations are open
	BeingServed
	// Completed is the state of a request when at least one conversation has been completed
	Completed
)

// String returns the string representation of a RequestState
func (r RequestState) String() string {
	return viper.GetStringSlice("RequestStates")[r]
}
