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

type ConversationResponse struct {
	ConversationID string            `json:"conversationid"`
	RequesterName  string            `json:"requestername"`
	Messages       []MessageResponse `json:"messages"`
}

type MessageResponse struct {
	Direction string    `json:"direction"`
	Date      time.Time `json:"date"`
	Text      string    `json:"text"`
	Read      bool      `json:"read"`
}

type AddMessageRequest struct {
	Text string `json:"text" validate:"required"`
}
