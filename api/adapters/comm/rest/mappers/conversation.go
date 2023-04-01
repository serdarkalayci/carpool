package mappers

import (
	"github.com/serdarkalayci/carpool/api/adapters/comm/rest/dto"
	"github.com/serdarkalayci/carpool/api/domain"
)

func MapConversation2ConversationResponse(conversation domain.Conversation) dto.ConversationResponse {
	return dto.ConversationResponse{
		ConversationID:    conversation.ConversationID,
		RequesterName:     conversation.RequesterName,
		RequesterApproved: conversation.RequesterApproved,
		SupplierApproved:  conversation.SupplierApproved,
		SupplierContact: dto.ContactDetails{
			Email: conversation.SupplierContact.Email,
			Phone: conversation.SupplierContact.Phone,
		},
		RequesterContact: dto.ContactDetails{
			Email: conversation.RequesterContact.Email,
			Phone: conversation.RequesterContact.Phone,
		},
		Messages: MapMessages2MessageResponses(conversation.Messages),
	}
}
