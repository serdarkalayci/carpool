package application

import (
	"testing"

	apperr "github.com/serdarkalayci/carpool/api/application/errors"
	"github.com/serdarkalayci/carpool/api/domain"
	"github.com/stretchr/testify/assert"
)

type mockRequestRepository struct{}

var (
	addRequestFunc  func(request domain.Request) error
	getRequestsFunc func(countryID string, origin string, destination string) (*[]domain.Request, error)
	getRequestFunc  func(requestID string) (*domain.Request, error)
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
