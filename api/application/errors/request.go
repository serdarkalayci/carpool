// Package errors is the package that holds the custom application errors
package errors

// ErrRequestNotFound is returned when a request is not found in the database.
type ErrRequestNotFound struct{}

func (e ErrRequestNotFound) Error() string {
	return "request not found"
}

// ErrRequestNotInserted is returned when a request cannot be inserted in the database.
type ErrRequestNotInserted struct{}

func (e ErrRequestNotInserted) Error() string {
	return "cannot add request"
}

// ErrRequestNotUpdated is returned when a request cannot be updated in the database.
type ErrRequestNotUpdated struct{}

func (e ErrRequestNotUpdated) Error() string {
	return "cannot update request"
}
