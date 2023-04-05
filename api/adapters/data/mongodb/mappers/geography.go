// Package mappers is the package that maps objects back and fort between dao and domain
package mappers

import (
	"github.com/serdarkalayci/carpool/api/adapters/data/mongodb/dao"
	"github.com/serdarkalayci/carpool/api/domain"
)

// MapCityDAO2City maps a CityDAO to a City
func MapCityDAO2City(cityDAO dao.CityDAO) domain.City {
	return domain.City{
		Name:   cityDAO.Name,
		Ballot: cityDAO.Ballot,
	}
}

// MapCountryDAO2Country maps a CountryDAO to a Country
func MapCountryDAO2Country(countryDAO dao.CountryDAO) domain.Country {
	cities := make([]domain.City, len(countryDAO.Cities))
	for i, cityDAO := range countryDAO.Cities {
		cities[i] = MapCityDAO2City(cityDAO)
	}
	return domain.Country{
		ID:     countryDAO.ID.Hex(),
		Name:   countryDAO.Name,
		Cities: cities,
	}
}

// MapCountryDAOs2Countries maps a slice of CountryDAO to a slice of Country
func MapCountryDAOs2Countries(countryDAOs []dao.CountryDAO) []domain.Country {
	countries := make([]domain.Country, len(countryDAOs))
	for i, countryDAO := range countryDAOs {
		countries[i] = MapCountryDAO2Country(countryDAO)
	}
	return countries
}
