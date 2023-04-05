// Package rest is responsible for rest communication layer
package rest

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/serdarkalayci/carpool/api/adapters/comm/rest/dto"
	"github.com/serdarkalayci/carpool/api/adapters/comm/rest/mappers"
	"github.com/serdarkalayci/carpool/api/application"
)

// swagger:route GET /country/{countryid} Country GetCountry
// Return the country if found
// responses:
//	200: OK
//	404: errorResponse

// GetCountry gets a single country if found
func (apiContext *APIContext) GetCountry(rw http.ResponseWriter, r *http.Request) {
	// parse the Rating id from the url
	vars := mux.Vars(r)

	// convert the id into an integer and return
	countryID := vars["countryid"]
	geographyService := application.NewGeographyService(apiContext.dbContext)
	user, err := geographyService.GetCountry(countryID)
	if err == nil {
		respondWithJSON(rw, r, 200, mappers.MapCountry2CountryDTO(user))
		return
	}
	respondWithError(rw, r, 404, err.Error())
}

// swagger:route GET /country Country GetCountries
// Return the country list
// responses:
//	200: OK
//	404: errorResponse

// GetCountries gets country list
func (apiContext *APIContext) GetCountries(rw http.ResponseWriter, r *http.Request) {
	// span := createSpan("Carpool.GetCountries", r)
	// defer span.Finish()

	geographyService := application.NewGeographyService(apiContext.dbContext)
	countries, err := geographyService.GetCountries()
	if err != nil {
		respondWithError(rw, r, 500, err.Error())
	} else {
		countryDTOs := make([]dto.CountryDTO, 0)
		for _, p := range countries {
			pDTO := mappers.MapCountry2CountryDTO(p)
			countryDTOs = append(countryDTOs, pDTO)
		}
		respondWithJSON(rw, r, 200, countryDTOs)
	}

}
