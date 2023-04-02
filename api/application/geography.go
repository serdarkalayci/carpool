package application

import (
	"github.com/serdarkalayci/carpool/api/domain"
)

type GeographyRepository interface {
	GetCountries() ([]domain.Country, error)
	GetCountry(ID string) (domain.Country, error)
	CheckBallotCity(countryID string, cityName string) (bool, error)
}

type GeographyService struct {
	dc DataContext
}

// NewGeographyService creates a new GeographyService instance and sets its repository
func NewGeographyService(dc DataContext) GeographyService {
	return GeographyService{
		dc: dc,
	}
}

func (gs GeographyService) GetCountries() ([]domain.Country, error) {
	return gs.dc.GeographyRepository.GetCountries()
}

func (gs GeographyService) GetCountry(countryID string) (domain.Country, error) {
	country, err := gs.dc.GeographyRepository.GetCountry(countryID)
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
