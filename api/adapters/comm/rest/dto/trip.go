// Package dto is the package that defines types for data transfer
package dto

import "time"

// TripListItem is the response object for a trip, which lacks some fields
type TripListItem struct {
	ID             string    `json:"id"`
	Origin         string    `json:"origin"`
	Destination    string    `json:"destination"`
	TripDate       time.Time `json:"tripdate"`
	AvailableSeats int       `json:"availableseats"`
}

// AddTripRequest is the request object for adding a trip
type AddTripRequest struct {
	CountryID      string   `json:"countryid" validate:"required"`
	Origin         string   `json:"origin" validate:"required"`
	Destination    string   `json:"destination" validate:"required"`
	Stops          []string `json:"stops"`
	TripDate       string   `json:"tripdate" validate:"required"`
	AvailableSeats int      `json:"availableseats" validate:"required"`
	Note           string   `json:"note"`
}

// TripDetailResponse is the response object for a trip in full
type TripDetailResponse struct {
	ID             string                 `json:"id"`
	SupplierName   string                 `json:"suppliername"`
	Country        string                 `json:"country"`
	Origin         string                 `json:"origin"`
	Destination    string                 `json:"destination"`
	Stops          []string               `json:"stops"`
	TripDate       time.Time              `json:"tripdate"`
	AvailableSeats int                    `json:"availableseats"`
	Note           string                 `json:"note"`
	Conversation   []ConversationResponse `json:"conversation,omitempty"`
}

// AddMessageRequest is the request object for adding a message
type AddMessageRequest struct {
	Text string `json:"text" validate:"required"`
}
