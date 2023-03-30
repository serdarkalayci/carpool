package rest

import (
	"net/http"

	"github.com/gorilla/mux"
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
		tripService := application.NewTripService(apiContext.tripRepo, apiContext.geographyRepo)
		err = tripService.AddTrip(trip)
		if err == nil {
			respondOK(rw, r, 200)
		} else if e, ok := err.(*domain.DuplicateKeyError); ok {
			respondWithError(rw, r, 400, e.Error())
		} else {
			log.Error().Err(err).Msg("error adding trip")
			respondWithError(rw, r, 500, "error adding trip")
		}
	} else {
		respondWithError(rw, r, 401, "Unauthorized")
	}
}

// swagger:route GET /trip Trip GetTrips
// Gets all the trips within the given country and optinally the given origin and destination
// responses:
//	200: OK
//	404: errorResponse

// GetTrips gets all the trips within the given country and optinally the given origin and destination
func (apiContext *APIContext) GetTrips(rw http.ResponseWriter, r *http.Request) {
	status, _, _ := checkLogin(r)
	if status {
		countryID := r.URL.Query().Get("countryid")
		origin := r.URL.Query().Get("origin")
		destination := r.URL.Query().Get("destination")
		tripService := application.NewTripService(apiContext.tripRepo, nil)
		trips, err := tripService.GetTrips(countryID, origin, destination)
		if err != nil {
			log.Error().Err(err).Msg("error getting trips")
			respondWithError(rw, r, 500, "error getting trips")
			return
		}
		respondWithJSON(rw, r, 200, mappers.MapTrips2TripListItems(trips))
	} else {
		respondWithError(rw, r, 401, "Unauthorized")
	}
}

// swagger:route GET /trip/{tripid}/message/{conversationid} Trip GetTrips
// Gets gets a specific conversation within a trip between the supplier and the requester
// responses:
//	200: OK
//	404: errorResponse

// GetConversation gets a specific conversation within a trip between the supplier and the requester
func (apiContext *APIContext) GetConversation(rw http.ResponseWriter, r *http.Request) {
	status, _, claims := checkLogin(r)
	if status {
		vars := mux.Vars(r)
		tripid := vars["tripid"]
		conversationid := vars["conversationid"]
		tripService := application.NewTripService(apiContext.tripRepo, nil)
		conversation, err := tripService.GetConversation(tripid, conversationid, claims.UserID)
		if err != nil {
			log.Error().Err(err).Msg("error getting conversation")
			respondWithError(rw, r, 500, "error getting conversation")
			return
		}
		respondWithJSON(rw, r, 200, mappers.MapConversation2ConversationResponse(*conversation))
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
		tripService := application.NewTripService(apiContext.tripRepo, nil)
		tripdetail, err := tripService.GetTrip(id, claims.UserID)
		if err != nil {
			log.Error().Err(err).Msg("error getting trips")
			respondWithError(rw, r, 500, "error getting trips")
			return
		}
		respondWithJSON(rw, r, 200, mappers.MapTripDetail2TripDetailResponse(*tripdetail))
	} else {
		respondWithError(rw, r, 401, "Unauthorized")
	}
}

// swagger:route POST /trip/{tripid}/conversation Conversation AddRequesterMessage
// Adds creates a new message to the trip by the requester.
// If there are no conversation initiated between the requester and the supplier, a new conversation is created.
// responses:
//	200: OK
//	404: errorResponse

// AddRequesterMessage creates a new message to the trip by the requester.
// If there are no conversation initiated between the requester and the supplier, a new conversation is created.
func (apiContext *APIContext) AddRequesterMessage(rw http.ResponseWriter, r *http.Request) {
	status, _, claims := checkLogin(r)
	if status {
		vars := mux.Vars(r)
		tripID := vars["tripid"]
		addMessageDTO := r.Context().Value(validatedMessage{}).(dto.AddMessageRequest)
		tripService := application.NewTripService(apiContext.tripRepo, nil)
		err := tripService.AddRequesterMessage(tripID, claims.UserID, claims.UserName, addMessageDTO.Text)
		if err == nil {
			respondOK(rw, r, 200)
		} else if e, ok := err.(*domain.DuplicateKeyError); ok {
			respondWithError(rw, r, 400, e.Error())
		} else {
			log.Error().Err(err).Msg("error adding message")
			respondWithError(rw, r, 500, "error adding message")
		}
	} else {
		respondWithError(rw, r, 401, "Unauthorized")
	}
}

// swagger:route POST /trip/{tripid}/conversation/{conversationid} Conversation AddSupplierMessage
// Adds creates a new message to the trip by the supplier.
// There should be a conversation initiated by a requester, so this method expects the conversationID.
// responses:
//	200: OK
//	404: errorResponse

// AddSupplierMessage creates a new message to the trip by the supplier.
func (apiContext *APIContext) AddSupplierMessage(rw http.ResponseWriter, r *http.Request) {
	status, _, claims := checkLogin(r)
	if status {
		vars := mux.Vars(r)
		tripID := vars["tripid"]
		conversationID := vars["conversationid"]
		addMessageDTO := r.Context().Value(validatedMessage{}).(dto.AddMessageRequest)
		tripService := application.NewTripService(apiContext.tripRepo, nil)
		err := tripService.AddSupplierMessage(tripID, claims.UserID, conversationID, addMessageDTO.Text)
		if err == nil {
			respondOK(rw, r, 200)
		} else if e, ok := err.(*domain.DuplicateKeyError); ok {
			respondWithError(rw, r, 400, e.Error())
		} else {
			log.Error().Err(err).Msg("error adding message")
			respondWithError(rw, r, 500, "error adding message")
		}
	} else {
		respondWithError(rw, r, 401, "Unauthorized")
	}
}
