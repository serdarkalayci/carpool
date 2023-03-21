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

func (d *DuplicateKeyError) Error() string {
	return "email and/or phone number already exists"
}
