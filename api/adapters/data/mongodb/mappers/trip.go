package mappers

import (
	"github.com/rs/zerolog/log"
	"github.com/serdarkalayci/carpool/api/adapters/data/mongodb/dao"
	"github.com/serdarkalayci/carpool/api/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func MapTripDAOs2Trips(tripsDAO []*dao.TripDAO) []*domain.Trip {
	var trips []*domain.Trip
	for _, tripDAO := range tripsDAO {
		trips = append(trips, MapTripDAO2Trip(tripDAO))
	}
	return trips
}

func MapTripDAO2Trip(tripDAO *dao.TripDAO) *domain.Trip {
	return &domain.Trip{
		ID:             tripDAO.ID.Hex(),
		SupplierID:     tripDAO.SupplierID.Hex(),
		CountryID:      tripDAO.CountryID.Hex(),
		Origin:         tripDAO.Origin,
		Stops:          tripDAO.Stops,
		Destination:    tripDAO.Destination,
		TripDate:       tripDAO.TripDate.Time(),
		AvailableSeats: tripDAO.AvailableSeats,
		Note:           tripDAO.Note,
	}
}

func MapTrip2TripDAO(trip domain.Trip) (dao.TripDAO, error) {
	supplierID, err := primitive.ObjectIDFromHex(trip.SupplierID)
	if err != nil {
		log.Error().Err(err).Msgf("Cannot parse ObjectID of SupplierID: %s", trip.SupplierID)
		return dao.TripDAO{}, err
	}
	countryID, err := primitive.ObjectIDFromHex(trip.CountryID)
	if err != nil {
		log.Error().Err(err).Msgf("Cannot parse ObjectID of CountryID: %s", trip.CountryID)
		return dao.TripDAO{}, err
	}
	return dao.TripDAO{
		ID:             primitive.NewObjectID(),
		SupplierID:     supplierID,
		CountryID:      countryID,
		Origin:         trip.Origin,
		Destination:    trip.Destination,
		Stops:          trip.Stops,
		TripDate:       primitive.NewDateTimeFromTime(trip.TripDate),
		AvailableSeats: trip.AvailableSeats,
		Note:           trip.Note,
	}, nil
}

func MapTripDetailDAO2TripDetail(tripDetailDAO *dao.TripDetailDAO) *domain.TripDetail {
	return &domain.TripDetail{
		ID:             tripDetailDAO.ID.Hex(),
		SupplierID:     tripDetailDAO.SupplierID.Hex(),
		SupplierName:   tripDetailDAO.SupplierName,
		Country:        tripDetailDAO.Country,
		Origin:         tripDetailDAO.Origin,
		Stops:          tripDetailDAO.Stops,
		Destination:    tripDetailDAO.Destination,
		TripDate:       tripDetailDAO.TripDate.Time(),
		AvailableSeats: tripDetailDAO.AvailableSeats,
		Note:           tripDetailDAO.Note,
	}
}

func MapConversationDAO2Conversation(conversationDAO *dao.ConversationDAO) *domain.Conversation {
	messages := MapMessageDAOs2Messages(conversationDAO.Messages)
	return &domain.Conversation{
		RequesterID:   conversationDAO.RequesterID.Hex(),
		RequesterName: conversationDAO.RequesterName,
		Messages:      messages,
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
