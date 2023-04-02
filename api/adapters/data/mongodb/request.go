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
	"go.mongodb.org/mongo-driver/mongo"
)

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

func (rr RequestRepository) AddRequest(request domain.Request) error {
	collection := rr.dbClient.Database(rr.dbName).Collection(viper.GetString("RequestsCollection"))
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	requestDAO := mappers.MapRequest2RequestDAO(&request)
	_, err := collection.InsertOne(ctx, requestDAO)
	if err != nil {
		log.Error().Err(err).Msgf("error inserting request: %v", requestDAO)
		return err
	}
	return nil
}

func (rr RequestRepository) GetRequests(countryID string) (*[]domain.Request, error) {
	collection := rr.dbClient.Database(rr.dbName).Collection(viper.GetString("RequestsCollection"))
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	cursor, err := collection.Find(ctx, bson.M{"countryID": countryID})
	if err != nil {
		log.Error().Err(err).Msgf("error getting requests for country: %s", countryID)
		return nil, err
	}
	defer cursor.Close(ctx)
	var requests []dao.RequestDAO
	if err = cursor.All(ctx, &requests); err != nil {
		log.Error().Err(err).Msgf("error getting requests with countryID: %s", countryID)
		return nil, err
	}
	return mappers.MapRequesDAOs2Requests(requests), nil
}
