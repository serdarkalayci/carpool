package mappers

import (
	"time"

	"github.com/serdarkalayci/carpool/api/adapters/comm/rest/dto"
	"github.com/serdarkalayci/carpool/api/domain"
)

func MapTrip2TripListItem(trip domain.Trip) dto.TripListItem {
	return dto.TripListItem{
		ID:             trip.ID,
		Origin:         trip.Origin,
		Destination:    trip.Destination,
		TripDate:       trip.TripDate,
		AvailableSeats: trip.AvailableSeats,
	}
}

func MapTrips2TripListItems(trips []*domain.Trip) []dto.TripListItem {
	var tripListItems []dto.TripListItem
	for _, trip := range trips {
		tripListItems = append(tripListItems, MapTrip2TripListItem(*trip))
	}
	return tripListItems
}

func MapAddTripRequest2Trip(trip dto.AddTripRequest) (domain.Trip, error) {
	tripDate, err := time.Parse("2006-01-02", trip.TripDate)
	if err != nil {
		return domain.Trip{}, err
	}
	return domain.Trip{
		CountryID:      trip.CountryID,
		Origin:         trip.Origin,
		Destination:    trip.Destination,
		Stops:          trip.Stops,
		TripDate:       tripDate,
		AvailableSeats: trip.AvailableSeats,
		Note:           trip.Note,
	}, nil
}

func MapTripDetail2TripDetailResponse(tripDetail domain.TripDetail) dto.TripDetailResponse {
	return dto.TripDetailResponse{
		ID:             tripDetail.ID,
		SupplierName:   tripDetail.SupplierName,
		Country:        tripDetail.Country,
		Origin:         tripDetail.Origin,
		Destination:    tripDetail.Destination,
		Stops:          tripDetail.Stops,
		TripDate:       tripDetail.TripDate,
		AvailableSeats: tripDetail.AvailableSeats,
		Note:           tripDetail.Note,
		Conversation:   MapConversations2ConversationResponses(tripDetail.Conversations),
	}
}

func MapConversation2ConversationResponse(conversation domain.Conversation) dto.ConversationResponse {
	return dto.ConversationResponse{
		ConversationID: conversation.ConversationID,
		RequesterName:  conversation.RequesterName,
		Messages:       MapMessages2MessageResponses(conversation.Messages),
	}
}

func MapConversations2ConversationResponses(conversations []domain.Conversation) []dto.ConversationResponse {
	var conversationResponses []dto.ConversationResponse
	for _, conversation := range conversations {
		conversationResponses = append(conversationResponses, MapConversation2ConversationResponse(conversation))
	}
	return conversationResponses
}

func MapMessage2MessageResponse(message domain.Message) dto.MessageResponse {
	return dto.MessageResponse{
		Direction: message.Direction,
		Date:      message.Date,
		Text:      message.Text,
		Read:      message.Read,
	}
}

func MapMessages2MessageResponses(messages []domain.Message) []dto.MessageResponse {
	var messageResponses []dto.MessageResponse
	for _, message := range messages {
		messageResponses = append(messageResponses, MapMessage2MessageResponse(message))
	}
	return messageResponses
}
