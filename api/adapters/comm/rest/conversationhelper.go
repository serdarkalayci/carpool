// Package rest is responsible for rest communication layer
package rest

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/rs/zerolog/log"
	"github.com/serdarkalayci/carpool/api/adapters/comm/rest/dto"
	apierr "github.com/serdarkalayci/carpool/api/adapters/comm/rest/errors"
)

type validatedConversation struct{}

func extractAddConversationPayload(r *http.Request) (conversation *dto.AddConversationRequest, e error) {
	payload, e := readPayload(r)
	if e != nil {
		return
	}
	err := json.Unmarshal(payload, &conversation)
	if err != nil {
		e = &apierr.ErrCannotParsePayload{}
		log.Error().Err(err).Msg(e.Error())
		return
	}
	return
}

func (apiContext *APIContext) validateNewConversation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		conversation, err := extractAddConversationPayload(r)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}
		// validate the conversation
		errs := apiContext.validation.Validate(conversation)
		if errs != nil && len(errs) != 0 {
			log.Error().Err(errs[0]).Msg("error validating the conversation")

			// return the validation conversations as an array
			respondWithJSON(rw, r, http.StatusUnprocessableEntity, errs.Errors())
			return
		}

		// add the rating to the context
		ctx := context.WithValue(r.Context(), validatedConversation{}, *conversation)
		r = r.WithContext(ctx)

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(rw, r)
	})
}

type validatedMessage struct{}

func extractAddMessagePayload(r *http.Request) (message *dto.AddMessageRequest, e error) {
	payload, e := readPayload(r)
	if e != nil {
		return
	}
	err := json.Unmarshal(payload, &message)
	if err != nil {
		e = &apierr.ErrCannotParsePayload{}
		log.Error().Err(e).Msg(e.Error())
		return
	}
	return
}

func (apiContext *APIContext) validateNewMessage(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		message, err := extractAddMessagePayload(r)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}
		// validate the message
		errs := apiContext.validation.Validate(message)
		if errs != nil && len(errs) != 0 {
			log.Error().Err(errs[0]).Msg("error validating the message")

			// return the validation messages as an array
			respondWithJSON(rw, r, http.StatusUnprocessableEntity, errs.Errors())
			return
		}

		// add the rating to the context
		ctx := context.WithValue(r.Context(), validatedMessage{}, *message)
		r = r.WithContext(ctx)

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(rw, r)
	})
}

type validatedApproval struct{}

func extractUpdateApprovalPayload(r *http.Request) (approval *dto.UpdateApprovalRequest, e error) {
	payload, e := readPayload(r)
	if e != nil {
		return
	}
	err := json.Unmarshal(payload, &approval)
	if err != nil {
		e = &apierr.ErrCannotParsePayload{}
		log.Error().Err(err).Msg(e.Error())
		return
	}
	return
}

func (apiContext *APIContext) validateUpdateApproval(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		approval, err := extractUpdateApprovalPayload(r)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}
		// validate the approval
		errs := apiContext.validation.Validate(approval)
		if errs != nil && len(errs) != 0 {
			log.Error().Err(errs[0]).Msg("error validating the approval")

			// return the validation approvals as an array
			respondWithJSON(rw, r, http.StatusUnprocessableEntity, errs.Errors())
			return
		}

		// add the rating to the context
		ctx := context.WithValue(r.Context(), validatedApproval{}, *approval)
		r = r.WithContext(ctx)

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(rw, r)
	})
}
