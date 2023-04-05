// Package mappers is the package that maps objects back and fort between dto and domain
package mappers

import (
	"time"

	"github.com/rs/zerolog/log"

	"github.com/serdarkalayci/carpool/api/adapters/comm/rest/dto"
	apierr "github.com/serdarkalayci/carpool/api/adapters/comm/rest/errors"
	"github.com/serdarkalayci/carpool/api/domain"
)

// MapTrip2TripListItem maps a domain.Trip to a dto.TripListItem
func MapTrip2TripListItem(trip domain.Trip) dto.TripListItem {
	return dto.TripListItem{
		ID:             trip.ID,
		Origin:         trip.Origin,
		Destination:    trip.Destination,
		TripDate:       trip.TripDate,
		AvailableSeats: trip.AvailableSeats,
	}
}

// MapTrips2TripListItems maps a slice of domain.Trip to a slice of dto.TripListItem
func MapTrips2TripListItems(trips []*domain.Trip) []dto.TripListItem {
	var tripListItems []dto.TripListItem
	for _, trip := range trips {
		tripListItems = append(tripListItems, MapTrip2TripListItem(*trip))
	}
	return tripListItems
}

// MapAddTripRequest2Trip maps a dto.AddTripRequest to a domain.Trip
func MapAddTripRequest2Trip(trip dto.AddTripRequest) (domain.Trip, error) {
	tripDate, err := time.Parse("2006-01-02", trip.TripDate)
	if err != nil {
		log.Error().Err(err).Msg("error parsing trip date")
		return domain.Trip{}, &apierr.ErrInvalidDateFormat{Date: trip.TripDate}
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

// MapTripDetail2TripDetailResponse maps a domain.TripDetail to a dto.TripDetailResponse
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

// MapConversations2ConversationResponses maps a slice of domain.Conversation to a slice of dto.ConversationResponse
func MapConversations2ConversationResponses(conversations []domain.Conversation) []dto.ConversationResponse {
	var conversationResponses []dto.ConversationResponse
	for _, conversation := range conversations {
		conversationResponses = append(conversationResponses, MapConversation2ConversationResponse(conversation))
	}
	return conversationResponses
}

// MapMessage2MessageResponse maps a domain.Message to a dto.MessageResponse
func MapMessage2MessageResponse(message domain.Message) dto.MessageResponse {
	return dto.MessageResponse{
		Direction: message.Direction,
		Date:      message.Date,
		Text:      message.Text,
		Read:      message.Read,
	}
}

// MapMessages2MessageResponses maps a slice of domain.Message to a slice of dto.MessageResponse
func MapMessages2MessageResponses(messages []domain.Message) []dto.MessageResponse {
	var messageResponses []dto.MessageResponse
	for _, message := range messages {
		messageResponses = append(messageResponses, MapMessage2MessageResponse(message))
	}
	return messageResponses
}
