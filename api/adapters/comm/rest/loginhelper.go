package rest

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/rs/zerolog/log"
	"github.com/serdarkalayci/carpool/api/adapters/comm/rest/dto"
	"github.com/spf13/viper"
)

type ValidatedLogin struct{}

// extractLoginPayload extracts login data from the request body
// Returns LoginRequest model if found, error otherwise
func extractLoginPayload(r *http.Request) (login *dto.LoginRequest, e error) {
	payload, e := readPayload(r)
	if e != nil {
		return
	}
	err := json.Unmarshal(payload, &login)
	if err != nil {
		e = errors.New(viper.GetString("CannotParsePayloadMsg"))
		log.Error().Err(err).Msg(viper.GetString("CannotParsePayloadMsg"))
		return
	}
	return
}

// validateLoginRequest Checks the integrity of login information in the request and calls next if ok
func (apiContext *APIContext) validateLoginRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		login, err := extractLoginPayload(r)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}
		// validate the login
		errs := apiContext.validation.Validate(login)
		if errs != nil && len(errs) != 0 {
			log.Error().Err(errs[0]).Msg("error validating the login")

			// return the validation messages as an array
			respondWithJSON(rw, r, http.StatusUnprocessableEntity, errs.Errors())
			return
		}

		// add the rating to the context
		ctx := context.WithValue(r.Context(), ValidatedLogin{}, *login)
		r = r.WithContext(ctx)

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(rw, r)
	})
}
