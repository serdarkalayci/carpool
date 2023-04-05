// Package errors is the package that holds the custom application errors
package errors

// ErrTripNotFound is returned when a trip is not found.
type ErrTripNotFound struct{}

func (e ErrTripNotFound) Error() string {
	return "trip not found"
}

// ErrTripNotInserted is returned when a trip is not inserted.
type ErrTripNotInserted struct{}

func (e ErrTripNotInserted) Error() string {
	return "trip not inserted"
}
