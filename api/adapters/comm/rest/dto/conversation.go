package dto

import (
	"time"
)

type ConversationResponse struct {
	ConversationID    string            `json:"conversationid"`
	RequesterName     string            `json:"requestername"`
	RequesterApproved bool              `json:"requesterapproved"`
	SupplierApproved  bool              `json:"supplierapproved"`
	RequestedCapacity int               `json:"requestedcapacity"`
	RequesterContact  ContactDetails    `json:"requestercontact,omitempty"`
	SupplierContact   ContactDetails    `json:"suppliercontact,omitempty"`
	Messages          []MessageResponse `json:"messages"`
}

type MessageResponse struct {
	Direction string    `json:"direction"`
	Date      time.Time `json:"date"`
	Text      string    `json:"text"`
	Read      bool      `json:"read"`
}

type ContactDetails struct {
	Email string `json:"email"`
	Phone string `json:"phone"`
}

type AddConversationRequest struct {
	TripID   string `json:"tripId" validate:"required"`
	Message  string `json:"message" validate:"required"`
	Capacity int    `json:"capacity" validate:"required"`
}

type UpdateApprovalRequest struct {
	Approved *bool `json:"approved" validate:"required"`
}
