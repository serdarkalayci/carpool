// Package rest is responsible for rest communication layer
package rest

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/serdarkalayci/carpool/api/adapters/comm/rest/dto"
	"github.com/serdarkalayci/carpool/api/adapters/comm/rest/mappers"
	"github.com/serdarkalayci/carpool/api/application"
	apperr "github.com/serdarkalayci/carpool/api/application/errors"
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
			respondWithError(rw, r, 400, err.Error())
			return
		}
		trip.SupplierID = claims.UserID
		tripService := application.NewTripService(apiContext.dbContext)
		err = tripService.AddTrip(trip)
		if err == nil {
			respondOK(rw, r, 200)
		} else if e, ok := err.(apperr.DuplicateKeyError); ok {
			respondWithError(rw, r, 400, e.Error())
		} else {
			respondWithError(rw, r, 500, err.Error())
		}
	} else {
		respondWithError(rw, r, 401, "Unauthorized")
	}
}

// swagger:route GET /trip Trip GetTrips
// Gets all the trips within the given country and optionally the given origin and destination
// responses:
//	200: OK
//	404: errorResponse

// GetTrips gets all the trips within the given country and optionally the given origin and destination
func (apiContext *APIContext) GetTrips(rw http.ResponseWriter, r *http.Request) {
	status, _, _ := checkLogin(r)
	if status {
		countryID := r.URL.Query().Get("countryid")
		origin := r.URL.Query().Get("origin")
		destination := r.URL.Query().Get("destination")
		tripService := application.NewTripService(apiContext.dbContext)
		trips, err := tripService.GetTrips(countryID, origin, destination)
		if err != nil {
			respondWithError(rw, r, 500, err.Error())
			return
		}
		respondWithJSON(rw, r, 200, mappers.MapTrips2TripListItems(trips))
	} else {
		respondWithError(rw, r, 401, "Unauthorized")
	}
}

// swagger:route GET /trip/{id} Trip GetTrip
// Gets a single trip with all the details
// responses:
//	200: OK
//	404: errorResponse

// GetTrip Gets a single trip with all the details
func (apiContext *APIContext) GetTrip(rw http.ResponseWriter, r *http.Request) {
	status, _, claims := checkLogin(r)
	if status {
		vars := mux.Vars(r)
		id := vars["id"]
		tripService := application.NewTripService(apiContext.dbContext)
		tripdetail, err := tripService.GetTrip(id, claims.UserID)
		if err != nil {
			respondWithError(rw, r, 500, err.Error())
			return
		}
		respondWithJSON(rw, r, 200, mappers.MapTripDetail2TripDetailResponse(*tripdetail))
	} else {
		respondWithError(rw, r, 401, "Unauthorized")
	}
}
