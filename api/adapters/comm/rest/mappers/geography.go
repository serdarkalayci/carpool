// Package mappers is the package that maps objects back and fort between dto and domain
package mappers

import (
	"github.com/serdarkalayci/carpool/api/adapters/comm/rest/dto"
	"github.com/serdarkalayci/carpool/api/domain"
)

// MapCountry2CountryDTO maps domain.Country to dto.CountryDTO
func MapCountry2CountryDTO(country domain.Country) dto.CountryDTO {
	return dto.CountryDTO{
		ID:           country.ID,
		Name:         country.Name,
		Cities:       MapCities2CitiesDTO(country.Cities),
		BallotCities: MapCities2CitiesDTO(country.BallotCities),
	}
}

// MapCities2CitiesDTO maps []domain.City to []dto.CityDTO
func MapCities2CitiesDTO(cities []domain.City) []dto.CityDTO {
	citiesDTO := make([]dto.CityDTO, 0)
	for _, city := range cities {
		citiesDTO = append(citiesDTO, MapCity2CityDTO(city))
	}
	return citiesDTO
}

// MapCity2CityDTO maps domain.City to dto.CityDTO
func MapCity2CityDTO(city domain.City) dto.CityDTO {
	return dto.CityDTO{
		Name: city.Name,
	}
}
