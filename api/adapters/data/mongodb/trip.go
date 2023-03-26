package mongodb

import (
	"context"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/serdarkalayci/carpool/api/adapters/data/mongodb/mappers"
	"github.com/serdarkalayci/carpool/api/domain"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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
		return err
	}
	trip.ID = result.InsertedID.(primitive.ObjectID).Hex()
	return nil
}

// func (tr TripRepository) GetTripsByCountryID(countryID string) ([]*domain.Trip, error) {
// 	collection := tr.dbClient.Database(tr.dbName).Collection(viper.GetString("TripsCollection"))
// 	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
// 	defer cancel()
// 	opts := options.Find().SetProjection(bson.D{{"cities", 0}})
// 	var tripsDAO []*dao.TripDAO
// 	cursor, err := collection.Find(ctx, bson.M{"countryid": countryID}, opts)
// 	if err != nil {
// 		log.Error().Err(err).Msgf("error getting trips with countryID: %s", countryID)
// 		return nil, err
// 	}
// 	if err = cursor.All(ctx, &tripsDAO); err != nil {
// 		log.Error().Err(err).Msgf("error getting trips with countryID: %s", countryID)
// 		return nil, err
// 	}
// 	return mappers.MapTripDAO2Trip(tripsDAO), nil
// }
