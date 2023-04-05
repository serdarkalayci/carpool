// Package dto is the package that defines types for data transfer
package dto

// AddUserRequest type defines a model for adding an user
type AddUserRequest struct {
	Name     string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Phone    string `json:"phone" validate:"required,e164"`
}

// UpdateUserRequest type defines a model for updating an user
type UpdateUserRequest struct {
	ID       string `json:"id" validate:"required"`
	Name     string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// ShortUserResponse type defines a model for returning a user shy of its personal details
type ShortUserResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// LongUserResponse type defines a model for returning a user shy including its personal details
type LongUserResponse struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

// LoginRequest type defines a model for getting an user's data for login operation
type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

// ConfirmUserRequest type defines a model for accepting a confirmation code for a user
type ConfirmUserRequest struct {
	Code string `json:"code" validate:"required"`
}
