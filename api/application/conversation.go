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
	dc DataContext
}

// NewConversationService creates a new ConversationService instance and sets its repository
func NewConversationService(dc DataContext) ConversationService {
	return ConversationService{
		dc: dc,
	}
}

func (cs ConversationService) InitiateConversation(tripID string, requesterID string, capacity int, message string) error {
	// First let's get Supplier details from trip
	trip, err := cs.dc.TripRepository.GetTripByID(tripID)
	if err != nil {
		log.Logger.Error().Err(err).Msgf("error getting trip with tripID: %s", tripID)
		return err
	}
	// Then let's get Requester details from user
	requester, err := cs.dc.UserRepository.GetUser(requesterID)
	if err != nil {
		log.Logger.Error().Err(err).Msgf("error getting user with userID: %s", requesterID)
		return err
	}
	// Also let's get the supplier details from user because we need their contact details
	supplier, err := cs.dc.UserRepository.GetUser(trip.SupplierID)
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
	err = cs.dc.ConversationRepository.InitiateConversation(conversation)
	if err != nil {
		log.Logger.Error().Err(err).Msgf("error initiating conversation for tripID: %s", tripID)
		return err
	}
	return nil
}

func (cs ConversationService) AddMessage(conversationID string, userID string, message string) error {
	// First check that this user is the supplier or the requester of this trip, so that we can decide the direction of the message
	owner, err := cs.dc.ConversationRepository.CheckConversationOwnership(conversationID, userID)
	if err != nil {
		log.Logger.Error().Err(err).Msgf("error checking ownership of conversationID: %s", conversationID)
		return err
	}
	direction := "in"
	if owner {
		direction = "out"
	}
	return cs.dc.ConversationRepository.AddMessage(conversationID, message, direction)
}

func (cs ConversationService) GetConversation(conversationID string, userID string) (*domain.Conversation, error) {
	conversation, err := cs.dc.ConversationRepository.GetConversationByID(conversationID)
	if err != nil {
		log.Logger.Error().Err(err).Msg("error getting conversation")
		return nil, err
	}
	// Check that this user is the supplier or the requester of this conversation so that we can decide if they can see this conversation or not
	if conversation.RequesterID != userID && conversation.SupplierID != userID {
		log.Logger.Error().Err(err).Msg("user is not the owner of this conversation")
		return nil, ErrNotAuthorizedForConversation{}
	}

	// Lets decide which messages to mark as read by looking who gets the details of the conversation
	direction := "in"
	if userID == conversation.RequesterID {
		direction = "out"
	}
	err = cs.dc.ConversationRepository.MarkConversationsRead(conversation.ConversationID, direction)
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

func (cs ConversationService) GetConversations(tripID string) ([]domain.Conversation, error) {
	conversations, err := cs.dc.ConversationRepository.GetConversations(tripID)
	if err != nil {
		log.Logger.Error().Err(err).Msg("error getting conversations")
		return nil, err
	}
	// We'll hide the contact details until both sides approve the trip
	for i := range conversations {
		if !conversations[i].SupplierApproved || !conversations[i].RequesterApproved {
			conversations[i].RequesterContact = domain.ContactDetails{}
			conversations[i].SupplierContact = domain.ContactDetails{}
		}
	}
	return conversations, nil
}

func (cs ConversationService) UpdateApproval(conversationID string, userID string, approved bool) error {
	if cs.dc.TripRepository == nil {
		log.Error().Msg("tripRepository is not set")
		panic("tripRepository is not set")
	}
	// First check that this user is the supplier of this trip
	supplier, err := cs.dc.ConversationRepository.CheckConversationOwnership(conversationID, userID)
	if err != nil {
		return err
	}
	conversation, err := cs.dc.ConversationRepository.GetConversationByID(conversationID)
	if err != nil {
		return err
	}
	to := ""
	message := ""
	var supplierapproved, requesterapproved string = "", ""
	bothapproved := false
	bothrejected := false
	capacity := conversation.RequestedCapacity
	switch {
	case approved && supplier:
		supplierapproved = "true"
		// Send mail to requester that the supplier has approved the trip
		to = conversation.RequesterContact.Email
		message = fmt.Sprintf(viper.GetViper().GetString("ApprovalMessagePositive"), conversation.SupplierName)
		// If the requester has already approved the trip, then we'll mark the trip as both approved
		if conversation.RequesterApproved == true {
			bothapproved = true
			capacity = capacity * -1
		}
	case approved && !supplier:
		requesterapproved = "true"
		// Send mail to supplier that the requester has approved the trip
		to = conversation.SupplierContact.Email
		message = fmt.Sprintf(viper.GetViper().GetString("ApprovalMessagePositive"), conversation.RequesterName)
		// If the supplier has already approved the trip, then we'll mark the trip as both approved
		if conversation.SupplierApproved == true {
			bothapproved = true
			capacity = capacity * -1
		}
	case !approved && supplier:
		supplierapproved = "false"
		// Send mail to requester that the supplier has rejected the trip
		to = conversation.RequesterContact.Email
		message = fmt.Sprintf(viper.GetViper().GetString("ApprovalMessageNegative"), conversation.SupplierName)
		bothrejected = true
	case !approved && !supplier:
		requesterapproved = "false"
		// Send mail to supplier that the requester has rejected the trip
		to = conversation.SupplierContact.Email
		message = fmt.Sprintf(viper.GetViper().GetString("ApprovalMessageNegative"), conversation.RequesterName)
		bothrejected = true
	}
	// If both parties have approved the trip, we have to drop the capacity of the trip
	// If the trip capacity is not enough, we have to return an error and not update the conversation
	if bothapproved {
		ts := NewTripService(cs.dc)
		err = ts.SetTripCapacity(conversation.TripID, capacity)
		if err != nil {
			return err
		}
	}
	// If the current state of the conversation is both approved, and the new state is both rejected, we have to increase the capacity of the trip
	if conversation.RequesterApproved && conversation.SupplierApproved && bothrejected {
		ts := NewTripService(cs.dc)
		err = ts.SetTripCapacity(conversation.TripID, capacity)
		if err != nil {
			return err
		}
	}
	// Now we can update the conversation
	err = cs.dc.ConversationRepository.UpdateApproval(conversationID, supplierapproved, requesterapproved)
	if err != nil {
		return err
	}

	// Send mail to the other party
	sendEmail(to, viper.GetString("ApprovalSubject"), message)
	return nil
}

type ErrNotTheOwner struct{}

func (e ErrNotTheOwner) Error() string {
	return "this user is not the supplier of this trip"
}

type ErrTheOwner struct{}

func (e ErrTheOwner) Error() string {
	return "this user is the supplier of this trip, therefore cannot inititate conversation"
}

type ErrNotAuthorizedForConversation struct{}

func (e ErrNotAuthorizedForConversation) Error() string {
	return "this user is not authorized to see this conversation"
}

type ErrConversationNotFound struct{}

func (e ErrConversationNotFound) Error() string {
	return "conversation not found"
}

type ErrConversationNotInserted struct{}

func (e ErrConversationNotInserted) Error() string {
	return "conversation not inserted"
}

type ErrConversationNotUpdated struct{}

func (e ErrConversationNotUpdated) Error() string {
	return "conversation not updated"
}

type ErrMessageNotInserted struct{}

func (e ErrMessageNotInserted) Error() string {
	return "message not inserted"
}

type ErrMessageNotUpdated struct{}

func (e ErrMessageNotUpdated) Error() string {
	return "message not updated"
}
