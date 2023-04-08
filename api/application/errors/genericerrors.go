// Package errors is the package that holds the custom application errors
package errors

import "fmt"

// ErrInvalidID is an error that is returned when an ID is invalid, eg. not a valid UUID or a bson.ObjectId.
type ErrInvalidID struct {
	Name  string
	Value string
}

func (e ErrInvalidID) Error() string {
	return fmt.Sprintf("invalid %s ID: %s", e.Name, e.Value)
}

// ErrMissingRepository is an error that is returned when a required respository is not set.
type ErrMissingRepository struct{}

func (e ErrMissingRepository) Error() string {
	return "an internal problem occurred"
}
