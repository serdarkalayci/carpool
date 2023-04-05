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

type validatedTrip struct{}

// extractAddTripPayload extracts user data from the request body
// Returns TripRequest model if found, error otherwise
func extractAddTripPayload(r *http.Request) (trip *dto.AddTripRequest, e error) {
	payload, e := readPayload(r)
	if e != nil {
		return
	}
	err := json.Unmarshal(payload, &trip)
	if err != nil {
		e = &apierr.ErrCannotParsePayload{}
		log.Error().Err(err).Msg(e.Error())
		return
	}
	return
}

// validateNewTrip Checks the integrity of new user in the request and calls next if ok
func (apiContext *APIContext) validateNewTrip(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		user, err := extractAddTripPayload(r)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}
		// validate the user
		errs := apiContext.validation.Validate(user)
		if errs != nil && len(errs) != 0 {
			log.Error().Err(errs[0]).Msg("error validating the user")

			// return the validation messages as an array
			respondWithJSON(rw, r, http.StatusUnprocessableEntity, errs.Errors())
			return
		}

		// add the rating to the context
		ctx := context.WithValue(r.Context(), validatedTrip{}, *user)
		r = r.WithContext(ctx)

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(rw, r)
	})
}
