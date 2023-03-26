package mappers

import (
	"github.com/rs/zerolog/log"
	"github.com/serdarkalayci/carpool/api/adapters/data/mongodb/dao"
	"github.com/serdarkalayci/carpool/api/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func MapTripDAOs2Trips(tripsDAO []*dao.TripDAO) []*domain.Trip {
	var trips []*domain.Trip
	for _, tripDAO := range tripsDAO {
		trips = append(trips, MapTripDAO2Trip(tripDAO))
	}
	return trips
}

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

func MapTrip2TripDAO(trip domain.Trip) (dao.TripDAO, error) {
	supplierID, err := primitive.ObjectIDFromHex(trip.SupplierID)
	if err != nil {
		log.Error().Err(err).Msgf("Cannot parse ObjectID of SupplierID: %s", trip.SupplierID)
		return dao.TripDAO{}, err
	}
	countryID, err := primitive.ObjectIDFromHex(trip.CountryID)
	if err != nil {
		log.Error().Err(err).Msgf("Cannot parse ObjectID of CountryID: %s", trip.CountryID)
		return dao.TripDAO{}, err
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
