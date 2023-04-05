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

// GeographyRepository represent a structure that will communicate to MongoDB to accomplish geography related transactions
type GeographyRepository struct {
	dbClient *mongo.Client
	dbName   string
}

func newGeographyRepository(client *mongo.Client, databaseName string) GeographyRepository {
	return GeographyRepository{
		dbClient: client,
		dbName:   databaseName,
	}
}

// GetCountries returns the list of countries shy of their cities
func (gr GeographyRepository) GetCountries() ([]domain.Country, error) {
	collection := gr.dbClient.Database(gr.dbName).Collection(viper.GetString("GeographyCollection"))
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	opts := options.Find().SetProjection(bson.D{{"cities", 0}})
	cursor, err := collection.Find(ctx, bson.M{}, opts)
	if err != nil {
		log.Error().Err(err).Msgf("error getting Countries")
		return nil, apperr.ErrCountriesNotFound{}
	}
	countries := make([]dao.CountryDAO, 0)
	defer cursor.Close(ctx)
	if err = cursor.All(ctx, &countries); err != nil {
		return nil, apperr.ErrCountriesNotFound{}
	}
	return mappers.MapCountryDAOs2Countries(countries), nil
}

// GetCountry returns one country with the given ID and its cities
func (gr GeographyRepository) GetCountry(ID string) (domain.Country, error) {
	collection := gr.dbClient.Database(gr.dbName).Collection(viper.GetString("GeographyCollection"))
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	objID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		log.Error().Err(err).Msgf("error parsing CountryID: %s", ID)
		return domain.Country{}, apperr.ErrInvalidID{Name: "CountryID", Value: ID}
	}
	var countryDAO dao.CountryDAO
	err = collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&countryDAO)
	if err != nil {
		log.Error().Err(err).Msgf("error getting country with CountryID: %s", ID)
		return domain.Country{}, apperr.ErrCountryNotFound{}
	}
	return mappers.MapCountryDAO2Country(countryDAO), nil
}

// CheckBallotCity checks if the given city is a ballot city
func (gr GeographyRepository) CheckBallotCity(countryID string, cityName string) (bool, error) {
	collection := gr.dbClient.Database(gr.dbName).Collection(viper.GetString("GeographyCollection"))
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	countryObjID, err := primitive.ObjectIDFromHex(countryID)
	if err != nil {
		log.Error().Err(err).Msgf("error parsing CountryID: %s", countryID)
		return false, apperr.ErrInvalidID{Name: "CountryID", Value: countryID}
	}
	count, err := collection.CountDocuments(ctx, bson.M{"_id": countryObjID, "cities": bson.M{"$elemMatch": bson.M{"name": cityName, "ballot": true}}})
	if err != nil {
		log.Error().Err(err).Msgf("error checking ballot city with CountryID: %s and CityName: %s", countryID, cityName)
		return false, apperr.ErrCountryNotFound{}
	}
	return count == 1, nil
}
