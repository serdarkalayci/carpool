// Package dto is the package that defines types for data transfer
package dto

// CountryDTO is the response object for a country
type CountryDTO struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	Cities       []CityDTO `json:"cities"`
	BallotCities []CityDTO `json:"ballotCities"`
}

// CityDTO is the response object for a city
type CityDTO struct {
	Name string `json:"name"`
}
