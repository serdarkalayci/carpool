// Package errors contains the generic errors for the REST adapter
package errors

import "fmt"

// ErrCannotParsePayload is the error returned when the payload cannot be parsed
type ErrCannotParsePayload struct{}

func (e *ErrCannotParsePayload) Error() string {
	return "cannot parse payload"
}

// ErrCannotReadPayload is the error returned when the payload cannot be read
type ErrCannotReadPayload struct{}

func (e ErrCannotReadPayload) Error() string {
	return "cannot read payload"
}

// ErrPayloadMissing is the error returned when the payload is missing
type ErrPayloadMissing struct{}

func (e ErrPayloadMissing) Error() string {
	return "payload missing"
}

// ErrInvalidDateFormat is the error returned when the date format is invalid
type ErrInvalidDateFormat struct {
	Date string
}

func (e ErrInvalidDateFormat) Error() string {
	return fmt.Sprintf("invalid date %s for the format: yyyy-MM-dd", e.Date)
}
