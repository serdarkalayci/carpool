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

// RequestRepository is the repository for requests
type RequestRepository struct {
	dbClient *mongo.Client
	dbName   string
}

func newRequestRepository(client *mongo.Client, databaseName string) RequestRepository {
	return RequestRepository{
		dbClient: client,
		dbName:   databaseName,
	}
}

// AddRequest adds a request to the database
func (rr RequestRepository) AddRequest(request domain.Request) error {
	collection := rr.dbClient.Database(rr.dbName).Collection(viper.GetString("RequestCollection"))
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	requestDAO := mappers.MapRequest2RequestDAO(&request)
	requestDAO.ID = primitive.NewObjectID()
	_, err := collection.InsertOne(ctx, requestDAO)
	if err != nil {
		log.Error().Err(err).Msgf("error inserting request: %v", requestDAO)
		return apperr.ErrRequestNotInserted{}
	}
	return nil
}

// GetRequests returns all requests for a country, and if supplied filters by origin and destination
func (rr RequestRepository) GetRequests(countryID string, origin string, destination string) (*[]domain.Request, error) {
	collection := rr.dbClient.Database(rr.dbName).Collection(viper.GetString("RequestCollection"))
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	countryObjID, err := primitive.ObjectIDFromHex(countryID)
	if err != nil {
		log.Error().Err(err).Msgf("invalid countryID: %s", countryID)
		return nil, apperr.ErrInvalidID{Name: "countryID", Value: countryID}
	}
	filter := bson.M{"countryid": countryObjID}
	if origin != "" {
		filter["origin"] = origin
	}
	if destination != "" {
		filter["destination"] = destination
	}
	projection := bson.M{"_id": 1, "origin": 1, "destination": 1, "dates": 1, "requestedseats": 1, "state": 1}
	opts := options.Find().SetProjection(projection)
	cursor, err := collection.Find(ctx, filter, opts)
	if err != nil {
		log.Error().Err(err).Msgf("error getting requests for country: %s", countryID)
		return nil, apperr.ErrRequestNotFound{}
	}
	defer cursor.Close(ctx)
	var requests []dao.RequestDAO
	if err = cursor.All(ctx, &requests); err != nil {
		log.Error().Err(err).Msgf("error getting requests with countryID: %s", countryID)
		return nil, apperr.ErrRequestNotFound{}
	}
	if requests == nil {
		return nil, apperr.ErrRequestNotFound{}
	}
	return mappers.MapRequesDAOs2Requests(requests), nil
}

// GetRequest returns a request by its ID
func (rr RequestRepository) GetRequest(requestID string) (*domain.Request, error) {
	collection := rr.dbClient.Database(rr.dbName).Collection(viper.GetString("RequestCollection"))
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	requestObjID, err := primitive.ObjectIDFromHex(requestID)
	if err != nil {
		log.Error().Err(err).Msgf("invalid requestID: %s", requestID)
		return nil, apperr.ErrInvalidID{Name: "requestID", Value: requestID}
	}
	filter := bson.M{"_id": requestObjID}
	var request dao.RequestDAO
	err = collection.FindOne(ctx, filter).Decode(&request)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, apperr.ErrRequestNotFound{}
		}
		log.Error().Err(err).Msgf("error getting request with ID: %s", requestID)
		return nil, apperr.ErrRequestNotFound{}
	}
	return mappers.MapRequestDAO2Request(&request), nil
}

// SetRequestStatus sets the status of a request
func (rr RequestRepository) SetRequestStatus(requestID string, state int) error {
	collection := rr.dbClient.Database(rr.dbName).Collection(viper.GetString("RequestCollection"))
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	requestObjID, err := primitive.ObjectIDFromHex(requestID)
	if err != nil {
		log.Error().Err(err).Msgf("invalid requestID: %s", requestID)
		return apperr.ErrInvalidID{Name: "requestID", Value: requestID}
	}
	update := bson.M{"$set": bson.M{"state": state}}
	result, err := collection.UpdateByID(ctx, requestObjID, update)
	if err != nil {
		log.Error().Err(err).Msgf("error updating request with ID: %s", requestID)
		return apperr.ErrRequestNotUpdated{}
	}
	if result.MatchedCount == 0 {
		return apperr.ErrRequestNotFound{}
	}
	return nil
}
