package application

import (
	"testing"

	apperr "github.com/serdarkalayci/carpool/api/application/errors"
	"github.com/serdarkalayci/carpool/api/domain"
	"github.com/stretchr/testify/assert"
)

type mockTripRepository struct{}

var (
	addTripFunc           func(trip *domain.Trip) error
	getTripsFunc          func(countryID string, origin, destination string) ([]*domain.Trip, error)
	getTripByIDFunc       func(tripID string) (*domain.TripDetail, error)
	getConversationIDFunc func(tripID string, userID string) (string, error)
	getTripCapacityFunc   func(tripID string) (int, error)
	setTripCapacityFunc   func(tripID string, capacity int) error
)

// AddTrip adds a new trip to the repository
func (m mockTripRepository) AddTrip(trip *domain.Trip) error {
	return addTripFunc(trip)
}

// GetTrips gets the trips from the repository based on the countryID, origin and destination
func (m mockTripRepository) GetTrips(countryID string, origin, destination string) ([]*domain.Trip, error) {
	return getTripsFunc(countryID, origin, destination)
}

// GetTripByID gets the trip from the repository based on the tripID
func (m mockTripRepository) GetTripByID(tripID string) (*domain.TripDetail, error) {
	return getTripByIDFunc(tripID)
}

// GetConversationID gets the conversationID from the repository based on the tripID and userID
func (m mockTripRepository) GetConversationID(tripID string, userID string) (string, error) {
	return getConversationIDFunc(tripID, userID)
}

// GetTripCapacity gets the trip capacity from the repository based on the tripID
func (m mockTripRepository) GetTripCapacity(tripID string) (int, error) {
	return getTripCapacityFunc(tripID)
}

// SetTripCapacity sets the trip capacity to the repository based on the tripID
func (m mockTripRepository) SetTripCapacity(tripID string, capacity int) error {
	return setTripCapacityFunc(tripID, capacity)
}

func TestNewTripService(t *testing.T) {
	mc := &MockContext{}
	mc.SetRepositories(nil, nil, nil, &mockTripRepository{}, nil, nil)
	ts := NewTripService(mc)
	assert.NotNil(t, ts)
	gr := ts.dc.GetTripRepository()
	assert.NotNil(t, gr)
	assert.Implements(t, (*TripRepository)(nil), gr)
}

func TestAddTrip(t *testing.T) {
	mc := &MockContext{}
	mc.SetRepositories(nil, nil, nil, &mockTripRepository{}, nil, nil) // We didn't set the GeographyRepository here, so a call to AddTrip will fail
	ts := NewTripService(mc)
	var trip domain.Trip
	addTripFunc = func(trip *domain.Trip) error {
		return apperr.ErrInvalidDestination{}
	}
	err := ts.AddTrip(trip)
	assert.EqualError(t, err, "an internal problem occured")
	// Set the GeographyRepository
	mc.SetRepositories(nil, nil, &mockGeographyRepository{}, &mockTripRepository{}, nil, nil)
	// First check if CheckBallotCity returns an error
	checkBallotCityFunc = func(countryID string, city string) (bool, error) {
		return false, apperr.ErrInvalidDestination{}
	}
	err = ts.AddTrip(trip)
	assert.ErrorAs(t, err, &apperr.ErrInvalidDestination{})
	// Check the case when CheckBallotCity does not return an error but a false result
	checkBallotCityFunc = func(countryID string, city string) (bool, error) {
		return false, nil
	}
	err = ts.AddTrip(trip)
	assert.ErrorAs(t, err, &apperr.ErrInvalidDestination{})
	// Check the case when CheckBallotCity does not return an error and a true result
	checkBallotCityFunc = func(countryID string, city string) (bool, error) {
		return true, nil
	}
	// Check the case when addTripFunc returns an error
	addTripFunc = func(trip *domain.Trip) error {
		return apperr.ErrInvalidDestination{}
	}
	err = ts.AddTrip(trip)
	assert.ErrorAs(t, err, &apperr.ErrInvalidDestination{})
	// Check the case when addTripFunc does not return an error
	addTripFunc = func(trip *domain.Trip) error {
		return nil
	}
	err = ts.AddTrip(trip)
	assert.Nil(t, err)
}

func TestGetTrips(t *testing.T) {
	mc := &MockContext{}
	mc.SetRepositories(nil, nil, nil, &mockTripRepository{}, nil, nil)
	ts := NewTripService(mc)
	// Check the case when getTripsFunc returns an error
	getTripsFunc = func(countryID string, origin, destination string) ([]*domain.Trip, error) {
		return nil, apperr.ErrInvalidDestination{}
	}
	_, err := ts.GetTrips("TR", "Istanbul", "Ankara")
	assert.ErrorAs(t, err, &apperr.ErrInvalidDestination{})
}

func TestGetTrip(t *testing.T) {
	mc := &MockContext{}
	mc.SetRepositories(nil, nil, nil, &mockTripRepository{}, nil, nil) // We didn't set the ConversationRepository here, so a call to GetTrip`` will fail
	ts := NewTripService(mc)
	// Check the case when the ConversationRepository is not set
	addTripFunc = func(trip *domain.Trip) error {
		return apperr.ErrInvalidDestination{}
	}
	trip, err := ts.GetTrip("trip1", "user1")
	assert.EqualError(t, err, "an internal problem occured")
	assert.Nil(t, trip)
	// Set the ConversationRepository
	mc.SetRepositories(nil, nil, nil, &mockTripRepository{}, &mockConversationRepository{}, nil)
	// Check the case when GetTripByID returns an error
	getTripByIDFunc = func(tripID string) (*domain.TripDetail, error) {
		return nil, apperr.ErrTripNotFound{}
	}
	trip, err = ts.GetTrip("trip1", "user1")
	assert.ErrorAs(t, err, &apperr.ErrTripNotFound{})
	assert.Nil(t, trip)
}

func TestGetTripAsRequester(t *testing.T) {
	mc := &MockContext{}
	mc.SetRepositories(nil, nil, nil, &mockTripRepository{}, &mockConversationRepository{}, nil)
	ts := NewTripService(mc)
	// Check the case when GetTripByID does not return an error, and returns a trip detail where the user is not the supplier
	getTripByIDFunc = func(tripID string) (*domain.TripDetail, error) {
		return &domain.TripDetail{
			ID:         "trip1",
			SupplierID: "user2",
		}, nil
	}
	// in this case, the applicaton will also try to load the conversation of this requester under this trip
	getConversationIDFunc = func(tripID string, userID string) (string, error) {
		return "", apperr.ErrConversationNotFound{}
	}
	trip, err := ts.GetTrip("trip1", "user1")
	assert.ErrorAs(t, err, &apperr.ErrConversationNotFound{})
	assert.Nil(t, trip)
	// check the case when GetConversationID returns an error
	getConversationIDFunc = func(tripID string, userID string) (string, error) {
		return "conversation1", nil
	}
	getConversationByIDFunc = func(conversationID string) (*domain.Conversation, error) {
		return nil, apperr.ErrConversationNotFound{}
	}
	trip, err = ts.GetTrip("trip1", "user1")
	assert.ErrorAs(t, err, &apperr.ErrConversationNotFound{})
	assert.Nil(t, trip)
	// let's make GetConversationID return a conversation wıth messages on both sides

	getConversationByIDFunc = func(conversationID string) (*domain.Conversation, error) {
		return &domain.Conversation{
			TripID:      "trip1",
			RequesterID: "user1",
			SupplierID:  "user2",
			Messages: []domain.Message{
				{
					Direction: "in",
					Text:      "Hello",
					Read:      false,
				},
				{
					Direction: "out",
					Text:      "Hi",
					Read:      false,
				},
			},
		}, nil
	}
	markConversationsReadFunc = func(conversationID string, userID string) error {
		return nil
	}
	trip, err = ts.GetTrip("trip1", "user1")
	assert.Nil(t, err)
	assert.NotNil(t, trip)
	assert.Equal(t, "trip1", trip.ID)
	assert.Equal(t, 1, len(trip.Conversations))

}

func TestGetTripAsSupplier(t *testing.T) {
	mc := &MockContext{}
	mc.SetRepositories(nil, nil, nil, &mockTripRepository{}, &mockConversationRepository{}, nil)
	ts := NewTripService(mc)
	// Check the case when GetTripByID does not return an error, and returns a trip detail where the user is the supplier
	getTripByIDFunc = func(tripID string) (*domain.TripDetail, error) {
		return &domain.TripDetail{
			ID:         "trip1",
			SupplierID: "user1",
		}, nil
	}
	// in this case, the applicaton will also try to load the conversation list of this supplier under this trip
	// first check the case when GetConversations returns an error
	getConversationsFunc = func(tripID string) ([]domain.Conversation, error) {
		return nil, apperr.ErrConversationNotFound{}
	}
	trip, err := ts.GetTrip("trip1", "user1")
	assert.ErrorAs(t, err, &apperr.ErrConversationNotFound{})
	assert.Nil(t, trip)
	// check the case when GetConversations returns a slice of conversations
	getConversationsFunc = func(tripID string) ([]domain.Conversation, error) {
		return []domain.Conversation{
			{
				TripID:      "trip1",
				RequesterID: "user2",
				SupplierID:  "user1",
			},
			{
				TripID:      "trip1",
				RequesterID: "user3",
				SupplierID:  "user1",
			},
		}, nil
	}
	markConversationsReadFunc = func(conversationID string, userID string) error {
		return nil
	}
	trip, err = ts.GetTrip("trip1", "user1")
	assert.Nil(t, err)
	assert.NotNil(t, trip)
	assert.Equal(t, "trip1", trip.ID)
	assert.Equal(t, 2, len(trip.Conversations))
}

func TestGetTripByID(t *testing.T) {
	mc := &MockContext{}
	mc.SetRepositories(nil, nil, nil, &mockTripRepository{}, nil, nil)
	ts := NewTripService(mc)
	// Check the case when GetTripCapacity returns an error
	getTripByIDFunc = func(tripID string) (*domain.TripDetail, error) {
		return nil, apperr.ErrTripNotFound{}
	}
	trip, err := ts.GetTripByID("trip1")
	assert.ErrorAs(t, err, &apperr.ErrTripNotFound{})
	assert.Nil(t, trip)
	// Check the case when GetTripCapacity returns a trip detail
	getTripByIDFunc = func(tripID string) (*domain.TripDetail, error) {
		return &domain.TripDetail{
			ID: "trip1",
		}, nil
	}
	trip, err = ts.GetTripByID("trip1")
	assert.Nil(t, err)
	assert.NotNil(t, trip)
	assert.Equal(t, "trip1", trip.ID)
}

func TestSetTripCapacity(t *testing.T) {
	mc := &MockContext{}
	mc.SetRepositories(nil, nil, nil, &mockTripRepository{}, nil, nil)
	ts := NewTripService(mc)
	// Check the case when GetTripCapacity returns an error
	getTripCapacityFunc = func(tripID string) (int, error) {
		return 0, apperr.ErrTripNotFound{}
	}
	err := ts.SetTripCapacity("trip1", 2)
	assert.ErrorAs(t, err, &apperr.ErrTripNotFound{})
	// Check the case when GetTripCapacity does not return an error
	getTripCapacityFunc = func(tripID string) (int, error) {
		return 2, nil
	}
	// first check with a demand greater than the capacity
	err = ts.SetTripCapacity("trip1", -3)
	assert.ErrorAs(t, err, &apperr.ErrInvalidCapacity{})
	// check with a demand equal to the capacity
	// this time, the application will try to set the capacity of the trip
	// Check the case when SetTripCapacity returns an error
	setTripCapacityFunc = func(tripID string, capacity int) error {
		return apperr.ErrTripNotFound{}
	}
	err = ts.SetTripCapacity("trip1", -2)
	assert.ErrorAs(t, err, &apperr.ErrTripNotFound{})
	// Check the case when SetTripCapacity does not return an error
	setTripCapacityFunc = func(tripID string, capacity int) error {
		return nil
	}
	err = ts.SetTripCapacity("trip1", 2)
	assert.Nil(t, err)
}
