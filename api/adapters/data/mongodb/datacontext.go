package mongodb

import (
	"context"
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/serdarkalayci/carpool/api/application"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// NewDataContext returns a new mongoDB backed DataContext
func NewDataContext() application.DataContext {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// We try to get connectionstring value from the environment variables, if not found it falls back to local database
	connectionString := os.Getenv("ConnectionString")
	if connectionString == "" {
		connectionString = "mongodb://mongoadmin:secret@localhost:27017"
		log.Info().Msg("ConnectionString from Env not found, falling back to local DB")
	} else {
		dbUserName := os.Getenv("DbUserName")
		dbPassword := os.Getenv("DbPassword")
		if dbUserName != "" {
			connectionString = strings.Replace(connectionString, "{username}", dbUserName, -1)
		}
		if dbPassword != "" {
			connectionString = strings.Replace(connectionString, "{password}", dbPassword, -1)
		}
		log.Info().Msgf("ConnectionString from Env is used: '%s'", connectionString)
	}
	databaseName := os.Getenv("DatabaseName")
	if databaseName == "" {
		databaseName = "carpool"
		log.Info().Msg("DatabaseName from Env not found, falling back to default")
	} else {
		log.Info().Msgf("DatabaseName from Env is used: '%s'", databaseName)
	}
	client, err := mongo.NewClient(options.Client().ApplyURI(connectionString))
	err = client.Connect(ctx)
	if err != nil {
		log.Error().Err(err).Msgf("An error occured while connecting to tha database")
	} else {
		// Check the connection
		err = client.Ping(ctx, nil)

		if err != nil {
			log.Error().Err(err).Msg("An error occured while connecting to tha database")
		}
		log.Info().Msg("Connected to MongoDB!")
	}
	dataContext := application.DataContext{}
	dataContext.UserRepository = newUserRepository(client, databaseName)
	dataContext.HealthRepository = newHealthRepository(client, databaseName)
	dataContext.GeographyRepository = newGeographyRepository(client, databaseName)
	dataContext.TripRepository = newTripRepository(client, databaseName)
	dataContext.ConversationRepository = newConversationRepository(client, databaseName)
	dataContext.RequestRepository = newRequestRepository(client, databaseName)
	return dataContext
}
