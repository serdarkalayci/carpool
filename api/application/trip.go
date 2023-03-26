package application

import "github.com/serdarkalayci/carpool/api/domain"

type TripRepository interface {
	AddTrip(trip *domain.Trip) error
	GetTrips(countryID string, origin, destination string) ([]*domain.Trip, error)
}

// TripService is the struct to let outer layers to interact to the Trip Applicatopn
type TripService struct {
	tripRepository TripRepository
}

// NewTripService creates a new TripService instance and sets its repository
func NewTripService(tr TripRepository) TripService {
	if tr == nil {
		panic("missing tripRepository")
	}
	return TripService{
		tripRepository: tr,
	}
}

func (ts TripService) AddTrip(trip domain.Trip) error {
	return ts.tripRepository.AddTrip(&trip)
}

func (ts TripService) GetTrips(countryID string, origin, destination string) ([]*domain.Trip, error) {
	return ts.tripRepository.GetTrips(countryID, origin, destination)
}
