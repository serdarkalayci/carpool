package domain

// User type defines a user of the domain
type User struct {
	ID       string
	Name     string
	Email    string
	Password string
	Phone    string
	Active   bool
	Admin    bool
}

type DuplicateKeyError struct{}
type ConfirmationCodeError struct{}
type UserNotFoundError struct{}

func (d *DuplicateKeyError) Error() string {
	return "email and/or phone number already exists"
}

func (c *ConfirmationCodeError) Error() string {
	return "confirmation code not found, does not match or expired"
}

func (u *UserNotFoundError) Error() string {
	return "user not found"
}
