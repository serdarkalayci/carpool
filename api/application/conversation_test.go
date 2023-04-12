package application

import (
	"errors"
	"testing"

	apperr "github.com/serdarkalayci/carpool/api/application/errors"
	"github.com/serdarkalayci/carpool/api/domain"
	"github.com/stretchr/testify/assert"
)

type mockConversationRepository struct{}

var (
	checkConversationOwnershipFunc func(conversationID string, userID string) (bool, error)
	initiateConversationFunc       func(domain.Conversation) error
	addMessageFunc                 func(conversationID string, message string, direction string) error
	getConversationFunc            func(tripID string, userID string) (*domain.Conversation, error)
	getConversationByIDFunc        func(conversationID string) (*domain.Conversation, error)
	getConversationsFunc           func(tripID string) ([]domain.Conversation, error)
	markConversationsReadFunc      func(conversationID string, direction string) error
	updateApprovalFunc             func(conversationID string, supplierApprove string, requesterApprove string) error
)

// CheckConversationOwnership checks if a conversation's supplier is the current user
func (m mockConversationRepository) CheckConversationOwnership(conversationID string, userID string) (bool, error) {
	return checkConversationOwnershipFunc(conversationID, userID)
}

// InitiateConversation initiates a conversation
func (m mockConversationRepository) InitiateConversation(conversation domain.Conversation) error {
	return initiateConversationFunc(conversation)
}

// AddMessage adds a message to a conversation
func (m mockConversationRepository) AddMessage(conversationID string, message string, direction string) error {
	return addMessageFunc(conversationID, message, direction)
}

// GetConversation gets a conversation based on the tripID and userID
func (m mockConversationRepository) GetConversation(tripID string, userID string) (*domain.Conversation, error) {
	return getConversationFunc(tripID, userID)
}

// GetConversationByID gets a conversation based on the conversationID
func (m mockConversationRepository) GetConversationByID(conversationID string) (*domain.Conversation, error) {
	return getConversationByIDFunc(conversationID)
}

// GetConversations gets all conversations based on the tripID
func (m mockConversationRepository) GetConversations(tripID string) ([]domain.Conversation, error) {
	return getConversationsFunc(tripID)
}

// MarkConversationsRead marks a conversation as read
func (m mockConversationRepository) MarkConversationsRead(conversationID string, direction string) error {
	return markConversationsReadFunc(conversationID, direction)
}

// UpdateApproval updates the approval status of a conversation
func (m mockConversationRepository) UpdateApproval(conversationID string, supplierApprove string, requesterApprove string) error {
	return updateApprovalFunc(conversationID, supplierApprove, requesterApprove)
}

func TestInitiateConversationForRequest(t *testing.T) {
	mc := &MockContext{}
	mc.SetRepositories(&mockUserRepository{}, nil, nil, &mockTripRepository{}, &mockConversationRepository{}, &mockRequestRepository{})
	cs := NewConversationService(mc)
	// Check the case when GetTripCapacity returns an error
	getTripByIDFunc = func(tripID string) (*domain.TripDetail, error) {
		return nil, apperr.ErrTripNotFound{}
	}
	err := cs.InitiateConversationForRequest("trip1", "request1")
	assert.ErrorAs(t, err, &apperr.ErrTripNotFound{})
	// Check the case when GetTripCapacity returns a trip detail and GetRequest returns an error
	getTripByIDFunc = func(tripID string) (*domain.TripDetail, error) {
		return &domain.TripDetail{
			ID:             "trip1",
			SupplierID:     "supplier1",
			AvailableSeats: 3,
		}, nil
	}
	getRequestFunc = func(requestID string) (*domain.Request, error) {
		return nil, apperr.ErrRequestNotFound{}
	}
	err = cs.InitiateConversationForRequest("trip1", "request1")
	assert.ErrorAs(t, err, &apperr.ErrRequestNotFound{})
	// Fix the GetRequest function
	getRequestFunc = func(requestID string) (*domain.Request, error) {
		return &domain.Request{
			ID:             "request1",
			RequesterID:    "requester1",
			RequesterName:  "requester1",
			RequestedSeats: 2,
		}, nil
	}
	// now, the application will try to get requster by its ID and the supplier by its ID from the trip
	// because it's the same method which returns the users, we have to set it accordingly
	// first set it to return and error for the requester
	getUserFunc = func(userID string) (domain.User, error) {
		if userID == "requester1" {
			return domain.User{}, apperr.ErrUserNotFound{}
		}
		return domain.User{
			ID:   "supplier1",
			Name: "supplier1",
		}, nil
	}
	err = cs.InitiateConversationForRequest("trip1", "request1")
	assert.ErrorAs(t, err, &apperr.ErrUserNotFound{})
	// now, set it to return an error for the supplier
	getUserFunc = func(userID string) (domain.User, error) {
		if userID == "requester1" {
			return domain.User{
				ID:   "requester1",
				Name: "requester1",
			}, nil

		}
		return domain.User{}, apperr.ErrUserNotFound{}
	}
	err = cs.InitiateConversationForRequest("trip1", "request1")
	assert.ErrorAs(t, err, &apperr.ErrUserNotFound{})
	// now, set it to return a user for both the requester and the supplier
	// now, set it to return an error for the supplier
	getUserFunc = func(userID string) (domain.User, error) {
		if userID == "requester1" {
			return domain.User{
				ID:   "requester1",
				Name: "requester1",
			}, nil

		}
		return domain.User{
			ID:   "supplier1",
			Name: "supplier1",
		}, nil
	}
	// now, the application will try to initiate a conversation
	initiateConversationFunc = func(conversation domain.Conversation) error {
		return errors.New("some database error")
	}
	err = cs.InitiateConversationForRequest("trip1", "request1")
	assert.EqualError(t, err, "some database error")
	// now test the case where 1nitiateConversation does not return an error
	initiateConversationFunc = func(conversation domain.Conversation) error {
		return nil
	}
	err = cs.InitiateConversationForRequest("trip1", "request1")
	assert.NoError(t, err)
}
func TestInitiateConversation(t *testing.T) {
	mc := &MockContext{}
	mc.SetRepositories(&mockUserRepository{}, nil, nil, &mockTripRepository{}, &mockConversationRepository{}, nil)
	cs := NewConversationService(mc)
	// Check the case when GetTripCapacity returns an error
	getTripByIDFunc = func(tripID string) (*domain.TripDetail, error) {
		return nil, apperr.ErrTripNotFound{}
	}
	err := cs.InitiateConversation("trip1", "requester1", 3, "message1")
	assert.ErrorAs(t, err, &apperr.ErrTripNotFound{})
	// Check the case when GetTripCapacity returns a trip detail
	getTripByIDFunc = func(tripID string) (*domain.TripDetail, error) {
		return &domain.TripDetail{
			ID:             "trip1",
			SupplierID:     "supplier1",
			AvailableSeats: 3,
		}, nil
	}
	// now, the application will try to get requster by its ID and the supplier by its ID from the trip
	// because it's the same method which returns the users, we have to set it accordingly
	// first set it to return and error for the requester
	getUserFunc = func(userID string) (domain.User, error) {
		if userID == "requester1" {
			return domain.User{}, apperr.ErrUserNotFound{}
		}
		return domain.User{
			ID:   "supplier1",
			Name: "supplier1",
		}, nil
	}
	err = cs.InitiateConversation("trip1", "requester1", 3, "message1")
	assert.ErrorAs(t, err, &apperr.ErrUserNotFound{})
	// now, set it to return an error for the supplier
	getUserFunc = func(userID string) (domain.User, error) {
		if userID == "requester1" {
			return domain.User{
				ID:   "requester1",
				Name: "requester1",
			}, nil

		}
		return domain.User{}, apperr.ErrUserNotFound{}
	}
	err = cs.InitiateConversation("trip1", "requester1", 3, "message1")
	assert.ErrorAs(t, err, &apperr.ErrUserNotFound{})
	// now, set it to return a user for both the requester and the supplier
	// now, set it to return an error for the supplier
	getUserFunc = func(userID string) (domain.User, error) {
		if userID == "requester1" {
			return domain.User{
				ID:   "requester1",
				Name: "requester1",
			}, nil

		}
		return domain.User{
			ID:   "supplier1",
			Name: "supplier1",
		}, nil
	}
	// now, the application will try to initiate a conversation
	initiateConversationFunc = func(conversation domain.Conversation) error {
		return errors.New("some database error")
	}
	err = cs.InitiateConversation("trip1", "requester1", 3, "")
	assert.EqualError(t, err, "some database error")
	// now test the case where 1nitiateConversation does not return an error
	initiateConversationFunc = func(conversation domain.Conversation) error {
		return nil
	}
	err = cs.InitiateConversation("trip1", "requester1", 3, "")
	assert.NoError(t, err)
}

func TestAddMessage(t *testing.T) {
	mc := &MockContext{}
	mc.SetRepositories(nil, nil, nil, nil, &mockConversationRepository{}, nil)
	cs := NewConversationService(mc)
	// Check the case when CheckConversationOwnership returns an error
	checkConversationOwnershipFunc = func(conversationID string, userID string) (bool, error) {
		return false, apperr.ErrConversationNotFound{}
	}
	err := cs.AddMessage("conversation1", "requester1", "message1")
	assert.ErrorAs(t, err, &apperr.ErrConversationNotFound{})
	// Check the case when CheckConversationOwnership returns false
	checkConversationOwnershipFunc = func(conversationID string, userID string) (bool, error) {
		return true, nil
	}
	// now, the application will try to add a message
	addMessageFunc = func(conversationID string, message string, direction string) error {
		return errors.New("some database error")
	}
	err = cs.AddMessage("conversation1", "requester1", "message1")
	assert.EqualError(t, err, "some database error")
	// now test the case where addMessage does not return an error
	addMessageFunc = func(conversationID string, message string, direction string) error {
		return nil
	}
	err = cs.AddMessage("conversation1", "requester1", "message1")
	assert.NoError(t, err)
}

func TestUpdateApproval(t *testing.T) {
	mc := &MockContext{}
	mc.SetRepositories(nil, nil, nil, nil, &mockConversationRepository{}, nil)
	cs := NewConversationService(mc)
	// This method also needs TripRepository, but it's not supplied.
	err := cs.UpdateApproval("conversation1", "requester1", true)
	assert.ErrorAs(t, err, &apperr.ErrMissingRepository{})
	// Fix the missing repository
	mc.SetRepositories(nil, nil, nil, &mockTripRepository{}, &mockConversationRepository{}, nil)
	// Check the case when CheckConversationOwnership returns an error
	checkConversationOwnershipFunc = func(conversationID string, userID string) (bool, error) {
		return false, apperr.ErrConversationNotFound{}
	}
	err = cs.UpdateApproval("conversation1", "requester1", true)
	assert.ErrorAs(t, err, &apperr.ErrConversationNotFound{})
	// Check the case when CheckConversationOwnership returns true, meaning the user is the supplier
	// but this time GetCoversationByID returns an error
	checkConversationOwnershipFunc = func(conversationID string, userID string) (bool, error) {
		return true, nil
	}
	getConversationByIDFunc = func(conversationID string) (*domain.Conversation, error) {
		return nil, apperr.ErrConversationNotFound{}
	}
	err = cs.UpdateApproval("conversation1", "requester1", true)

}

func TestUpdateApprovalAsSupplier(t *testing.T) {
	mc := &MockContext{}
	mc.SetRepositories(nil, nil, nil, &mockTripRepository{}, &mockConversationRepository{}, nil)
	cs := NewConversationService(mc)

	checkConversationOwnershipFunc = func(conversationID string, userID string) (bool, error) {
		if userID == "supplier1" {
			return true, nil
		}
		return false, nil
	}
	// this setting will set supplierApproved to true, while the requester did not approve yet
	getConversationByIDFunc = func(conversationID string) (*domain.Conversation, error) {
		return &domain.Conversation{
			TripID:            "trip1",
			ConversationID:    "conversation1",
			RequestedCapacity: 3,
			SupplierApproved:  false,
			RequesterApproved: false,
		}, nil
	}
	// because they are not both approved nor rejected, the application will just update the approval status
	updateApprovalFunc = func(conversationID string, supplierApprove string, requesterApprove string) error {
		assert.Equal(t, "true", supplierApprove)
		return errors.New("some database error")
	}
	err := cs.UpdateApproval("conversation1", "supplier1", true)
	assert.EqualError(t, err, "some database error")
	// this setting will set supplierApproved to true, while the requester already approved\
	// because the state is both approved, the application will try to update the capacity of the trip
	getConversationByIDFunc = func(conversationID string) (*domain.Conversation, error) {
		return &domain.Conversation{
			TripID:            "trip1",
			ConversationID:    "conversation1",
			RequestedCapacity: 3,
			SupplierApproved:  false,
			RequesterApproved: true,
		}, nil
	}
	updateApprovalFunc = func(conversationID string, supplierApprove string, requesterApprove string) error {
		assert.Equal(t, "true", supplierApprove)
		return nil
	}
	// first check the case where SetTripCapacity returns an error
	setTripCapacityFunc = func(tripID string, capacity int) error {
		assert.Equal(t, "trip1", tripID)
		assert.Equal(t, 3, capacity)
		return nil
	}
	// fisrt let's check the case where GetTripCapacity returns a capacity that is less than the requested capacity
	getTripCapacityFunc = func(tripID string) (int, error) {
		assert.Equal(t, "trip1", tripID)
		return 2, nil
	}
	err = cs.UpdateApproval("conversation1", "supplier1", true)
	assert.EqualError(t, err, "requested capacity is greater than available capacity")
	// fix the error
	setTripCapacityFunc = func(tripID string, capacity int) error {
		assert.Equal(t, "trip1", tripID)
		assert.Equal(t, -3, capacity)
		return nil
	}
	// fisrt let's check the case where GetTripCapacity returns a capacity that is less than the requested capacity
	getTripCapacityFunc = func(tripID string) (int, error) {
		assert.Equal(t, "trip1", tripID)
		return 5, nil
	}
	err = cs.UpdateApproval("conversation1", "supplier1", true)
	assert.NoError(t, err)
	// let's check the case where the supplier rejects the request
	// first set the current state of the conversation to both approved, so that we can test capacity update
	getConversationByIDFunc = func(conversationID string) (*domain.Conversation, error) {
		return &domain.Conversation{
			TripID:            "trip1",
			ConversationID:    "conversation1",
			RequestedCapacity: 3,
			SupplierApproved:  true,
			RequesterApproved: true,
		}, nil
	}
	// and set SetTripCapacity to return an error
	setTripCapacityFunc = func(tripID string, capacity int) error {
		assert.Equal(t, "trip1", tripID)
		assert.Equal(t, 3, capacity)
		return errors.New("some database error")
	}
	err = cs.UpdateApproval("conversation1", "supplier1", false)
	assert.EqualError(t, err, "some database error")
	// fix the error and check if the calculations are correct
	setTripCapacityFunc = func(tripID string, capacity int) error {
		assert.Equal(t, "trip1", tripID)
		assert.Equal(t, 3, capacity)
		return nil
	}
}

func TestUpdateApprovalAsRequester(t *testing.T) {
	mc := &MockContext{}
	mc.SetRepositories(nil, nil, nil, &mockTripRepository{}, &mockConversationRepository{}, nil)
	cs := NewConversationService(mc)

	checkConversationOwnershipFunc = func(conversationID string, userID string) (bool, error) {
		if userID == "requester1" {
			return false, nil
		}
		return true, nil
	}
	// this setting will set requestApproved to true, while the supplier did not approve yet
	getConversationByIDFunc = func(conversationID string) (*domain.Conversation, error) {
		return &domain.Conversation{
			TripID:            "trip1",
			ConversationID:    "conversation1",
			RequestedCapacity: 3,
			SupplierApproved:  false,
			RequesterApproved: false,
		}, nil
	}
	// because they are not both approved nor rejected, the application will just update the approval status
	updateApprovalFunc = func(conversationID string, supplierApprove string, requesterApprove string) error {
		assert.Equal(t, "true", requesterApprove)
		return errors.New("some database error")
	}
	err := cs.UpdateApproval("conversation1", "requester1", true)
	assert.EqualError(t, err, "some database error")
	// this setting will set requesterApproved to true, while the supplier already approved
	// because the state is both approved, the application will try to update the capacity of the trip
	getConversationByIDFunc = func(conversationID string) (*domain.Conversation, error) {
		return &domain.Conversation{
			TripID:            "trip1",
			ConversationID:    "conversation1",
			RequestedCapacity: 3,
			SupplierApproved:  true,
			RequesterApproved: false,
		}, nil
	}
	updateApprovalFunc = func(conversationID string, supplierApprove string, requesterApprove string) error {
		assert.Equal(t, "true", requesterApprove)
		return nil
	}
	// first check the case where SetTripCapacity returns an error
	setTripCapacityFunc = func(tripID string, capacity int) error {
		assert.Equal(t, "trip1", tripID)
		assert.Equal(t, 3, capacity)
		return nil
	}
	// fisrt let's check the case where GetTripCapacity returns a capacity that is less than the requested capacity
	getTripCapacityFunc = func(tripID string) (int, error) {
		assert.Equal(t, "trip1", tripID)
		return 2, nil
	}
	err = cs.UpdateApproval("conversation1", "requester1", true)
	assert.EqualError(t, err, "requested capacity is greater than available capacity")
	// fix the error
	setTripCapacityFunc = func(tripID string, capacity int) error {
		assert.Equal(t, "trip1", tripID)
		assert.Equal(t, -3, capacity)
		return nil
	}
	// fisrt let's check the case where GetTripCapacity returns a capacity that is less than the requested capacity
	getTripCapacityFunc = func(tripID string) (int, error) {
		assert.Equal(t, "trip1", tripID)
		return 5, nil
	}
	err = cs.UpdateApproval("conversation1", "requester1", true)
	assert.NoError(t, err)
	// let's check the case where the requester rejects the request
	// first set the current state of the conversation to both approved, so that we can test capacity update
	getConversationByIDFunc = func(conversationID string) (*domain.Conversation, error) {
		return &domain.Conversation{
			TripID:            "trip1",
			ConversationID:    "conversation1",
			RequestedCapacity: 3,
			SupplierApproved:  true,
			RequesterApproved: true,
		}, nil
	}
	// and set SetTripCapacity to return an error
	setTripCapacityFunc = func(tripID string, capacity int) error {
		assert.Equal(t, "trip1", tripID)
		assert.Equal(t, 3, capacity)
		return errors.New("some database error")
	}
	err = cs.UpdateApproval("conversation1", "requester1", false)
	assert.EqualError(t, err, "some database error")
	// fix the error and check if the calculations are correct
	setTripCapacityFunc = func(tripID string, capacity int) error {
		assert.Equal(t, "trip1", tripID)
		assert.Equal(t, 3, capacity)
		return nil
	}
}
