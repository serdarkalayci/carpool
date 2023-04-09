// Package errors is the package that holds the custom application errors
package errors

// ErrUserNotFound is returned when a user is not found in the database.
type ErrUserNotFound struct{}

func (e ErrUserNotFound) Error() string {
	return "user not found"
}

// ErrWrongCredentials is returned when a user tries to login with wrong credentials.
type ErrWrongCredentials struct{}

func (e ErrWrongCredentials) Error() string {
	return "wrong credentials"
}

// ErrPasswordNotStrong is returned when a user tries to register with a password that is not strong enough.
type ErrPasswordNotStrong struct {
	Reason string
}

func (e ErrPasswordNotStrong) Error() string {
	return "password is not strong enough. " + e.Reason
}

// DuplicateKeyError is returned when a user tries to register with an email or phone number that already exists.
type DuplicateKeyError struct{}

func (d DuplicateKeyError) Error() string {
	return "email and/or phone number already exists"
}

// ConfirmationCodeError is returned when a user tries to confirm their account with an invalid code.
type ConfirmationCodeError struct{}

func (c ConfirmationCodeError) Error() string {
	return "confirmation code not found, does not match or expired"
}

// ErrUserNotInserted is returned when a user is not inserted into the database.
type ErrUserNotInserted struct{}

func (e ErrUserNotInserted) Error() string {
	return "user not inserted"
}

// ErrCodeNotInserted is returned when a confirmation code is not inserted into the database.
type ErrCodeNotInserted struct{}

func (e ErrCodeNotInserted) Error() string {
	return "confirmation code not inserted"
}
