package application

import (
	"github.com/rs/zerolog/log"
	"github.com/serdarkalayci/carpool/api/domain"
)

type TripRepository interface {
	AddTrip(trip *domain.Trip) error
	GetTrips(countryID string, origin, destination string) ([]*domain.Trip, error)
	GetTripByID(tripID string) (*domain.TripDetail, error)
	CheckConversation(tripID string, userID string) (string, error)
	CheckTripOwnership(tripID string, userID string) (bool, error)
	InitiateConversation(tripID string, userID string, userName string, message string) error
	AddMessage(conversationID string, message string, direction string) error
	GetConversation(tripID string, userID string) (*domain.Conversation, error)
	GetConversationByID(conversationID string) (*domain.Conversation, error)
	GetConversations(tripID string) ([]domain.Conversation, error)
	MarkConversationsRead(conversationID string, direction string) error
}

// TripService is the struct to let outer layers to interact to the Trip Applicatopn
type TripService struct {
	tripRepository TripRepository
}

// NewTripService creates a new TripService instance and sets its repository
func NewTripService(tr TripRepository) TripService {
	if tr == nil {
		panic("missing tripRepository")
	}
	return TripService{
		tripRepository: tr,
	}
}

func (ts TripService) AddTrip(trip domain.Trip) error {
	return ts.tripRepository.AddTrip(&trip)
}

func (ts TripService) GetTrips(countryID string, origin, destination string) ([]*domain.Trip, error) {
	return ts.tripRepository.GetTrips(countryID, origin, destination)
}

func (ts TripService) GetTrip(tripID string, userID string) (*domain.TripDetail, error) {
	tripDetail, err := ts.tripRepository.GetTripByID(tripID)
	if err != nil {
		log.Logger.Error().Err(err).Msg("error getting trip detail")
		return nil, err
	}
	// If the user is the requester, we need to get the conversation between them and the supplier
	if tripDetail.SupplierID != userID {
		conversation, err := ts.tripRepository.GetConversation(tripID, userID)
		if err != nil {
			log.Logger.Error().Err(err).Msg("error getting conversation")
			return nil, err
		}
		if conversation != nil {
			tripDetail.Conversations = append(tripDetail.Conversations, *conversation)
			// Because requester is reading the conversation, we need to mark the messages directed to them as read
			err = ts.tripRepository.MarkConversationsRead(conversation.ConversationID, "out")
			if err != nil {
				log.Logger.Error().Err(err).Msg("error marking outbound messages as read")
				return nil, err
			}
		}
	} else {
		// If the user is the supplier, we need to get the conversations between them and the requesters
		conversations, err := ts.tripRepository.GetConversations(tripID)
		if err != nil {
			return nil, err
		}
		tripDetail.Conversations = conversations
	}
	return tripDetail, nil
}

func (ts TripService) AddRequesterMessage(tripID string, userID string, userName string, message string) error {
	// First check that this user is not the supplier of this trip
	owner, err := ts.tripRepository.CheckTripOwnership(tripID, userID)
	if err != nil {
		return err
	}
	if owner {
		return domain.ErrNotTheOwner{}
	}
	// check if there is a conversation for this trip with this user
	conversationID, err := ts.tripRepository.CheckConversation(tripID, userID)
	if err != nil {
		return err
	}
	// if there is no conversation, create one
	if conversationID == "" {
		// first we have to get the user's name, so that we can write it to the conversation
		return ts.tripRepository.InitiateConversation(tripID, userID, userName, message)
	}
	// if there is a conversation, add the message to it
	direction := "in"
	return ts.tripRepository.AddMessage(conversationID, message, direction)
}

func (ts TripService) AddSupplierMessage(tripID string, userID string, conversationID, message string) error {
	// First check that this user is the supplier of this trip
	owner, err := ts.tripRepository.CheckTripOwnership(tripID, userID)
	if err != nil {
		return err
	}
	if !owner {
		return domain.ErrNotTheOwner{}
	}
	direction := "out"
	return ts.tripRepository.AddMessage(conversationID, message, direction)
}

func (ts TripService) GetConversation(tripID string, conversationID string, userID string) (*domain.Conversation, error) {
	// First check that this user is the supplier of this trip
	owner, err := ts.tripRepository.CheckTripOwnership(tripID, userID)
	if err != nil {
		return nil, err
	}
	if !owner {
		return nil, domain.ErrNotTheOwner{}
	}
	conversation, err := ts.tripRepository.GetConversationByID(conversationID)
	if err != nil {
		log.Logger.Error().Err(err).Msg("error getting conversation")
	}
	// Because supplier is reading the conversation, we need to mark the messages directed to them as read
	err = ts.tripRepository.MarkConversationsRead(conversation.ConversationID, "in")
	if err != nil {
		log.Logger.Error().Err(err).Msg("error marking inbound messages as read")
		return nil, err
	}
	return conversation, nil
}
