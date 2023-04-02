package mappers

import (
	"github.com/serdarkalayci/carpool/api/adapters/data/mongodb/dao"
	"github.com/serdarkalayci/carpool/api/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func MapConversationDAO2Conversation(conversationDAO *dao.ConversationDAO) *domain.Conversation {
	messages := MapMessageDAOs2Messages(conversationDAO.Messages)
	return &domain.Conversation{
		TripID:            conversationDAO.TripID.Hex(),
		ConversationID:    conversationDAO.ID.Hex(),
		RequesterID:       conversationDAO.RequesterID.Hex(),
		RequesterName:     conversationDAO.RequesterName,
		SupplierID:        conversationDAO.SupplierID.Hex(),
		SupplierName:      conversationDAO.SupplierName,
		RequestedCapacity: conversationDAO.RequestedCapacity,
		RequesterApproved: conversationDAO.RequesterApproved,
		SupplierApproved:  conversationDAO.SupplierApproved,
		RequesterContact:  domain.ContactDetails{Email: conversationDAO.RequesterContact.Email, Phone: conversationDAO.RequesterContact.Phone},
		SupplierContact:   domain.ContactDetails{Email: conversationDAO.SupplierContact.Email, Phone: conversationDAO.SupplierContact.Phone},
		Messages:          messages,
	}
}
func MapConversation2ConversationDAO(conversation *domain.Conversation) *dao.ConversationDAO {
	messages := MapMessages2MessageDAOs(conversation.Messages)
	tripID, _ := primitive.ObjectIDFromHex(conversation.TripID)
	requesterID, _ := primitive.ObjectIDFromHex(conversation.RequesterID)
	supplierID, _ := primitive.ObjectIDFromHex(conversation.SupplierID)

	return &dao.ConversationDAO{
		ID:                primitive.NewObjectID(),
		TripID:            tripID,
		RequesterID:       requesterID,
		RequesterName:     conversation.RequesterName,
		SupplierID:        supplierID,
		SupplierName:      conversation.SupplierName,
		RequestedCapacity: conversation.RequestedCapacity,
		RequesterApproved: conversation.RequesterApproved,
		SupplierApproved:  conversation.SupplierApproved,
		RequesterContact:  dao.ContactDetails{Email: conversation.RequesterContact.Email, Phone: conversation.RequesterContact.Phone},
		SupplierContact:   dao.ContactDetails{Email: conversation.SupplierContact.Email, Phone: conversation.SupplierContact.Phone},
		Messages:          messages,
	}
}

func MapConversationDAOs2Conversations(conversationsDAO []dao.ConversationDAO) []domain.Conversation {
	var conversations []domain.Conversation
	for _, conversationDAO := range conversationsDAO {
		conversations = append(conversations, *MapConversationDAO2Conversation(&conversationDAO))
	}
	return conversations
}

func MapMessageDAO2Message(messageDAO *dao.MessageDAO) *domain.Message {
	return &domain.Message{
		Direction: messageDAO.Direction,
		Date:      messageDAO.Date.Time(),
		Text:      messageDAO.Text,
		Read:      messageDAO.Read,
	}
}

func MapMessageDAOs2Messages(messagesDAO []dao.MessageDAO) []domain.Message {
	var messages []domain.Message
	for _, messageDAO := range messagesDAO {
		messages = append(messages, *MapMessageDAO2Message(&messageDAO))
	}
	return messages
}

func MapMessage2MessageDAO(message *domain.Message) *dao.MessageDAO {
	return &dao.MessageDAO{
		Direction: message.Direction,
		Date:      primitive.NewDateTimeFromTime(message.Date),
		Text:      message.Text,
		Read:      message.Read,
	}
}

func MapMessages2MessageDAOs(messages []domain.Message) []dao.MessageDAO {
	var messagesDAO []dao.MessageDAO
	for _, message := range messages {
		messagesDAO = append(messagesDAO, *MapMessage2MessageDAO(&message))
	}
	return messagesDAO
}
