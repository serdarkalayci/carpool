package application

import "github.com/serdarkalayci/carpool/api/domain"

type RequestRepository interface {
	AddRequest(request domain.Request) error
	GetRequests(countryID string) (*[]domain.Request, error)
}

type RequestService struct {
	dc DataContext
}

func NewRequestService(dc DataContext) RequestService {
	return RequestService{
		dc: dc,
	}
}

func (rs RequestService) GetRequests(countryID string) (*[]domain.Request, error) {
	return rs.dc.RequestRepository.GetRequests(countryID)
}

func (rs RequestService) AddRequest(request domain.Request) error {
	return rs.dc.RequestRepository.AddRequest(request)
}
