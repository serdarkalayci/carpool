// Package mappers is the package that maps objects back and fort between dto and domain
package mappers

import (
	"github.com/serdarkalayci/carpool/api/adapters/comm/rest/dto"
	"github.com/serdarkalayci/carpool/api/domain"
)

// MapConversation2ConversationResponse maps a domain.Conversation to a dto.ConversationResponse
func MapConversation2ConversationResponse(conversation domain.Conversation) dto.ConversationResponse {
	return dto.ConversationResponse{
		ConversationID:    conversation.ConversationID,
		RequesterName:     conversation.RequesterName,
		RequestedCapacity: conversation.RequestedCapacity,
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
