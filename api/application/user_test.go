package application

import "github.com/serdarkalayci/carpool/api/domain"

type mockUserRepository struct{}

var (
	getUserFunc               func(ID string) (domain.User, error)
	checkUserFunc             func(email string) (domain.User, error)
	addUserFunc               func(u domain.User) (string, error)
	addConfirmationCodeFunc   func(userID string, confirmationCode string) error
	checkConfirmationCodeFunc func(userID string, confirmationCode string) error
	activateUserFunc          func(userID string) error
)

// GetUser gets the user with the given ID
func (m mockUserRepository) GetUser(ID string) (domain.User, error) {
	return getUserFunc(ID)
}

// CheckUser checks if the user with the given email exists
func (m mockUserRepository) CheckUser(email string) (domain.User, error) {
	return checkUserFunc(email)
}

// AddUser adds a new user to the database
func (m mockUserRepository) AddUser(u domain.User) (string, error) {
	return addUserFunc(u)
}

// AddConfirmationCode adds a confirmation code to the user with the given ID
func (m mockUserRepository) AddConfirmationCode(userID string, confirmationCode string) error {
	return addConfirmationCodeFunc(userID, confirmationCode)
}

// CheckConfirmationCode checks if the confirmation code is correct for the user with the given ID
func (m mockUserRepository) CheckConfirmationCode(userID string, confirmationCode string) error {
	return checkConfirmationCodeFunc(userID, confirmationCode)
}

// ActivateUser activates the user with the given ID
func (m mockUserRepository) ActivateUser(userID string) error {
	return activateUserFunc(userID)
}
