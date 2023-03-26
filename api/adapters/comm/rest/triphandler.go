package rest

import (
	"net/http"

	"github.com/rs/zerolog/log"
	"github.com/serdarkalayci/carpool/api/adapters/comm/rest/dto"
	"github.com/serdarkalayci/carpool/api/adapters/comm/rest/mappers"
	"github.com/serdarkalayci/carpool/api/application"
	"github.com/serdarkalayci/carpool/api/domain"
)

// swagger:route POST /trip Trip AddTrip
// Adds a new trip to the system
// responses:
//	200: OK
//	404: errorResponse

// AddTrip creates a new trip on the system
func (apiContext *APIContext) AddTrip(rw http.ResponseWriter, r *http.Request) {
	status, _, claims := checkLogin(r)
	if status {
		tripDTO := r.Context().Value(validatedTrip{}).(dto.AddTripRequest)
		trip, err := mappers.MapAddTripRequest2Trip(tripDTO)
		if err != nil {
			log.Error().Err(err).Msg("error mapping trip")
			respondWithError(rw, r, 400, "error mapping trip")
			return
		}
		trip.SupplierID = claims.UserID
		tripService := application.NewTripService(apiContext.tripRepo)
		err = tripService.AddTrip(trip)
		if err == nil {
			respondOK(rw, r, 200)
		} else if e, ok := err.(*domain.DuplicateKeyError); ok {
			respondWithError(rw, r, 400, e.Error())
		} else {
			log.Error().Err(err).Msg("error adding user")
			respondWithError(rw, r, 500, "error adding user")
		}
	} else {
		respondWithError(rw, r, 401, "Unauthorized")
	}
}
