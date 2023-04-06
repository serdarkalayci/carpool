package application

import "github.com/serdarkalayci/carpool/api/domain"

type mockConversationRepository struct{}

var (
	checkConversationOwnershipFunc func(conversationID string, userID string) (bool, error)
	initiateConversationFunc       func(domain.Conversation) error
	addMessageFunc                 func(conversationID string, message string, direction string) error
	getConversationFunc            func(tripID string, userID string) (*domain.Conversation, error)
	getConversationByIDFunc        func(conversationID string) (*domain.Conversation, error)
	getConversationsFunc           func(tripID string) ([]domain.Conversation, error)
	markConversationsReadFunc      func(conversationID string, direction string) error
	updateApprovalFunc             func(conversationID string, supplierApprove string, requesterApprove string) error
)

// CheckConversationOwnership checks if a conversation's supplier is the current user
func (m mockConversationRepository) CheckConversationOwnership(conversationID string, userID string) (bool, error) {
	return checkConversationOwnershipFunc(conversationID, userID)
}

// InitiateConversation initiates a conversation
func (m mockConversationRepository) InitiateConversation(conversation domain.Conversation) error {
	return initiateConversationFunc(conversation)
}

// AddMessage adds a message to a conversation
func (m mockConversationRepository) AddMessage(conversationID string, message string, direction string) error {
	return addMessageFunc(conversationID, message, direction)
}

// GetConversation gets a conversation based on the tripID and userID
func (m mockConversationRepository) GetConversation(tripID string, userID string) (*domain.Conversation, error) {
	return getConversationFunc(tripID, userID)
}

// GetConversationByID gets a conversation based on the conversationID
func (m mockConversationRepository) GetConversationByID(conversationID string) (*domain.Conversation, error) {
	return getConversationByIDFunc(conversationID)
}

// GetConversations gets all conversations based on the tripID
func (m mockConversationRepository) GetConversations(tripID string) ([]domain.Conversation, error) {
	return getConversationsFunc(tripID)
}

// MarkConversationsRead marks a conversation as read
func (m mockConversationRepository) MarkConversationsRead(conversationID string, direction string) error {
	return markConversationsReadFunc(conversationID, direction)
}

// UpdateApproval updates the approval status of a conversation
func (m mockConversationRepository) UpdateApproval(conversationID string, supplierApprove string, requesterApprove string) error {
	return updateApprovalFunc(conversationID, supplierApprove, requesterApprove)
}
