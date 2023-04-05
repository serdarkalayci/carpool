// Package domain is the package that holds the very basic domain objects
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
