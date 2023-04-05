// Package mappers is the package that maps objects back and fort between dao and domain
package mappers

import (
	"github.com/rs/zerolog/log"
	"github.com/serdarkalayci/carpool/api/adapters/data/mongodb/dao"
	apperr "github.com/serdarkalayci/carpool/api/application/errors"
	"github.com/serdarkalayci/carpool/api/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// MapTripDAOs2Trips maps a slice of TripDAO to a slice of Trip
func MapTripDAOs2Trips(tripsDAO []*dao.TripDAO) []*domain.Trip {
	var trips []*domain.Trip
	for _, tripDAO := range tripsDAO {
		trips = append(trips, MapTripDAO2Trip(tripDAO))
	}
	return trips
}

// MapTripDAO2Trip maps a TripDAO to a Trip
func MapTripDAO2Trip(tripDAO *dao.TripDAO) *domain.Trip {
	return &domain.Trip{
		ID:             tripDAO.ID.Hex(),
		SupplierID:     tripDAO.SupplierID.Hex(),
		CountryID:      tripDAO.CountryID.Hex(),
		Origin:         tripDAO.Origin,
		Stops:          tripDAO.Stops,
		Destination:    tripDAO.Destination,
		TripDate:       tripDAO.TripDate.Time(),
		AvailableSeats: tripDAO.AvailableSeats,
		Note:           tripDAO.Note,
	}
}

// MapTrip2TripDAO maps a Trip to a TripDAO
func MapTrip2TripDAO(trip domain.Trip) (dao.TripDAO, error) {
	supplierID, err := primitive.ObjectIDFromHex(trip.SupplierID)
	if err != nil {
		log.Error().Err(err).Msgf("Cannot parse ObjectID of SupplierID: %s", trip.SupplierID)
		return dao.TripDAO{}, apperr.ErrInvalidID{Name: "SupplierID", Value: trip.SupplierID}
	}
	countryID, err := primitive.ObjectIDFromHex(trip.CountryID)
	if err != nil {
		log.Error().Err(err).Msgf("Cannot parse ObjectID of CountryID: %s", trip.CountryID)
		return dao.TripDAO{}, apperr.ErrInvalidID{Name: "CountryID", Value: trip.CountryID}
	}
	return dao.TripDAO{
		ID:             primitive.NewObjectID(),
		SupplierID:     supplierID,
		CountryID:      countryID,
		Origin:         trip.Origin,
		Destination:    trip.Destination,
		Stops:          trip.Stops,
		TripDate:       primitive.NewDateTimeFromTime(trip.TripDate),
		AvailableSeats: trip.AvailableSeats,
		Note:           trip.Note,
	}, nil
}

// MapTripDetailDAO2TripDetail maps a TripDetailDAO to a TripDetail
func MapTripDetailDAO2TripDetail(tripDetailDAO *dao.TripDetailDAO) *domain.TripDetail {
	return &domain.TripDetail{
		ID:             tripDetailDAO.ID.Hex(),
		SupplierID:     tripDetailDAO.SupplierID.Hex(),
		SupplierName:   tripDetailDAO.SupplierName,
		Country:        tripDetailDAO.Country,
		Origin:         tripDetailDAO.Origin,
		Stops:          tripDetailDAO.Stops,
		Destination:    tripDetailDAO.Destination,
		TripDate:       tripDetailDAO.TripDate.Time(),
		AvailableSeats: tripDetailDAO.AvailableSeats,
		Note:           tripDetailDAO.Note,
	}
}
