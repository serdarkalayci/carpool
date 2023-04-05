// Package errors is the package that holds the custom application errors
package errors

// ErrNotAuthorizedForConversation is returned when a user is neither the supplier nor the requester of a conversation.
type ErrNotAuthorizedForConversation struct{}

func (e ErrNotAuthorizedForConversation) Error() string {
	return "this user is not authorized to see this conversation"
}

// ErrConversationNotFound is returned when a conversation is not found.
type ErrConversationNotFound struct{}

func (e ErrConversationNotFound) Error() string {
	return "conversation not found"
}

// ErrConversationNotInserted is returned when a conversation is not inserted.
type ErrConversationNotInserted struct{}

func (e ErrConversationNotInserted) Error() string {
	return "conversation not inserted"
}

// ErrConversationNotUpdated is returned when a conversation is not updated.
type ErrConversationNotUpdated struct{}

func (e ErrConversationNotUpdated) Error() string {
	return "conversation not updated"
}

// ErrMessageNotInserted is returned when a message is not inserted to the conversation.
type ErrMessageNotInserted struct{}

func (e ErrMessageNotInserted) Error() string {
	return "message not inserted"
}

// ErrMessageNotUpdated is returned when a message is not updated.
type ErrMessageNotUpdated struct{}

func (e ErrMessageNotUpdated) Error() string {
	return "message not updated"
}

// ErrInvalidCapacity is returned when a requested capacity is greater than the available capacity.
type ErrInvalidCapacity struct{}

func (e ErrInvalidCapacity) Error() string {
	return "requested capacity is greater than available capacity"
}
