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
		tripService := application.NewConversationService(apiContext.conversationRepo, nil, nil)
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

// swagger:route POST /conversation/{conversationid} Conversation AddMessage
// Adds creates a new message to the conversation.
// responses:
//	200: OK
//	404: errorResponse

// AddSupplierMessage creates a new message to the trip by the supplier.
func (apiContext *APIContext) AddMessage(rw http.ResponseWriter, r *http.Request) {
	status, _, claims := checkLogin(r)
	if status {
		vars := mux.Vars(r)
		conversationID := vars["conversationid"]
		addMessageDTO := r.Context().Value(validatedMessage{}).(dto.AddMessageRequest)
		tripService := application.NewConversationService(apiContext.conversationRepo, nil, nil)
		err := tripService.AddMessage(conversationID, claims.UserID, addMessageDTO.Text)
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
