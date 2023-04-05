// Package errors is the package that holds the custom application errors
package errors

// ErrInvalidDestination is returned when the destination is not a ballot city.
type ErrInvalidDestination struct{}

func (e ErrInvalidDestination) Error() string {
	return "destination is not a ballot city"
}

// ErrCountriesNotFound is returned when the are no countries to be returned.
type ErrCountriesNotFound struct{}

func (e ErrCountriesNotFound) Error() string {
	return "countries not found"
}

// ErrCountryNotFound is returned when the country is not found.
type ErrCountryNotFound struct{}

func (e ErrCountryNotFound) Error() string {
	return "country not found"
}
