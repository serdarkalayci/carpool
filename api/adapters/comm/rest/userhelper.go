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

type validatedUser struct{}

// ExtractAddUserPayload extracts user data from the request body
// Returns UserRequest model if found, error otherwise
func extractAddUserPayload(r *http.Request) (user *dto.AddUserRequest, e error) {
	payload, e := readPayload(r)
	if e != nil {
		return
	}
	err := json.Unmarshal(payload, &user)
	if err != nil {
		e = errors.New(viper.GetString("CannotParsePayloadMsg"))
		log.Error().Err(err).Msg(viper.GetString("CannotParsePayloadMsg"))
		return
	}
	return
}

// validateNewUser Checks the integrity of new user in the request and calls next if ok
func (apiContext *APIContext) validateNewUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		user, err := extractAddUserPayload(r)
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

		ctx := context.WithValue(r.Context(), validatedUser{}, *user)
		r = r.WithContext(ctx)

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(rw, r)
	})
}
