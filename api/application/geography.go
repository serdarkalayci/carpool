// Package application is the package that holds the application logic between database and communication layers
package application

import (
	"github.com/serdarkalayci/carpool/api/domain"
)

// GeographyRepository is the interface that wraps the basic GetCountries and GetCountry methods.
type GeographyRepository interface {
	GetCountries() ([]domain.Country, error)
	GetCountry(ID string) (domain.Country, error)
	CheckBallotCity(countryID string, cityName string) (bool, error)
}

// GeographyService is the struct that wraps the basic GetCountries and GetCountry methods.
type GeographyService struct {
	dc DataContextCarrier
}

// NewGeographyService creates a new GeographyService instance and sets its repository
func NewGeographyService(dc DataContextCarrier) GeographyService {
	return GeographyService{
		dc: dc,
	}
}

// GetCountries returns all countries
func (gs GeographyService) GetCountries() ([]domain.Country, error) {
	return gs.dc.GetGeographyRepository().GetCountries()
}

// GetCountry returns a country by ID
func (gs GeographyService) GetCountry(countryID string) (domain.Country, error) {
	country, err := gs.dc.GetGeographyRepository().GetCountry(countryID)
	if err != nil {
		return domain.Country{}, err
	}
	country.BallotCities = getCitiesWithBallot(country)
	return country, nil
}

func getCitiesWithBallot(country domain.Country) []domain.City {
	var citiesWithBallot []domain.City
	for _, city := range country.Cities {
		if city.Ballot {
			citiesWithBallot = append(citiesWithBallot, city)
		}
	}
	return citiesWithBallot
}
