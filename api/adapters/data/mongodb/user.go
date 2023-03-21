package mongodb

import (
	"context"
	"fmt"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/serdarkalayci/carpool/api/adapters/data/mongodb/dao"
	"github.com/serdarkalayci/carpool/api/adapters/data/mongodb/mappers"
	"github.com/serdarkalayci/carpool/api/domain"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// UserRepository represent a structure that will communicate to MongoDB to accomplish user related transactions
type UserRepository struct {
	dbClient *mongo.Client
	dbName   string
}

func newUserRepository(client *mongo.Client, databaseName string) UserRepository {
	return UserRepository{
		dbClient: client,
		dbName:   databaseName,
	}
}

// GetUser returns one user with the given ID if it exists in the array, returns not found error otherwise
func (ur UserRepository) GetUser(ID string) (domain.User, error) {
	collection := ur.dbClient.Database(ur.dbName).Collection(viper.GetString("UsersCollection")) ///ToDo: Change static string to configuration value
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	objID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		log.Error().Err(err).Msgf("error parsing UserID: %s", ID)
		return domain.User{}, err
	}
	var userDAO dao.UserDAO
	err = collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&userDAO)
	if err != nil {
		log.Error().Err(err).Msgf("error getting user with UserID: %s", ID)
		return domain.User{}, err
	}
	return mappers.MapUserDAO2User(userDAO), nil
}

// AddUser adds a new user to the array in the memory
func (ur UserRepository) AddUser(u domain.User) error {
	collection := ur.dbClient.Database(ur.dbName).Collection(viper.GetString("UsersCollection")) ///ToDo: Change static string to configuration value
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	userDAO := mappers.MapUser2NewUserDAO(u)
	result, err := collection.InsertOne(ctx, userDAO)
	if err != nil {
		log.Error().Err(err).Msg("error while writing user")
		if mongo.IsDuplicateKeyError(err) {
			return &domain.DuplicateKeyError{}
		}
		return err
	}
	log.Info().Msgf("user written: %s", result.InsertedID)
	return nil
}

// CheckUser checks the username & password if it matches any user from the array
func (ur UserRepository) CheckUser(username string, password string) (domain.User, error) {
	collection := ur.dbClient.Database(ur.dbName).Collection(viper.GetString("UsersCollection")) ///ToDo: Change static string to configuration value
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	var userDAO dao.UserDAO
	err := collection.FindOne(ctx, bson.M{"username": username, "password": password}).Decode(&userDAO)
	if err != nil {
		log.Error().Err(err).Msgf("error getting user with username: %s", username)
		return domain.User{}, err
	}
	return mappers.MapUserDAO2User(userDAO), nil
}

// CheckUserName checks the username if it exists in the database
func (ur UserRepository) CheckUserName(username string) (bool, error) {
	collection := ur.dbClient.Database(ur.dbName).Collection(viper.GetString("UsersCollection")) ///ToDo: Change static string to configuration value
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	var userDAO dao.UserDAO
	err := collection.FindOne(ctx, bson.M{"username": username}).Decode(&userDAO)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return true, nil
		}
		log.Error().Err(err).Msgf("error getting user with username: %s", username)
		return false, err
	}
	return false, nil
}

// UpdateUser updates an existing user on the user array
func (ur UserRepository) UpdateUser(u domain.User) error {
	return fmt.Errorf("not impelemented")
}

// DeleteUser deletes a user from the user array
func (ur UserRepository) DeleteUser(u domain.User) error {
	return fmt.Errorf("not impelemented")
}
