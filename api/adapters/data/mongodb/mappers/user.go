// Package mappers is the package that maps objects back and fort between dao and domain
package mappers

import (
	"github.com/serdarkalayci/carpool/api/adapters/data/mongodb/dao"
	"github.com/serdarkalayci/carpool/api/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// MapUser2NewUserDAO maps domain User to dao User
func MapUser2NewUserDAO(u domain.User) dao.UserDAO {
	userDAO := dao.UserDAO{}
	userDAO.ID = primitive.NewObjectID()
	userDAO.Name = u.Name
	userDAO.Password = u.Password
	userDAO.Email = u.Email
	userDAO.Phone = u.Phone
	return userDAO
}

// MapUserDAO2User maps dao User to domain User
func MapUserDAO2User(u dao.UserDAO) domain.User {
	user := domain.User{}
	user.ID = u.ID.Hex()
	user.Name = u.Name
	user.Password = u.Password
	user.Email = u.Email
	user.Phone = u.Phone
	user.Active = u.Active
	user.Admin = u.Admin
	return user
}
