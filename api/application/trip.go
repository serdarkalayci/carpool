package application

import "github.com/serdarkalayci/carpool/api/domain"

type TripRepository interface {
	AddTrip(trip *domain.Trip) error
	GetTrips(countryID string, origin, destination string) ([]*domain.Trip, error)
	GetTripByID(tripID string) (*domain.TripDetail, error)
	CheckConversation(tripID string, userID string) (string, string, error)
	InitiateConversation(tripID string, userID string, userName string, message string) error
	AddMessage(conversationID string, message string, direction string) error
	GetConversation(tripID string, userID string) (*domain.Conversation, error)
	GetConversations(tripID string) ([]domain.Conversation, error)
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
		return nil, err
	}
	// If the user is the requester, we need to get the conversation between them and the supplier
	if tripDetail.SupplierID != userID {
		conversation, err := ts.tripRepository.GetConversation(tripID, userID)
		if err != nil {
			return nil, err
		}
		if conversation != nil {
			tripDetail.Conversations = append(tripDetail.Conversations, *conversation)
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

func (ts TripService) AddMessage(tripID string, userID string, userName string, message string) error {
	// check if there is a conversation for this trip with this user
	conversationID, requesterID, err := ts.tripRepository.CheckConversation(tripID, userID)
	if err != nil {
		return err
	}
	// if there is no conversation, create one
	if conversationID == "" {
		// first we have to get the user's name, so that we can write it to the conversation
		return ts.tripRepository.InitiateConversation(tripID, userID, userName, message)
	}
	// if there is a conversation, add the message to it
	direction := "out"
	// if the user is the requester, the message is incoming
	if requesterID == userID {
		direction = "in"
	}
	return ts.tripRepository.AddMessage(conversationID, message, direction)
}
