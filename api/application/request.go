// Package application is the package that holds the application logic between database and communication layers
package application

import (
	"github.com/serdarkalayci/carpool/api/domain"
)

// RequestRepository is an interface for request repository. RequestService depends on this interface.
type RequestRepository interface {
	AddRequest(request domain.Request) error
	GetRequests(countryID string, origin string, destination string) (*[]domain.Request, error)
	GetRequest(requestID string) (*domain.Request, error)
	SetRequestStatus(requestID string, state int) error
}

// RequestService is the struct that holds the methods for request service. RequestHandler depends on this interface.
type RequestService struct {
	dc DataContextCarrier
}

// NewRequestService is the constructor for RequestService. It sets its underlying data context.
func NewRequestService(dc DataContextCarrier) RequestService {
	return RequestService{
		dc: dc,
	}
}

// GetRequests is the method that gets requests from the repository using its countryID, and filters using .
func (rs RequestService) GetRequests(countryID string, origin string, destination string) (*[]domain.Request, error) {
	return rs.dc.GetRequestRepository().GetRequests(countryID, origin, destination)
}

// AddRequest is the method that adds a request to the repository.
func (rs RequestService) AddRequest(request domain.Request) error {
	return rs.dc.GetRequestRepository().AddRequest(request)
}

// GetRequest is the method that gets a request from the repository bu its ID.
func (rs RequestService) GetRequest(requestID string) (*domain.Request, error) {
	return rs.dc.GetRequestRepository().GetRequest(requestID)
}

// RelateRequestToTrip is the method that relates a request to a trip by creating a conversation between them.
func (rs RequestService) RelateRequestToTrip(requestID string, tripID string) error {
	cs := NewConversationService(rs.dc)
	err := cs.InitiateConversationForRequest(tripID, requestID)
	if err != nil {
		return err
	}
	return rs.dc.GetRequestRepository().SetRequestStatus(requestID, 1)
}
