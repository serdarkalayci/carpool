package application

import (
	"github.com/rs/zerolog/log"
	"github.com/serdarkalayci/carpool/api/domain"
)

type TripRepository interface {
	AddTrip(trip *domain.Trip) error
	GetTrips(countryID string, origin, destination string) ([]*domain.Trip, error)
	GetTripByID(tripID string) (*domain.TripDetail, error)
}

// TripService is the struct to let outer layers to interact to the Trip Applicaton
type TripService struct {
	tripRepository         TripRepository
	geographyRepository    GeographyRepository
	conversationRepository ConversationRepository
}

// NewTripService creates a new TripService instance and sets its repository
func NewTripService(tr TripRepository, gr GeographyRepository, cr ConversationRepository) TripService {
	if tr == nil {
		panic("missing tripRepository")
	}
	ts := TripService{
		tripRepository: tr,
	}
	if gr != nil {
		ts.geographyRepository = gr
	}
	if cr != nil {
		ts.conversationRepository = cr
	}
	return ts
}

func (ts TripService) AddTrip(trip domain.Trip) error {
	// Check if the destination is valid
	if ts.geographyRepository == nil {
		log.Logger.Error().Msg("geographyRepository is not set")
		panic("missing geographyRepository")
	}
	correct, err := ts.geographyRepository.CheckBallotCity(trip.CountryID, trip.Destination)
	if err != nil {
		log.Logger.Error().Err(err).Msgf("error checking destination city with countryID: %s and cityName: %s", trip.CountryID, trip.Destination)
		return err
	}
	if !correct {
		log.Logger.Info().Msgf("destination is not a ballot city countryID: %s and cityName: %s", trip.CountryID, trip.Destination)
		return domain.ErrInvalidDestination{}
	}
	return ts.tripRepository.AddTrip(&trip)
}

func (ts TripService) GetTrips(countryID string, origin, destination string) ([]*domain.Trip, error) {
	return ts.tripRepository.GetTrips(countryID, origin, destination)
}

func (ts TripService) GetTrip(tripID string, userID string) (*domain.TripDetail, error) {
	if ts.conversationRepository == nil {
		log.Logger.Error().Msg("conversationRepository is not set")
		panic("missing conversationRepository")
	}
	tripDetail, err := ts.tripRepository.GetTripByID(tripID)
	if err != nil {
		log.Logger.Error().Err(err).Msg("error getting trip detail")
		return nil, err
	}
	// If the user is the requester, we need to get the conversation between them and the supplier
	if tripDetail.SupplierID != userID {
		conversation, err := ts.conversationRepository.GetConversation(tripID, userID)
		if err != nil {
			log.Logger.Error().Err(err).Msg("error getting conversation")
			return nil, err
		}
		if conversation != nil {
			tripDetail.Conversations = append(tripDetail.Conversations, *conversation)
			// Because requester is reading the conversation, we need to mark the messages directed to them as read
			err = ts.conversationRepository.MarkConversationsRead(conversation.ConversationID, "out")
			if err != nil {
				log.Logger.Error().Err(err).Msg("error marking outbound messages as read")
				return nil, err
			}
		}
	} else {
		// If the user is the supplier, we need to get the conversations between them and the requesters
		conversations, err := ts.conversationRepository.GetConversations(tripID)
		if err != nil {
			return nil, err
		}
		tripDetail.Conversations = conversations
	}
	return tripDetail, nil
}
