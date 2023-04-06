package application

import (
	"errors"
	"testing"

	"github.com/serdarkalayci/carpool/api/domain"
	"github.com/stretchr/testify/assert"
)

type mockGeographyRepository struct{}

var (
	getCountriesFunc    func() ([]domain.Country, error)
	getCountryFunc      func(ID string) (domain.Country, error)
	checkBallotCityFunc func(countryID string, cityName string) (bool, error)
)

func (mgr mockGeographyRepository) GetCountries() ([]domain.Country, error) {
	return getCountriesFunc()
}

func (mgr mockGeographyRepository) GetCountry(ID string) (domain.Country, error) {
	return getCountryFunc(ID)
}

func (mgr mockGeographyRepository) CheckBallotCity(countryID string, cityName string) (bool, error) {
	return checkBallotCityFunc(countryID, cityName)
}

func TestNewGeographyService(t *testing.T) {
	mc := &MockContext{}
	mc.SetRepositories(nil, nil, &mockGeographyRepository{}, nil, nil, nil)
	gs := NewGeographyService(mc)
	assert.NotNil(t, gs)
	gr := gs.dc.GetGeographyRepository()
	assert.NotNil(t, gr)
	assert.Implements(t, (*GeographyRepository)(nil), gr)
}

func TestGetCountries(t *testing.T) {
	mc := &MockContext{}
	mc.SetRepositories(nil, nil, &mockGeographyRepository{}, nil, nil, nil)
	gs := NewGeographyService(mc)
	getCountriesFunc = func() ([]domain.Country, error) {
		return []domain.Country{
			{
				ID:   "1",
				Name: "Turkey",
			},
			{
				ID:   "2",
				Name: "USA",
			},
		}, nil
	}
	countries, err := gs.GetCountries()
	assert.Nil(t, err)
	assert.Equal(t, 2, len(countries))
	assert.Equal(t, "Turkey", countries[0].Name)
	assert.Equal(t, "USA", countries[1].Name)
}

func TestGetCountry(t *testing.T) {
	mc := &MockContext{}
	mc.SetRepositories(nil, nil, &mockGeographyRepository{}, nil, nil, nil)
	gs := NewGeographyService(mc)
	// Check the case where getCountriesFunc returns an error
	getCountryFunc = func(ID string) (domain.Country, error) {
		return domain.Country{}, errors.New("error getting country")
	}
	country, err := gs.GetCountry("1")
	assert.Errorf(t, err, "error getting country")
	assert.Equal(t, "", country.Name)
	// Check the case where getCountriesFunc returns a country, and with a ballot city
	getCountryFunc = func(ID string) (domain.Country, error) {
		return domain.Country{
			ID:   ID,
			Name: "Turkey",
			Cities: []domain.City{
				{
					Name:   "Istanbul",
					Ballot: true,
				},
				{
					Name:   "Ankara",
					Ballot: false,
				},
			},
		}, nil
	}
	country, err = gs.GetCountry("1")
	assert.Nil(t, err)
	assert.Equal(t, "Turkey", country.Name)

}
