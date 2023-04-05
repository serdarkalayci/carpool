// Package dto is the package that defines types for data transfer
package dto

import (
	"time"
)

// ConversationResponse is the response object for a conversation
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

// MessageResponse is the response object for a message
type MessageResponse struct {
	Direction string    `json:"direction"`
	Date      time.Time `json:"date"`
	Text      string    `json:"text"`
	Read      bool      `json:"read"`
}

// ContactDetails is the response object for contact details
type ContactDetails struct {
	Email string `json:"email"`
	Phone string `json:"phone"`
}

// AddConversationRequest is the request object for adding a message
type AddConversationRequest struct {
	TripID   string `json:"tripId" validate:"required"`
	Message  string `json:"message" validate:"required"`
	Capacity int    `json:"capacity" validate:"required"`
}

// UpdateApprovalRequest is the request object for updating the approval status
type UpdateApprovalRequest struct {
	Approved *bool `json:"approved" validate:"required"`
}
