package dto

import "time"

type TripListItem struct {
	ID             string    `json:"id"`
	Origin         string    `json:"origin"`
	Destination    string    `json:"destination"`
	TripDate       time.Time `json:"tripdate"`
	AvailableSeats int       `json:"availableseats"`
}

type AddTripRequest struct {
	CountryID      string   `json:"countryid" validate:"required"`
	Origin         string   `json:"origin" validate:"required"`
	Destination    string   `json:"destination" validate:"required"`
	Stops          []string `json:"stops"`
	TripDate       string   `json:"tripdate" validate:"required"`
	AvailableSeats int      `json:"availableseats" validate:"required"`
	Note           string   `json:"note"`
}
