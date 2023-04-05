// Package dao is the package that holds the database access objects
package dao

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// UserDAO represents the struct of User type to be stored in mongoDB
type UserDAO struct {
	ID       primitive.ObjectID `bson:"_id"`
	Name     string             `bson:"name"`
	Password string             `bson:"password"`
	Email    string             `bson:"email"`
	Phone    string             `bson:"phone"`
	Active   bool               `bson:"active"`
	Admin    bool               `bson:"admin"`
}

// ConfirmationDAO represents the struct of Confirmation code and UserID to be stored in mongoDB
type ConfirmationDAO struct {
	UserID       primitive.ObjectID `bson:"userid"`
	Code         string             `bson:"code"`
	ValidityDate primitive.DateTime `bson:"validitydate"`
}
