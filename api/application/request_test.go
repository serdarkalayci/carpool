package application

import (
	"testing"

	apperr "github.com/serdarkalayci/carpool/api/application/errors"
	"github.com/serdarkalayci/carpool/api/domain"
	"github.com/stretchr/testify/assert"
)

type mockRequestRepository struct{}

var (
	addRequestFunc       func(request domain.Request) error
	getRequestsFunc      func(countryID string, origin string, destination string) (*[]domain.Request, error)
	getRequestFunc       func(requestID string) (*domain.Request, error)
	setRequestStatusFunc func(requestID string, state int) error
)

// AddRequest adds a new trip request to the database
func (m mockRequestRepository) AddRequest(request domain.Request) error {
	return addRequestFunc(request)
}

// GetRequests returns all the requests in the database
func (m mockRequestRepository) GetRequests(countryID string, origin string, destination string) (*[]domain.Request, error) {
	return getRequestsFunc(countryID, origin, destination)
}

// GetRequest returns a request with the given ID
func (m mockRequestRepository) GetRequest(requestID string) (*domain.Request, error) {
	return getRequestFunc(requestID)
}

// SetRequestStatus sets the status of a request
func (m mockRequestRepository) SetRequestStatus(requestID string, state int) error {
	return setRequestStatusFunc(requestID, state)
}

func TestAddRequest(t *testing.T) {
	mc := &MockContext{}
	mc.SetRepositories(nil, nil, nil, nil, nil, &mockRequestRepository{})
	rs := NewRequestService(mc)
	addRequestFunc = func(request domain.Request) error {
		return apperr.ErrRequestNotFound{}
	}
	err := rs.AddRequest(domain.Request{})
	assert.ErrorAs(t, err, &apperr.ErrRequestNotFound{})
	addRequestFunc = func(request domain.Request) error {
		return nil
	}
	err = rs.AddRequest(domain.Request{})
	assert.NoError(t, err)
}

func TestGetRequests(t *testing.T) {
	mc := &MockContext{}
	mc.SetRepositories(nil, nil, nil, nil, nil, &mockRequestRepository{})
	rs := NewRequestService(mc)
	getRequestsFunc = func(countryID string, origin string, destination string) (*[]domain.Request, error) {
		return nil, apperr.ErrRequestNotFound{}
	}
	_, err := rs.GetRequests("", "", "")
	assert.ErrorAs(t, err, &apperr.ErrRequestNotFound{})
	getRequestsFunc = func(countryID string, origin string, destination string) (*[]domain.Request, error) {
		return nil, nil
	}
	_, err = rs.GetRequests("", "", "")
	assert.NoError(t, err)
}

func TestGetRequest(t *testing.T) {
	mc := &MockContext{}
	mc.SetRepositories(nil, nil, nil, nil, nil, &mockRequestRepository{})
	rs := NewRequestService(mc)
	getRequestFunc = func(requestID string) (*domain.Request, error) {
		return nil, apperr.ErrRequestNotFound{}
	}
	_, err := rs.GetRequest("")
	assert.ErrorAs(t, err, &apperr.ErrRequestNotFound{})
	getRequestFunc = func(requestID string) (*domain.Request, error) {
		return nil, nil
	}
	_, err = rs.GetRequest("")
	assert.NoError(t, err)
}

func TestRelateRequestToTrip(t *testing.T) {
	mc := &MockContext{}
	mc.SetRepositories(&mockUserRepository{}, nil, nil, &mockTripRepository{}, &mockConversationRepository{}, &mockRequestRepository{})
	rs := NewRequestService(mc)
	// First check the case where InitiateConversationForRequest fails
	getTripByIDFunc = func(tripID string) (*domain.TripDetail, error) {
		return nil, apperr.ErrTripNotFound{}
	}
	err := rs.RelateRequestToTrip("request1", "trip1")
	assert.ErrorAs(t, err, &apperr.ErrTripNotFound{})
	getTripByIDFunc = func(tripID string) (*domain.TripDetail, error) {
		return &domain.TripDetail{
			ID:             "trip1",
			SupplierID:     "supplier1",
			AvailableSeats: 3,
		}, nil
	}
	getRequestFunc = func(requestID string) (*domain.Request, error) {
		return &domain.Request{
			ID:             "request1",
			RequesterID:    "requester1",
			RequesterName:  "requester1",
			RequestedSeats: 2,
		}, nil
	}
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
	initiateConversationFunc = func(conversation domain.Conversation) error {
		return nil
	}
	setRequestStatusFunc = func(requestID string, state int) error {
		return apperr.ErrRequestNotFound{}
	}
	err = rs.RelateRequestToTrip("request1", "trip1")
	assert.ErrorAs(t, err, &apperr.ErrRequestNotFound{})
	setRequestStatusFunc = func(requestID string, state int) error {
		return nil
	}
	err = rs.RelateRequestToTrip("request1", "trip1")
	assert.NoError(t, err)
}
