package mongodb

import (
	"context"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/serdarkalayci/carpool/api/adapters/data/mongodb/dao"
	"github.com/serdarkalayci/carpool/api/adapters/data/mongodb/mappers"
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

func (tr TripRepository) GetTrips(countryID string, origin, destination string) ([]*domain.Trip, error) {
	collection := tr.dbClient.Database(tr.dbName).Collection(viper.GetString("TripsCollection"))
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	opts := options.Find().SetProjection(bson.M{"origin": 1, "destination": 1, "tripdate": 1, "availableseats": 1})
	var tripsDAO []*dao.TripDAO
	objID, err := primitive.ObjectIDFromHex(countryID)
	if err != nil {
		log.Error().Err(err).Msgf("error parsing countryID: %s", countryID)
		return nil, err
	}
	filter := bson.M{"countryid": objID}
	if origin != "" {
		filter["origin"] = origin
	}
	if destination != "" {
		filter["destination"] = destination
	}
	cursor, err := collection.Find(ctx, filter, opts)
	if err != nil {
		log.Error().Err(err).Msgf("error getting trips with countryID: %s", countryID)
		return nil, err
	}
	if err = cursor.All(ctx, &tripsDAO); err != nil {
		log.Error().Err(err).Msgf("error getting trips with countryID: %s", countryID)
		return nil, err
	}
	return mappers.MapTripDAOs2Trips(tripsDAO), nil
}

func (tr TripRepository) GetTripByID(tripID string) (*domain.TripDetail, error) {
	collection := tr.dbClient.Database(tr.dbName).Collection(viper.GetString("TripDetailsView"))
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	objID, err := primitive.ObjectIDFromHex(tripID)
	if err != nil {
		log.Error().Err(err).Msgf("error parsing tripID: %s", tripID)
		return nil, err
	}
	filter := bson.M{"_id": objID}
	var tripDetailDAO dao.TripDetailDAO
	err = collection.FindOne(ctx, filter).Decode(&tripDetailDAO)
	if err != nil {
		log.Error().Err(err).Msgf("error getting trip with tripID: %s", tripID)
		return nil, err
	}
	return mappers.MapTripDetailDAO2TripDetail(&tripDetailDAO), nil
}
