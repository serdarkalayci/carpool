package application

import (
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
