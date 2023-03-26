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

// swagger:route GET /user/{id} User GetUser
// Return the user if found
// responses:
//	200: OK
//	404: errorResponse

// GetUser gets a single user if found
func (apiContext *APIContext) GetUser(rw http.ResponseWriter, r *http.Request) {
	// parse the Rating id from the url
	vars := mux.Vars(r)

	// convert the id into an integer and return
	id := vars["id"]
	userService := application.NewUserService(apiContext.userRepo)
	user, err := userService.GetUser(id)
	if err == nil {
		respondWithJSON(rw, r, 200, mappers.MapUser2SUserResponse(user))
	}
}

// swagger:route POST /user/{id} User AddUser
// Adds a new user to the system
// responses:
//	200: OK
//	404: errorResponse

// AddUser creates a new user on the system
func (apiContext *APIContext) AddUser(rw http.ResponseWriter, r *http.Request) {
	// Get user data from payload
	userDTO := r.Context().Value(validatedUser{}).(dto.AddUserRequest)
	user := mappers.MapAddUserRequest2User(userDTO)
	userService := application.NewUserService(apiContext.userRepo)
	err := userService.AddUser(user)
	if err == nil {
		respondOK(rw, r, 200)
	} else if e, ok := err.(*domain.DuplicateKeyError); ok {
		respondWithError(rw, r, 400, e.Error())
	} else {
		log.Error().Err(err).Msg("error adding user")
		respondWithError(rw, r, 500, "error adding user")
	}
}

// ConfirmUser confirms a user if found
func (apiContext *APIContext) ConfirmUser(rw http.ResponseWriter, r *http.Request) {
	// parse the Rating id from the url
	vars := mux.Vars(r)

	// convert the id into an integer and return
	userid := vars["id"]
	confirmationCode := vars["code"]
	userService := application.NewUserService(apiContext.userRepo)
	err := userService.CheckConfirmationCode(userid, confirmationCode)
	if err != nil {
		respondWithError(rw, r, 401, "user not confirmed")
		return
	}
	respondOK(rw, r, 200)
}
