package mongodb

import (
	"context"
	"errors"
	"strconv"
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

// ConversationRepository represent a structure that will communicate to MongoDB to accomplish user related transactions
type ConversationRepository struct {
	dbClient *mongo.Client
	dbName   string
}

func newConversationRepository(client *mongo.Client, databaseName string) ConversationRepository {
	return ConversationRepository{
		dbClient: client,
		dbName:   databaseName,
	}
}

func (cr ConversationRepository) CheckConversation(tripID string, userID string) (convoID string, err error) {
	collection := cr.dbClient.Database(cr.dbName).Collection(viper.GetString("ConversationsCollection"))
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	tripObjID, err := primitive.ObjectIDFromHex(tripID)
	if err != nil {
		log.Error().Err(err).Msgf("error parsing tripID: %s", tripID)
		return "", err
	}
	userObjID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		log.Error().Err(err).Msgf("error parsing userID: %s", userID)
		return "", err
	}
	filter := bson.M{"tripid": tripObjID, "requesterid": userObjID}
	var conversation dao.ConversationDAO
	err = collection.FindOne(ctx, filter).Decode(&conversation)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return "", nil
		}
		log.Error().Err(err).Msgf("error getting trip with tripID: %s", tripID)
		return "", err
	}
	return conversation.ID.Hex(), nil
}

func (cr ConversationRepository) CheckConversationOwnership(conversationID string, userID string) (bool, error) {
	collection := cr.dbClient.Database(cr.dbName).Collection(viper.GetString("ConversationsCollection"))
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	convObjID, err := primitive.ObjectIDFromHex(conversationID)
	if err != nil {
		log.Error().Err(err).Msgf("error parsing tripID: %s", conversationID)
		return false, err
	}
	userObjID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		log.Error().Err(err).Msgf("error parsing userID: %s", userID)
		return false, err
	}
	filter := bson.M{"_id": convObjID, "supplierid": userObjID}
	count, err := collection.CountDocuments(ctx, filter)
	if err != nil {
		log.Error().Err(err).Msgf("error getting conversation with conversation: %s", conversationID)
		return false, err
	}
	return count == 1, nil
}

func (cr ConversationRepository) InitiateConversation(conversation domain.Conversation) error {
	collection := cr.dbClient.Database(cr.dbName).Collection(viper.GetString("ConversationsCollection"))
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	conversationDAO := mappers.MapConversation2ConversationDAO(&conversation)
	_, err := collection.InsertOne(ctx, conversationDAO)
	if err != nil {
		log.Error().Err(err).Msgf("error inserting conversation: %v", conversationDAO)
		return err
	}
	return nil
}

func (cr ConversationRepository) AddMessage(conversationID string, message string, direction string) error {
	collection := cr.dbClient.Database(cr.dbName).Collection(viper.GetString("ConversationsCollection"))
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	convoObjID, err := primitive.ObjectIDFromHex(conversationID)
	if err != nil {
		log.Error().Err(err).Msgf("error parsing conversationID: %s", conversationID)
		return err
	}
	update := bson.M{"$push": bson.M{"messages": dao.MessageDAO{Date: primitive.NewDateTimeFromTime(time.Now()), Text: message, Direction: direction}}}
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

// GetConversation returns the conversation for a trip when the Requester is the user
func (cr ConversationRepository) GetConversation(tripID string, userID string) (*domain.Conversation, error) {
	collection := cr.dbClient.Database(cr.dbName).Collection(viper.GetString("ConversationsCollection"))
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	tripObjID, err := primitive.ObjectIDFromHex(tripID)
	if err != nil {
		log.Error().Err(err).Msgf("error parsing tripID: %s", tripID)
		return nil, err
	}
	userObjID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		log.Error().Err(err).Msgf("error parsing userID: %s", userID)
		return nil, err
	}
	filter := bson.M{"tripid": tripObjID, "requesterid": userObjID}
	var conversationDAO dao.ConversationDAO
	err = collection.FindOne(ctx, filter).Decode(&conversationDAO)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		log.Error().Err(err).Msgf("error getting trip with tripID: %s", tripID)
		return nil, err
	}
	return mappers.MapConversationDAO2Conversation(&conversationDAO), nil
}

// GetConversationByID returns the conversation for a trip when the supplier is the user
func (cr ConversationRepository) GetConversationByID(conversationID string) (*domain.Conversation, error) {
	collection := cr.dbClient.Database(cr.dbName).Collection(viper.GetString("ConversationsCollection"))
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	convObjID, err := primitive.ObjectIDFromHex(conversationID)
	if err != nil {
		log.Error().Err(err).Msgf("error parsing conversationID: %s", conversationID)
		return nil, err
	}
	var conversationDAO dao.ConversationDAO
	err = collection.FindOne(ctx, bson.M{"_id": convObjID}).Decode(&conversationDAO)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		log.Error().Err(err).Msgf("error getting conversation with conversationID: %s", conversationID)
		return nil, err
	}
	return mappers.MapConversationDAO2Conversation(&conversationDAO), nil
}

// GetConversations returns the conversation for a trip when the Supplier is the user, so it returns all conversations without their messages
func (cr ConversationRepository) GetConversations(tripID string) ([]domain.Conversation, error) {
	collection := cr.dbClient.Database(cr.dbName).Collection(viper.GetString("ConversationsCollection"))
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	tripObjID, err := primitive.ObjectIDFromHex(tripID)
	if err != nil {
		log.Error().Err(err).Msgf("error parsing tripID: %s", tripID)
		return nil, err
	}
	filter := bson.M{"tripid": tripObjID}
	var conversations []dao.ConversationDAO
	opts := options.Find().SetProjection(bson.M{"messages": 0})
	cursor, err := collection.Find(ctx, filter, opts)
	if err != nil {
		log.Error().Err(err).Msgf("error getting trip with tripID: %s", tripID)
		return nil, err
	}
	defer cursor.Close(ctx)
	if err = cursor.All(ctx, &conversations); err != nil {
		log.Error().Err(err).Msgf("error getting conversations with tripID: %s", tripID)
		return nil, err
	}
	return mappers.MapConversationDAOs2Conversations(conversations), nil
}

func (cr ConversationRepository) MarkConversationsRead(conversationID string, direction string) error {
	collection := cr.dbClient.Database(cr.dbName).Collection(viper.GetString("ConversationsCollection"))
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	convoObjID, err := primitive.ObjectIDFromHex(conversationID)
	if err != nil {
		log.Error().Err(err).Msgf("error parsing conversationID: %s", conversationID)
		return err
	}
	filter := bson.M{"_id": convoObjID, "messages.direction": direction}
	update := bson.M{"$set": bson.M{"messages.$.read": true}}
	_, err = collection.UpdateMany(ctx, filter, update)
	if err != nil {
		log.Error().Err(err).Msgf("error updating conversatopn: %v", conversationID)
		return err
	}
	return nil
}

func (cr ConversationRepository) UpdateApproval(conversationID string, supplierApprove string, requesterApprove string) error {
	collection := cr.dbClient.Database(cr.dbName).Collection(viper.GetString("ConversationsCollection"))
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	convoObjID, err := primitive.ObjectIDFromHex(conversationID)
	if err != nil {
		log.Error().Err(err).Msgf("error parsing conversationID: %s", conversationID)
		return err
	}
	filter := bson.M{"_id": convoObjID}
	upd := bson.M{}
	if supplierApprove != "" {
		upd["supplierapprove"], _ = strconv.ParseBool(supplierApprove)
	}
	if requesterApprove != "" {
		upd["requesterapprove"], _ = strconv.ParseBool(requesterApprove)
	}
	update := bson.M{"$set": upd}
	result, err := collection.UpdateOne(ctx, filter, update)
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
