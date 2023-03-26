package mappers

import (
	"time"

	"github.com/serdarkalayci/carpool/api/adapters/comm/rest/dto"
	"github.com/serdarkalayci/carpool/api/domain"
)

func MapTrip2TripListItem(trip domain.Trip) dto.TripListItem {
	return dto.TripListItem{
		ID:             trip.ID,
		Origin:         trip.Origin,
		Destination:    trip.Destination,
		TripDate:       trip.TripDate,
		AvailableSeats: trip.AvailableSeats,
	}
}

func MapAddTripRequest2Trip(trip dto.AddTripRequest) (domain.Trip, error) {
	tripDate, err := time.Parse("2006-01-02", trip.TripDate)
	if err != nil {
		return domain.Trip{}, err
	}
	return domain.Trip{
		CountryID:      trip.CountryID,
		Origin:         trip.Origin,
		Destination:    trip.Destination,
		Stops:          trip.Stops,
		TripDate:       tripDate,
		AvailableSeats: trip.AvailableSeats,
		Note:           trip.Note,
	}, nil
}
