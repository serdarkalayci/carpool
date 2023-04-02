package rest

import (
	"net/http"

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
			respondWithError(rw, r, 400, "error mapping request")
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
			respondWithError(rw, r, 500, "error adding request")
		}
	} else {
		respondWithError(rw, r, 401, "Unauthorized")
	}
}
