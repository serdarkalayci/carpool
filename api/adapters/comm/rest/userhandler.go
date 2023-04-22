// Package rest is responsible for rest communication layer
package rest

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/serdarkalayci/carpool/api/adapters/comm/rest/dto"
	"github.com/serdarkalayci/carpool/api/adapters/comm/rest/mappers"
	"github.com/serdarkalayci/carpool/api/application"
	apperr "github.com/serdarkalayci/carpool/api/application/errors"
)

// swagger:route GET /user/{userid} User GetUser
// Return the user if found
// responses:
//	200: OK
//	404: errorResponse

// GetUser gets a single user if found
func (apiContext *APIContext) GetUser(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userid := vars["userid"]
	userService := application.NewUserService(apiContext.dbContext)
	user, err := userService.GetUser(userid)
	if err == nil {
		respondWithJSON(rw, r, 200, mappers.MapUser2SUserResponse(user))
	}
}

// swagger:route POST /user User AddUser
// Adds a new user to the system
// responses:
//	200: OK
//	404: errorResponse

// AddUser creates a new user on the system
func (apiContext *APIContext) AddUser(rw http.ResponseWriter, r *http.Request) {
	// Get user data from payload
	userDTO := r.Context().Value(validatedUser{}).(dto.AddUserRequest)
	user := mappers.MapAddUserRequest2User(userDTO)
	userService := application.NewUserService(apiContext.dbContext)
	err := userService.AddUser(user)
	if err == nil {
		respondOK(rw, r, 200)
	} else if e, ok := err.(apperr.DuplicateKeyError); ok {
		respondWithError(rw, r, 400, e.Error())
	} else {
		respondWithError(rw, r, 500, err.Error())
	}
}

// swagger:route PUT /user/{userid} User GetUser
// Return the user if found
// responses:
//	200: OK
//	404: errorResponse

// ConfirmUser confirms a user if found
func (apiContext *APIContext) ConfirmUser(rw http.ResponseWriter, r *http.Request) {
	// parse the Rating id from the url
	vars := mux.Vars(r)
	userid := vars["userid"]
	confirmation := r.Context().Value(validatedConfirmUser{}).(dto.ConfirmUserRequest)
	userService := application.NewUserService(apiContext.dbContext)
	err := userService.CheckConfirmationCode(userid, confirmation.Code)
	if err != nil {
		respondWithError(rw, r, 401, err.Error())
		return
	}
	respondOK(rw, r, 200)
}
