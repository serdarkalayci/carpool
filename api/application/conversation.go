// Package application is the package that holds the application logic between database and communication layers
package application

import (
	"fmt"
	"time"

	"github.com/rs/zerolog/log"
	apperr "github.com/serdarkalayci/carpool/api/application/errors"
	"github.com/serdarkalayci/carpool/api/domain"
	"github.com/spf13/viper"
)

// ConversationRepository is the interface that wraps the basic conversation repository methods
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

// ConversationService is the struct to let outer layers to interact to the Conversation Application
type ConversationService struct {
	dc DataContextCarrier
}

// NewConversationService creates a new ConversationService instance and sets its repository
func NewConversationService(dc DataContextCarrier) ConversationService {
	return ConversationService{
		dc: dc,
	}
}

// InitiateConversation initiates a conversation between a supplier and a requester
func (cs ConversationService) InitiateConversation(tripID string, requesterID string, capacity int, message string) error {
	// First let's get Supplier details from trip. For this we need to set up the trip service
	ts := NewTripService(cs.dc)
	trip, err := ts.GetTripByID(tripID)
	if err != nil {
		return err
	}
	// Then let's get Requester details from user
	us := NewUserService(cs.dc)
	requester, err := us.GetUser(requesterID)
	if err != nil {
		return err
	}
	// Also let's get the supplier details from user because we need their contact details
	supplier, err := us.GetUser(trip.SupplierID)
	if err != nil {
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
	err = cs.dc.GetConversationRepository().InitiateConversation(conversation)
	if err != nil {
		return err
	}
	return nil
}

// AddMessage adds a message to a conversation
func (cs ConversationService) AddMessage(conversationID string, userID string, message string) error {
	// First check that this user is the supplier or the requester of this trip, so that we can decide the direction of the message
	owner, err := cs.dc.GetConversationRepository().CheckConversationOwnership(conversationID, userID)
	if err != nil {
		return err
	}
	direction := "in"
	if owner {
		direction = "out"
	}
	return cs.dc.GetConversationRepository().AddMessage(conversationID, message, direction)
}

// GetConversation gets a conversation by its ID
func (cs ConversationService) GetConversation(conversationID string, userID string) (*domain.Conversation, error) {
	conversation, err := cs.dc.GetConversationRepository().GetConversationByID(conversationID)
	if err != nil {
		return nil, err
	}
	// Check that this user is the supplier or the requester of this conversation so that we can decide if they can see this conversation or not
	if conversation.RequesterID != userID && conversation.SupplierID != userID {
		log.Logger.Error().Err(err).Msgf("user %s is neither the requester or the supplier of this conversation %s", userID, conversationID)
		return nil, apperr.ErrNotAuthorizedForConversation{}
	}

	// Lets decide which messages to mark as read by looking who gets the details of the conversation
	direction := "in"
	if userID == conversation.RequesterID {
		direction = "out"
	}
	err = cs.dc.GetConversationRepository().MarkConversationsRead(conversation.ConversationID, direction)
	if err != nil {
		return nil, err
	}
	// We'll hide the contact details until both sides approve the trip
	if !conversation.SupplierApproved || !conversation.RequesterApproved {
		conversation.RequesterContact = domain.ContactDetails{}
		conversation.SupplierContact = domain.ContactDetails{}
	}
	return conversation, nil
}

// GetConversations gets all conversations for a trip
func (cs ConversationService) GetConversations(tripID string) ([]domain.Conversation, error) {
	conversations, err := cs.dc.GetConversationRepository().GetConversations(tripID)
	if err != nil {
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

// UpdateApproval updates the approval status of a conversation
func (cs ConversationService) UpdateApproval(conversationID string, userID string, approved bool) error {
	if cs.dc.GetTripRepository() == nil {
		return apperr.ErrMissingRepository{}
	}
	// First check that this user is the supplier of this trip
	supplier, err := cs.dc.GetConversationRepository().CheckConversationOwnership(conversationID, userID)
	if err != nil {
		return err
	}
	conversation, err := cs.dc.GetConversationRepository().GetConversationByID(conversationID)
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
	err = cs.dc.GetConversationRepository().UpdateApproval(conversationID, supplierapproved, requesterapproved)
	if err != nil {
		return err
	}

	// Send mail to the other party
	sendEmail(to, viper.GetString("ApprovalSubject"), message)
	return nil
}
