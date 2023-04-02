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
	UpdateApproval(conversationID string, supplierApprove string, requesterApprove string) error
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

func (cs ConversationService) InitiateConversation(tripID string, requesterID string, capacity int, message string) error {
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
	// Also let's get the supplier details from user because we need their contact details
	supplier, err := cs.userRepository.GetUser(trip.SupplierID)
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
		TripID:            tripID,
		RequesterID:       requesterID,
		RequesterName:     requester.Name,
		SupplierID:        trip.SupplierID,
		SupplierName:      trip.SupplierName,
		RequestedCapacity: capacity,
		Messages: []domain.Message{
			{
				Direction: direction,
				Text:      message,
				Date:      time.Now(),
			},
		},
		RequesterContact: domain.ContactDetails{
			Email: requester.Email,
			Phone: requester.Phone,
		},
		SupplierContact: domain.ContactDetails{
			Email: supplier.Email,
			Phone: supplier.Phone,
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

func (cs ConversationService) GetConversation(conversationID string, userID string) (*domain.Conversation, error) {
	conversation, err := cs.conversationRepository.GetConversationByID(conversationID)
	if err != nil {
		log.Logger.Error().Err(err).Msg("error getting conversation")
		return nil, err
	}
	// Lets decide which messages to mark as read by looking who gets the details of the conversation
	direction := "in"
	if userID == conversation.RequesterID {
		direction = "out"
	}
	err = cs.conversationRepository.MarkConversationsRead(conversation.ConversationID, direction)
	if err != nil {
		log.Logger.Error().Err(err).Msg("error marking corresponding messages as read")
		return nil, err
	}
	// We'll hide the contact details until both sides approve the trip
	if !conversation.SupplierApproved || !conversation.RequesterApproved {
		conversation.RequesterContact = domain.ContactDetails{}
		conversation.SupplierContact = domain.ContactDetails{}
	}
	return conversation, nil
}

func (cs ConversationService) UpdateApproval(conversationID string, userID string, approved bool) error {
	if cs.tripRepository == nil {
		log.Error().Msg("tripRepository is not set")
		panic("tripRepository is not set")
	}
	// First check that this user is the supplier of this trip
	supplier, err := cs.conversationRepository.CheckConversationOwnership(conversationID, userID)
	if err != nil {
		return err
	}
	conversation, err := cs.conversationRepository.GetConversationByID(conversationID)
	if err != nil {
		return err
	}
	to := ""
	message := ""
	var supplierapproved, requesterapproved string = "", ""
	bothapproved := false
	switch {
	case approved && supplier:
		supplierapproved = "true"
		// Send mail to requester that the supplier has approved the trip
		to = conversation.RequesterContact.Email
		message = fmt.Sprintf(viper.GetViper().GetString("ApprovalMessagePositive"), conversation.SupplierName)
		// If the requester has already approved the trip, then we'll mark the trip as both approved
		if conversation.RequesterApproved == true {
			bothapproved = true
		}
	case approved && !supplier:
		requesterapproved = "true"
		// Send mail to supplier that the requester has approved the trip
		to = conversation.SupplierContact.Email
		message = fmt.Sprintf(viper.GetViper().GetString("ApprovalMessagePositive"), conversation.RequesterName)
		// If the supplier has already approved the trip, then we'll mark the trip as both approved
		if conversation.SupplierApproved == true {
			bothapproved = true
		}
	case !approved && supplier:
		supplierapproved = "false"
		requesterapproved = "false"
		// Send mail to requester that the supplier has rejected the trip
		to = conversation.RequesterContact.Email
		message = fmt.Sprintf(viper.GetViper().GetString("ApprovalMessageNegative"), conversation.SupplierName)
	case !approved && !supplier:
		supplierapproved = "false"
		requesterapproved = "false"
		// Send mail to supplier that the requester has rejected the trip
		to = conversation.SupplierContact.Email
		message = fmt.Sprintf(viper.GetViper().GetString("ApprovalMessageNegative"), conversation.RequesterName)
	}
	// If both parties have approved the trip, we have to drop the capacity of the trip
	// If the trip capacity is not enough, we have to return an error and not update the conversation
	if bothapproved {
		ts := NewTripService(cs.tripRepository, nil, nil)
		err = ts.SetTripCapacity(conversation.TripID, conversation.RequestedCapacity)
		if err != nil {
			return err
		}
	}
	// Now we can update the conversation
	err = cs.conversationRepository.UpdateApproval(conversationID, supplierapproved, requesterapproved)
	if err != nil {
		return err
	}

	// Send mail to the other party
	sendEmail(to, viper.GetString("ApprovalSubject"), message)
	return nil
}
