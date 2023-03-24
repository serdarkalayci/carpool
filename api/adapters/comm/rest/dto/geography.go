package dto

type CountryDTO struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	Cities       []CityDTO `json:"cities"`
	BallotCities []CityDTO `json:"ballotCities"`
}
type CityDTO struct {
	Name string `json:"name"`
}
