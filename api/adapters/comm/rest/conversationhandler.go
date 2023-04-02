package rest

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
	"github.com/serdarkalayci/carpool/api/adapters/comm/rest/dto"
	"github.com/serdarkalayci/carpool/api/adapters/comm/rest/mappers"
	"github.com/serdarkalayci/carpool/api/application"
)

// swagger:route GET /conversation/{conversationid} Conversation GetConmversation
// Gets gets a specific conversation within a trip between the supplier and the requester
// responses:
//	200: OK
//	404: errorResponse

// GetConversation gets a specific conversation within a trip between the supplier and the requester
func (apiContext *APIContext) GetConversation(rw http.ResponseWriter, r *http.Request) {
	status, _, claims := checkLogin(r)
	if status {
		vars := mux.Vars(r)
		conversationid := vars["conversationid"]
		tripService := application.NewConversationService(apiContext.dbContext)
		conversation, err := tripService.GetConversation(conversationid, claims.UserID)
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

// swagger:route POST /conversation Conversation AddConversation
// Adds creates a new conversation for a trip.
// responses:
//	200: OK
//	404: errorResponse

// AddConversation creates a new conversation for a trip.
func (apiContext *APIContext) AddConversation(rw http.ResponseWriter, r *http.Request) {
	status, _, claims := checkLogin(r)
	if status {
		addConversationDTO := r.Context().Value(validatedConversation{}).(dto.AddConversationRequest)
		conversationService := application.NewConversationService(apiContext.dbContext)
		err := conversationService.InitiateConversation(addConversationDTO.TripID, claims.UserID, addConversationDTO.Capacity, addConversationDTO.Message)
		if err == nil {
			respondOK(rw, r, 200)
		} else {
			log.Error().Err(err).Msg("error adding conversation")
			respondWithError(rw, r, 500, "error adding conversation")
		}
	} else {
		respondWithError(rw, r, 401, "Unauthorized")
	}
}

// swagger:route PUT /conversation/{conversationid} Conversation AddMessage
// Adds a new message to the conversation.
// responses:
//	200: OK
//	404: errorResponse

// AddMessage creates a new message to the conversation.
func (apiContext *APIContext) AddMessage(rw http.ResponseWriter, r *http.Request) {
	status, _, claims := checkLogin(r)
	if status {
		vars := mux.Vars(r)
		conversationID := vars["conversationid"]
		addMessageDTO := r.Context().Value(validatedMessage{}).(dto.AddMessageRequest)
		tripService := application.NewConversationService(apiContext.dbContext)
		err := tripService.AddMessage(conversationID, claims.UserID, addMessageDTO.Text)
		if err == nil {
			respondOK(rw, r, 200)
		} else {
			log.Error().Err(err).Msg("error adding message")
			respondWithError(rw, r, 500, "error adding message")
		}
	} else {
		respondWithError(rw, r, 401, "Unauthorized")
	}
}

// swagger:route PUT /conversation/{conversationid}/approve Conversation UpdateApproval
// Updates the approval status of the conversation.
// responses:
//	200: OK
//	404: errorResponse

// UpdateApproval updates the approval status of the conversation.
func (apiContext *APIContext) UpdateApproval(rw http.ResponseWriter, r *http.Request) {
	status, _, claims := checkLogin(r)
	if status {
		vars := mux.Vars(r)
		conversationID := vars["conversationid"]
		updateApprovalDTO := r.Context().Value(validatedApproval{}).(dto.UpdateApprovalRequest)
		tripService := application.NewConversationService(apiContext.dbContext)
		err := tripService.UpdateApproval(conversationID, claims.UserID, *updateApprovalDTO.Approved)
		if err == nil {
			respondOK(rw, r, 200)
		} else {
			log.Error().Err(err).Msg("error updating approval")
			respondWithError(rw, r, 500, "error updating approval")
		}
	} else {
		respondWithError(rw, r, 401, "Unauthorized")
	}
}
