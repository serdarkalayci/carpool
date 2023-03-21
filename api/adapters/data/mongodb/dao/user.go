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
	Active   bool               `bson:"active"`
}
