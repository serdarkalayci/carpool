// Package mongodb is the package that holds the database logic for mongodb database
package mongodb

import (
	"context"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/serdarkalayci/carpool/api/adapters/data/mongodb/dao"
	"github.com/serdarkalayci/carpool/api/adapters/data/mongodb/mappers"
	apperr "github.com/serdarkalayci/carpool/api/application/errors"
	"github.com/serdarkalayci/carpool/api/domain"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// TripRepository represent a structure that will communicate to MongoDB to accomplish user related transactions
type TripRepository struct {
	dbClient *mongo.Client
	dbName   string
}

func newTripRepository(client *mongo.Client, databaseName string) TripRepository {
	return TripRepository{
		dbClient: client,
		dbName:   databaseName,
	}
}

// AddTrip adds a new trip to the database
func (tr TripRepository) AddTrip(trip *domain.Trip) error {
	collection := tr.dbClient.Database(tr.dbName).Collection(viper.GetString("TripsCollection"))
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	tripDAO, err := mappers.MapTrip2TripDAO(*trip)
	if err != nil {
		log.Error().Err(err).Msgf("error mapping trip: %v", trip)
		return err
	}
	result, err := collection.InsertOne(ctx, tripDAO)
	if err != nil {
		log.Error().Err(err).Msgf("error inserting trip: %v", trip)
		return apperr.ErrTripNotInserted{}
	}
	trip.ID = result.InsertedID.(primitive.ObjectID).Hex()
	return nil
}

// GetTrips returns trips from the database
func (tr TripRepository) GetTrips(countryID string, origin, destination string) ([]*domain.Trip, error) {
	collection := tr.dbClient.Database(tr.dbName).Collection(viper.GetString("TripsCollection"))
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	opts := options.Find().SetProjection(bson.M{"origin": 1, "destination": 1, "tripdate": 1, "availableseats": 1})
	var tripsDAO []*dao.TripDAO
	objID, err := primitive.ObjectIDFromHex(countryID)
	if err != nil {
		log.Error().Err(err).Msgf("error parsing countryID: %s", countryID)
		return nil, apperr.ErrInvalidID{Name: "CountryID", Value: countryID}
	}
	today, _ := time.Parse("2006-01-02", time.Now().Format("2006-01-02"))
	filter := bson.M{"countryid": objID, "tripdate": bson.M{"$gte": today}}
	if origin != "" {
		filter["origin"] = origin
	}
	if destination != "" {
		filter["destination"] = destination
	}
	cursor, err := collection.Find(ctx, filter, opts)
	if err != nil {
		log.Error().Err(err).Msgf("error getting trips with countryID: %s", countryID)
		return nil, apperr.ErrTripNotFound{}
	}
	if err = cursor.All(ctx, &tripsDAO); err != nil {
		log.Error().Err(err).Msgf("error getting trips with countryID: %s", countryID)
		return nil, apperr.ErrTripNotFound{}
	}
	return mappers.MapTripDAOs2Trips(tripsDAO), nil
}

// GetTripByID returns a trip from the database based on it id
func (tr TripRepository) GetTripByID(tripID string) (*domain.TripDetail, error) {
	collection := tr.dbClient.Database(tr.dbName).Collection(viper.GetString("TripDetailsView"))
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	objID, err := primitive.ObjectIDFromHex(tripID)
	if err != nil {
		log.Error().Err(err).Msgf("error parsing tripID: %s", tripID)
		return nil, apperr.ErrInvalidID{Name: "TripID", Value: tripID}
	}
	filter := bson.M{"_id": objID}
	var tripDetailDAO dao.TripDetailDAO
	err = collection.FindOne(ctx, filter).Decode(&tripDetailDAO)
	if err != nil {
		log.Error().Err(err).Msgf("error getting trip with tripID: %s", tripID)
		return nil, apperr.ErrTripNotFound{}
	}
	return mappers.MapTripDetailDAO2TripDetail(&tripDetailDAO), nil
}

// GetConversationID returns the conversation id of a trip and a user as requester
func (tr TripRepository) GetConversationID(tripID string, userID string) (string, error) {
	collection := tr.dbClient.Database(tr.dbName).Collection(viper.GetString("ConversationsCollection"))
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	tripObjID, err := primitive.ObjectIDFromHex(tripID)
	if err != nil {
		log.Error().Err(err).Msgf("error parsing tripID: %s", tripID)
		return "", apperr.ErrInvalidID{Name: "TripID", Value: tripID}
	}
	userObjID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		log.Error().Err(err).Msgf("error parsing userID: %s", userID)
		return "", apperr.ErrInvalidID{Name: "UserID", Value: userID}
	}
	filter := bson.M{"tripid": tripObjID, "requesterid": userObjID}
	projection := bson.M{"_id": 1}
	opts := options.FindOne().SetProjection(projection)
	conversationDAO := dao.ConversationDAO{}
	err = collection.FindOne(ctx, filter, opts).Decode(&conversationDAO)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return "", nil
		}
		log.Error().Err(err).Msgf("error getting conversation with tripID: %s", tripID)
		return "", err
	}
	return conversationDAO.ID.Hex(), nil
}

// GetTripCapacity returns the capacity of a trip
func (tr TripRepository) GetTripCapacity(tripID string) (int, error) {
	collection := tr.dbClient.Database(tr.dbName).Collection(viper.GetString("TripsCollection"))
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	objID, err := primitive.ObjectIDFromHex(tripID)
	if err != nil {
		log.Error().Err(err).Msgf("error parsing tripID: %s", tripID)
		return 0, err
	}
	filter := bson.M{"_id": objID}
	projection := bson.M{"availableseats": 1}
	opts := options.FindOne().SetProjection(projection)
	tripDAO := dao.TripDAO{}
	err = collection.FindOne(ctx, filter, opts).Decode(&tripDAO)
	if err != nil {
		log.Error().Err(err).Msgf("error getting trip with tripID: %s", tripID)
		return 0, err
	}
	return tripDAO.AvailableSeats, nil
}

// SetTripCapacity updates the capacity of a trip
func (tr TripRepository) SetTripCapacity(tripID string, capacity int) error {
	collection := tr.dbClient.Database(tr.dbName).Collection(viper.GetString("TripsCollection"))
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	objID, err := primitive.ObjectIDFromHex(tripID)
	if err != nil {
		log.Error().Err(err).Msgf("error parsing tripID: %s", tripID)
		return err
	}
	filter := bson.M{"_id": objID}
	update := bson.M{"$inc": bson.M{"availableseats": capacity}}
	_, err = collection.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Error().Err(err).Msgf("error updating trip with tripID: %s", tripID)
		return err
	}
	return nil
}
