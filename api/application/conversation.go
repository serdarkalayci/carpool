package application

import (
	"fmt"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/serdarkalayci/carpool/api/domain"
	"github.com/spf13/viper"
)

type ConversationRepository interface {
	CheckConversationOwnership(conversationID string, userID string) (bool, error)
	InitiateConversation(domain.Conversation) error
	AddMessage(conversationID string, message string, direction string) error
	GetConversation(tripID string, userID string) (*domain.Conversation, error)
	GetConversationByID(conversationID string) (*domain.Conversation, error)
	GetConversations(tripID string) ([]domain.Conversation, error)
	MarkConversationsRead(conversationID string, direction string) error
}

// ConversationService is the struct to let outer layers to interact to the Conversation Applicaton
type ConversationService struct {
	conversationRepository ConversationRepository
	tripRepository         TripRepository
	userRepository         UserRepository
}

// NewConversationService creates a new ConversationService instance and sets its repository
func NewConversationService(cr ConversationRepository, tr TripRepository, ur UserRepository) ConversationService {
	if cr == nil {
		panic("missing conversationRepository")
	}
	cs := ConversationService{
		conversationRepository: cr,
	}
	if tr != nil {
		cs.tripRepository = tr
	}
	if ur != nil {
		cs.userRepository = ur
	}
	return cs
}

func (cs ConversationService) InitiateConversation(tripID string, requesterID, message string) error {
	// We need TripRepository to get the supplier's name
	if cs.tripRepository == nil {
		log.Error().Msg("tripRepository is not set")
		panic("tripRepository is not set")
	}
	// We need UserRepository to get the requester's name
	if cs.userRepository == nil {
		log.Error().Msg("userRepository is not set")
		panic("userRepository is not set")
	}
	// First let's get Supplier details from trip
	trip, err := cs.tripRepository.GetTripByID(tripID)
	if err != nil {
		log.Logger.Error().Err(err).Msgf("error getting trip with tripID: %s", tripID)
		return err
	}
	// Then let's get Requester details from user
	requester, err := cs.userRepository.GetUser(requesterID)
	if err != nil {
		log.Logger.Error().Err(err).Msgf("error getting user with userID: %s", requesterID)
		return err
	}
	// Now we can initiate the conversation
	// if message is empty, then this must be an invitation from the supplier, so we'll autogenerate the message and its direction will be "out"
	// else it must be the requester's first message, so we'll use the message and its direction will be "in"
	direction := "in"
	if message == "" {
		message = fmt.Sprintf(viper.GetViper().GetString("InvitationMessage"), trip.SupplierName)
		direction = "out"
	}
	conversation := domain.Conversation{
		TripID:        tripID,
		RequesterID:   requesterID,
		RequesterName: requester.Name,
		SupplierID:    trip.SupplierID,
		SupplierName:  trip.SupplierName,
		Messages: []domain.Message{
			{
				Direction: direction,
				Text:      message,
				Date:      time.Now(),
			},
		},
	}
	err = cs.conversationRepository.InitiateConversation(conversation)
	if err != nil {
		log.Logger.Error().Err(err).Msgf("error initiating conversation for tripID: %s", tripID)
		return err
	}
	return nil
}

func (cs ConversationService) AddMessage(conversationID string, userID string, message string) error {
	// First check that this user is the supplier or the requester of this trip, so that we can decide the direction of the message
	owner, err := cs.conversationRepository.CheckConversationOwnership(conversationID, userID)
	if err != nil {
		log.Logger.Error().Err(err).Msgf("error checking ownership of conversationID: %s", conversationID)
		return err
	}
	direction := "in"
	if owner {
		direction = "out"
	}
	return cs.conversationRepository.AddMessage(conversationID, message, direction)
}

func (cs ConversationService) GetConversation(tripID string, conversationID string, userID string) (*domain.Conversation, error) {
	// First check that this user is the supplier of this trip
	owner, err := cs.conversationRepository.CheckConversationOwnership(tripID, userID)
	if err != nil {
		return nil, err
	}
	if !owner {
		return nil, domain.ErrNotTheOwner{}
	}
	conversation, err := cs.conversationRepository.GetConversationByID(conversationID)
	if err != nil {
		log.Logger.Error().Err(err).Msg("error getting conversation")
	}
	// Because supplier is reading the conversation, we need to mark the messages directed to them as read
	err = cs.conversationRepository.MarkConversationsRead(conversation.ConversationID, "in")
	if err != nil {
		log.Logger.Error().Err(err).Msg("error marking inbound messages as read")
		return nil, err
	}
	return conversation, nil
}
