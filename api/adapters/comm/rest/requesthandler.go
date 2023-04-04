package rest

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
	"github.com/serdarkalayci/carpool/api/adapters/comm/rest/dto"
	"github.com/serdarkalayci/carpool/api/adapters/comm/rest/mappers"
	"github.com/serdarkalayci/carpool/api/application"
)

// swagger:route POST /request Request AddRequest
// Adds a new request to the system
// responses:
//	200: OK
//	404: errorResponse

// AddRequest creates a new request on the system
func (apiContext *APIContext) AddRequest(rw http.ResponseWriter, r *http.Request) {
	status, _, claims := checkLogin(r)
	if status {
		requestDTO := r.Context().Value(validatedRequest{}).(dto.AddRequestRequest)
		request, err := mappers.MapAddRequestRequest2Request(requestDTO)
		if err != nil {
			log.Error().Err(err).Msg("error mapping request")
			respondWithError(rw, r, 400, err.Error())
			return
		}
		request.RequesterID = claims.UserID
		request.RequesterName = claims.UserName
		requestService := application.NewRequestService(apiContext.dbContext)
		err = requestService.AddRequest(request)
		if err == nil {
			respondOK(rw, r, 200)
		} else {
			log.Error().Err(err).Msg("error adding request")
			respondWithError(rw, r, 500, err.Error())
		}
	} else {
		respondWithError(rw, r, 401, "Unauthorized")
	}
}

func (apiContext *APIContext) GetRequests(rw http.ResponseWriter, r *http.Request) {
	status, _, _ := checkLogin(r)
	if status {
		requestService := application.NewRequestService(apiContext.dbContext)
		countryID := r.URL.Query().Get("countryid")
		origin := r.URL.Query().Get("origin")
		destination := r.URL.Query().Get("destination")
		requests, err := requestService.GetRequests(countryID, origin, destination)
		if err == nil {
			requestListResponse := mappers.MapRequests2RequestListResponses(requests)
			respondWithJSON(rw, r, 200, requestListResponse)
		} else {
			log.Error().Err(err).Msg("error getting requests")
			respondWithError(rw, r, 500, err.Error())
		}
	} else {
		respondWithError(rw, r, 401, "Unauthorized")
	}
}

func (apiContext *APIContext) GetRequest(rw http.ResponseWriter, r *http.Request) {
	status, _, _ := checkLogin(r)
	if status {
		requestService := application.NewRequestService(apiContext.dbContext)
		vars := mux.Vars(r)
		requestID := vars["requestid"]
		request, err := requestService.GetRequest(requestID)
		if err == nil {
			requestListResponse := mappers.MapRequest2RequestResponse(request)
			respondWithJSON(rw, r, 200, requestListResponse)
		} else {
			log.Error().Err(err).Msg("error getting requests")
			respondWithError(rw, r, 500, err.Error())
		}
	} else {
		respondWithError(rw, r, 401, "Unauthorized")
	}
}
