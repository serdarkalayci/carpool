package dto

type AddConversationRequest struct {
	TripID  string `json:"tripId" validate:"required"`
	Message string `json:"message" validate:"required"`
}

type UpdateApprovalRequest struct {
	Approved *bool `json:"approved" validate:"required"`
}
