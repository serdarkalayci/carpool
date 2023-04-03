package application

import "github.com/serdarkalayci/carpool/api/domain"

type RequestRepository interface {
	AddRequest(request domain.Request) error
	GetRequests(countryID string, origin string, destination string) (*[]domain.Request, error)
	GetRequest(requestID string) (*domain.Request, error)
}

type RequestService struct {
	dc DataContext
}

func NewRequestService(dc DataContext) RequestService {
	return RequestService{
		dc: dc,
	}
}

func (rs RequestService) GetRequests(countryID string, origin string, destination string) (*[]domain.Request, error) {
	return rs.dc.RequestRepository.GetRequests(countryID, origin, destination)
}

func (rs RequestService) AddRequest(request domain.Request) error {
	return rs.dc.RequestRepository.AddRequest(request)
}

func (rs RequestService) GetRequest(requestID string) (*domain.Request, error) {
	return rs.dc.RequestRepository.GetRequest(requestID)
}
