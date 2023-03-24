package mappers

import (
	"github.com/serdarkalayci/carpool/api/adapters/data/mongodb/dao"
	"github.com/serdarkalayci/carpool/api/domain"
)

func MapCityDAO2City(cityDAO dao.CityDAO) domain.City {
	return domain.City{
		Name:   cityDAO.Name,
		Ballot: cityDAO.Ballot,
	}
}

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
