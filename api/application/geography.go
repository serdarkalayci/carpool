package application

import (
	"github.com/serdarkalayci/carpool/api/domain"
)

type GeographyRepository interface {
	GetCountries() ([]domain.Country, error)
	GetCountry(ID string) (domain.Country, error)
}

type GeographyService struct {
	geographyRepository GeographyRepository
}

// NewGeographyService creates a new GeographyService instance and sets its repository
func NewGeographyService(gr GeographyRepository) GeographyService {
	if gr == nil {
		panic("missing geographyRepository")
	}
	return GeographyService{
		geographyRepository: gr,
	}
}

func (gs GeographyService) GetCountries() ([]domain.Country, error) {
	return gs.geographyRepository.GetCountries()
}

func (gs GeographyService) GetCountry(countryID string) (domain.Country, error) {
	country, err := gs.geographyRepository.GetCountry(countryID)
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
