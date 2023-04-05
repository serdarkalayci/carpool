// Package dto is the package that defines types for data transfer
package dto

// AddRequestRequest is the request object for adding a request
type AddRequestRequest struct {
	CountryID      string   `json:"countryid" validate:"required"`
	Origin         string   `json:"origin" validate:"required"`
	Destination    string   `json:"destination" validate:"required"`
	RequestedSeats int      `json:"requestedseats" validate:"required"`
	Dates          []string `json:"dates" validate:"required"`
}

// RequestListResponse is the response object for a list of requests, which lacks some fields
type RequestListResponse struct {
	ID             string   `json:"id"`
	Origin         string   `json:"origin"`
	Destination    string   `json:"destination"`
	RequestedSeats int      `json:"requestedseats"`
	Dates          []string `json:"dates"`
	State          string
}

// RequestResponse is the response object for a request in full
type RequestResponse struct {
	ID             string   `json:"id"`
	RequesterID    string   `json:"requesterid"`
	RequesterName  string   `json:"requestername"`
	CountryID      string   `json:"countryid"`
	Origin         string   `json:"origin"`
	Destination    string   `json:"destination"`
	RequestedSeats int      `json:"requestedseats"`
	Dates          []string `json:"dates"`
	State          string   `json:"state"`
}
