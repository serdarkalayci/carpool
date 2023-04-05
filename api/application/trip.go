// Package application is the package that holds the application logic between database and communication layers
package application

import (
	"github.com/rs/zerolog/log"
	apperr "github.com/serdarkalayci/carpool/api/application/errors"
	"github.com/serdarkalayci/carpool/api/domain"
)

// TripRepository is the interface to let outer layers to interact to the Trip Repository
type TripRepository interface {
	AddTrip(trip *domain.Trip) error
	GetTrips(countryID string, origin, destination string) ([]*domain.Trip, error)
	GetTripByID(tripID string) (*domain.TripDetail, error)
	GetConversationID(tripID string, userID string) (string, error)
	GetTripCapacity(tripID string) (int, error)
	SetTripCapacity(tripID string, capacity int) error
}

// TripService is the struct to let outer layers to interact to the Trip Application
type TripService struct {
	dc DataContext
}

// NewTripService creates a new TripService instance and sets its repository
func NewTripService(dc DataContext) TripService {
	return TripService{
		dc: dc,
	}
}

// AddTrip adds a new trip to the repository
func (ts TripService) AddTrip(trip domain.Trip) error {
	// Check if the destination is valid
	if ts.dc.GeographyRepository == nil {
		log.Fatal().Msg("geographyRepository is not set")
	}
	correct, err := ts.dc.GeographyRepository.CheckBallotCity(trip.CountryID, trip.Destination)
	if err != nil {
		return err
	}
	if !correct {
		log.Logger.Info().Msgf("destination is not a ballot city countryID: %s and cityName: %s", trip.CountryID, trip.Destination)
		return apperr.ErrInvalidDestination{}
	}
	return ts.dc.TripRepository.AddTrip(&trip)
}

// GetTrips gets the trips from the repository based on the countryID, origin and destination
func (ts TripService) GetTrips(countryID string, origin, destination string) ([]*domain.Trip, error) {
	return ts.dc.TripRepository.GetTrips(countryID, origin, destination)
}

// GetTrip gets the trip from the repository based on the tripID and userID
func (ts TripService) GetTrip(tripID string, userID string) (*domain.TripDetail, error) {
	if ts.dc.ConversationRepository == nil {
		log.Fatal().Msg("conversationRepository is not set")
	}
	tripDetail, err := ts.dc.TripRepository.GetTripByID(tripID)
	if err != nil {
		return nil, err
	}
	// If the user is the requester, we need to get the conversation between them and the supplier
	if tripDetail.SupplierID != userID {
		conversationID, err := ts.dc.TripRepository.GetConversationID(tripID, userID)
		if err != nil {
			return nil, err
		}
		if conversationID != "" {
			cs := NewConversationService(ts.dc)
			conversation, err := cs.GetConversation(conversationID, userID)
			if err != nil {
				return nil, err
			}
			if conversation != nil {
				tripDetail.Conversations = append(tripDetail.Conversations, *conversation)
				// Because requester is reading the conversation, we need to mark the messages directed to them as read
				err = ts.dc.ConversationRepository.MarkConversationsRead(conversation.ConversationID, "out")
				if err != nil {
					return nil, err
				}
			}
		}
	} else {
		// If the user is the supplier, we need to get the conversations between them and the requesters
		cs := NewConversationService(ts.dc)
		conversations, err := cs.GetConversations(tripID)
		if err != nil {
			return nil, err
		}
		tripDetail.Conversations = conversations
	}
	return tripDetail, nil
}

// SetTripCapacity sets the trip capacity to the repository
func (ts TripService) SetTripCapacity(tripID string, demand int) error {
	currentCap, err := ts.dc.TripRepository.GetTripCapacity(tripID)
	if err != nil {
		return err
	}
	if currentCap+demand < 0 {
		log.Logger.Info().Msgf("demand %d is more than the capacity %d for tripID: %s", demand, currentCap, tripID)
		return apperr.ErrInvalidCapacity{}
	}
	return ts.dc.TripRepository.SetTripCapacity(tripID, demand)
}
