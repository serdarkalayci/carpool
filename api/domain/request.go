package domain

import (
	"time"

	"github.com/spf13/viper"
)

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

type RequestState int

const (
	Requested RequestState = iota
	BeingServed
	Completed
)

func (r RequestState) String() string {
	return viper.GetStringSlice("RequestStates")[r]
}
