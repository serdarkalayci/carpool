package application

import "fmt"

type ErrInvalidID struct {
	Name string
	ID   string
}

func (e ErrInvalidID) Error() string {
	return fmt.Sprintf("invalid %s ID: %s", e.Name, e.ID)
}
