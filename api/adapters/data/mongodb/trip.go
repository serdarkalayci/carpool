package mongodb

import (
	"context"
	"errors"
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

func (tr TripRepository) CheckConversation(tripID string, userID string) (convoID string, requesterID string, err error) {
	collection := tr.dbClient.Database(tr.dbName).Collection(viper.GetString("ConversationsCollection"))
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	tripObjID, err := primitive.ObjectIDFromHex(tripID)
	if err != nil {
		log.Error().Err(err).Msgf("error parsing tripID: %s", tripID)
		return "", "", err
	}
	userObjID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		log.Error().Err(err).Msgf("error parsing userID: %s", userID)
		return "", "", err
	}
	filter := bson.M{"tripid": tripObjID, "requesterid": userObjID}
	var conversation dao.Conversation
	err = collection.FindOne(ctx, filter).Decode(&conversation)
	if err != nil {
		log.Error().Err(err).Msgf("error getting trip with tripID: %s", tripID)
		return "", "", err
	}
	return conversation.ID.Hex(), conversation.RequesterID.Hex(), nil
}

func (tr TripRepository) InitiateConversation(tripID string, userID string, message string) error {
	collection := tr.dbClient.Database(tr.dbName).Collection(viper.GetString("ConversationsCollection"))
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	tripObjID, err := primitive.ObjectIDFromHex(tripID)
	if err != nil {
		log.Error().Err(err).Msgf("error parsing tripID: %s", tripID)
		return err
	}
	userObjID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		log.Error().Err(err).Msgf("error parsing userID: %s", userID)
		return err
	}
	conversationDAO := dao.Conversation{
		ID:          primitive.NewObjectID(),
		TripID:      tripObjID,
		RequesterID: userObjID,
		Messages:    []dao.Message{{Date: primitive.NewDateTimeFromTime(time.Now()), Text: message, Direction: "in"}},
	}
	_, err = collection.InsertOne(ctx, conversationDAO)
	if err != nil {
		log.Error().Err(err).Msgf("error inserting conversation: %v", conversationDAO)
		return err
	}
	return nil
}

func (tr TripRepository) AddMessage(conversationID string, message string, direction string) error {
	collection := tr.dbClient.Database(tr.dbName).Collection(viper.GetString("ConversationsCollection"))
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	convoObjID, err := primitive.ObjectIDFromHex(conversationID)
	if err != nil {
		log.Error().Err(err).Msgf("error parsing conversationID: %s", conversationID)
		return err
	}
	update := bson.M{"$push": bson.M{"messages": dao.Message{Date: primitive.NewDateTimeFromTime(time.Now()), Text: message, Direction: direction}}}
	result, err := collection.UpdateByID(ctx, convoObjID, update)
	if err != nil {
		log.Error().Err(err).Msgf("error updating conversatopn: %v", conversationID)
		return err
	}
	if result.MatchedCount == 0 {
		log.Error().Err(err).Msgf("conversation not found: %v", conversationID)
		return errors.New("conversation not found")
	}
	return nil
}
