// Package rest is responsible for rest communication layer
package rest

import "net/http"

// CorsHandler swagger:route OPTIONS /
//
// # Handler to respond to CORS preflight requests
//
// Responses:
//
//	200: OK
func CorsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, lang")
	w.Header().Set("Access-Control-Allow-Methods", "GET, PUT, POST, OPTIONS")
	w.WriteHeader(200)
}
