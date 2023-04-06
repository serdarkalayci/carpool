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
	trip, err = ts.GetTrip("trip1", "user1")
	assert.ErrorAs(t, err, &apperr.ErrConversationNotFound{})
	assert.Nil(t, trip)
	// let's make GetConversationID return a conversation

}
