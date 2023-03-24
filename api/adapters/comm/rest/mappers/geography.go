package mappers

import (
	"github.com/serdarkalayci/carpool/api/adapters/comm/rest/dto"
	"github.com/serdarkalayci/carpool/api/domain"
)

func MapCountry2CountryDTO(country domain.Country) dto.CountryDTO {
	return dto.CountryDTO{
		ID:           country.ID,
		Name:         country.Name,
		Cities:       MapCities2CitiesDTO(country.Cities),
		BallotCities: MapCities2CitiesDTO(country.BallotCities),
	}
}

func MapCities2CitiesDTO(cities []domain.City) []dto.CityDTO {
	citiesDTO := make([]dto.CityDTO, 0)
	for _, city := range cities {
		citiesDTO = append(citiesDTO, MapCity2CityDTO(city))
	}
	return citiesDTO
}

func MapCity2CityDTO(city domain.City) dto.CityDTO {
	return dto.CityDTO{
		Name: city.Name,
	}
}
