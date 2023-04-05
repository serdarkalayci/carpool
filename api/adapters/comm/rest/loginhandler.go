// Package rest is responsible for rest communication layer
package rest

import (
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"

	"github.com/serdarkalayci/carpool/api/application"

	"github.com/rs/zerolog/log"
	"github.com/serdarkalayci/carpool/api/adapters/comm/rest/dto"
)

// Claims represents the data for a user login
type Claims struct {
	UserID   string               `json:"userid"`
	UserName string               `json:"username"`
	Payload  jwt.RegisteredClaims `json:"payload"`
}

// Valid checks the validity of the claims
func (c Claims) Valid() error {
	return c.Payload.Valid()
}

const secretKey = "the_most_secure_secret"
const cookieName = "carpooltoken"

var hs = []byte(secretKey)

// Login swagger:route POST PUT /user Login
//
// # Handler to log in the user, returns a JWT Token
//
// Responses:
//
//	       200: OK
//			  400: Bad Request
//			  500: Internal Server Error

// Login handles the login request
func (apiContext *APIContext) Login(w http.ResponseWriter, r *http.Request) {
	userLogin := r.Context().Value(validatedLogin{}).(dto.LoginRequest)
	userService := application.NewUserService(apiContext.dbContext)
	user, err := userService.CheckUser(userLogin.Email, userLogin.Password)
	if err != nil {
		respondWithError(w, r, 401, "User not found")
		log.Error().Err(err).Msg("user not found")
		return
	}
	// Create the JWT claims, which includes the username and expiry time
	now := time.Now()
	rclaims := jwt.RegisteredClaims{
		Audience:  jwt.ClaimStrings{"https://carpool.io"},
		ExpiresAt: jwt.NewNumericDate(now.Add(30 * time.Minute)),
		ID:        "carpool",
		IssuedAt:  jwt.NewNumericDate(now),
		Issuer:    "carpool",
		NotBefore: jwt.NewNumericDate(now),
		Subject:   "carpoollogin",
	}

	pl := Claims{
		Payload:  rclaims,
		UserID:   user.ID,
		UserName: user.Name,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, pl)

	tokenstring, err := token.SignedString(hs)
	if err != nil {
		log.Error().Err(err).Msg("error creating the token")
		respondWithError(w, r, 500, "Token creation failed")
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	// Finally, we set the client cookie for "token" as the JWT we just generated
	// we also set an expiry time which is the same as the token itself
	http.SetCookie(w, &http.Cookie{
		Name:  cookieName,
		Value: tokenstring,
		Path:  "/",
	})
}

func checkLogin(r *http.Request) (status bool, httpStatusCode int, claims *Claims) {
	// We can obtain the session token from the requests cookies, which come with every request
	// Initialize a new instance of `Claims`
	c, err := r.Cookie(cookieName)
	status = false
	if err != nil {
		if err == http.ErrNoCookie {
			// If the cookie is not set, return an unauthorized status
			httpStatusCode = http.StatusUnauthorized
			return
		}
		// For any other type of error, return a bad request status
		httpStatusCode = http.StatusBadRequest
		return
	}

	// Get the JWT string from the cookie
	tokenstring := c.Value
	// Initialize a new instance of `Claims`
	claims = &Claims{}

	token, err := jwt.ParseWithClaims(tokenstring, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return hs, nil
	})
	if err != nil {
		log.Error().Err(err).Msg("error validating the token")
		httpStatusCode = http.StatusUnauthorized
		return
	}
	claims = token.Claims.(*Claims)
	status = true
	return
}

// Refresh swagger:route POST PUT /login Refresh
//
// # Handler to refresh a JWT Token
//
// Responses:
//
//	       200: OK
//			  400: Bad Request
//			  500: Internal Server Error

// Refresh handles the refresh request, and refreshes the validity period of the token
func (apiContext *APIContext) Refresh(w http.ResponseWriter, r *http.Request) {
	status, _, claims := checkLogin(r)
	if status {
		// We ensure that a new token is not issued until enough time has elapsed
		// In this case, a new token will only be issued if the old token is within
		// 30 seconds of expiry. Otherwise, return a bad request status
		if claims.Payload.ExpiresAt.Sub(time.Now()) > 30*time.Minute {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// Now, create a new token for the current use, with a renewed expiration time
		expirationTime := time.Now().Add(60 * time.Minute)
		claims.Payload.ExpiresAt = jwt.NewNumericDate(expirationTime)
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		tokenstring, err := token.SignedString(hs)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// Set the new token as the users `token` cookie
		http.SetCookie(w, &http.Cookie{
			Name:    cookieName,
			Value:   tokenstring,
			Expires: expirationTime,
		})
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}
