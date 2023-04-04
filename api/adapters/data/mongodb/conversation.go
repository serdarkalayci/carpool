package mongodb

import (
	"context"
	"strconv"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/serdarkalayci/carpool/api/adapters/data/mongodb/dao"
	"github.com/serdarkalayci/carpool/api/adapters/data/mongodb/mappers"
	"github.com/serdarkalayci/carpool/api/application"
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

// CheckConversation checks if a conversation exists for current user a user under a trip
func (cr ConversationRepository) CheckConversation(tripID string, userID string) (convoID string, err error) {
	collection := cr.dbClient.Database(cr.dbName).Collection(viper.GetString("ConversationsCollection"))
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	tripObjID, err := primitive.ObjectIDFromHex(tripID)
	if err != nil {
		log.Error().Err(err).Msgf("error parsing tripID: %s", tripID)
		return "", application.ErrInvalidID{Name: "tripID", Value: tripID}
	}
	userObjID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		log.Error().Err(err).Msgf("error parsing userID: %s", userID)
		return "", application.ErrInvalidID{Name: "userID", Value: userID}
	}
	filter := bson.M{"tripid": tripObjID, "requesterid": userObjID}
	var conversation dao.ConversationDAO
	err = collection.FindOne(ctx, filter).Decode(&conversation)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return "", nil
		}
		log.Error().Err(err).Msgf("error getting trip with tripID: %s", tripID)
		return "", application.ErrConversationNotFound{}
	}
	return conversation.ID.Hex(), nil
}

// CheckConversationOwnership checks if a conversation's supplier is the current user
func (cr ConversationRepository) CheckConversationOwnership(conversationID string, userID string) (bool, error) {
	collection := cr.dbClient.Database(cr.dbName).Collection(viper.GetString("ConversationsCollection"))
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	convObjID, err := primitive.ObjectIDFromHex(conversationID)
	if err != nil {
		log.Error().Err(err).Msgf("error parsing tripID: %s", conversationID)
		return false, application.ErrInvalidID{Name: "conversationID", Value: conversationID}
	}
	userObjID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		log.Error().Err(err).Msgf("error parsing userID: %s", userID)
		return false, application.ErrInvalidID{Name: "userID", Value: userID}
	}
	filter := bson.M{"_id": convObjID, "supplierid": userObjID}
	count, err := collection.CountDocuments(ctx, filter)
	if err != nil {
		log.Error().Err(err).Msgf("error getting conversation with conversation: %s", conversationID)
		return false, application.ErrConversationNotFound{}
	}
	return count == 1, nil
}

// InitiateConversation creates a new conversation for a trip and a user
func (cr ConversationRepository) InitiateConversation(conversation domain.Conversation) error {
	collection := cr.dbClient.Database(cr.dbName).Collection(viper.GetString("ConversationsCollection"))
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	conversationDAO := mappers.MapConversation2ConversationDAO(&conversation)
	_, err := collection.InsertOne(ctx, conversationDAO)
	if err != nil {
		log.Error().Err(err).Msgf("error inserting conversation: %v", conversationDAO)
		return application.ErrConversationNotInserted{}
	}
	return nil
}

// AddMessage adds a message to a conversation
func (cr ConversationRepository) AddMessage(conversationID string, message string, direction string) error {
	collection := cr.dbClient.Database(cr.dbName).Collection(viper.GetString("ConversationsCollection"))
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	convoObjID, err := primitive.ObjectIDFromHex(conversationID)
	if err != nil {
		log.Error().Err(err).Msgf("error parsing conversationID: %s", conversationID)
		return application.ErrInvalidID{Name: "conversationID", Value: conversationID}
	}
	update := bson.M{"$push": bson.M{"messages": dao.MessageDAO{Date: primitive.NewDateTimeFromTime(time.Now()), Text: message, Direction: direction}}}
	result, err := collection.UpdateByID(ctx, convoObjID, update)
	if err != nil {
		log.Error().Err(err).Msgf("error inserting message to conversation ID: %s", conversationID)
		return application.ErrMessageNotInserted{}
	}
	if result.MatchedCount == 0 {
		log.Error().Err(err).Msgf("conversation not found: %v", conversationID)
		return application.ErrConversationNotFound{}
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
		return nil, application.ErrInvalidID{Name: "tripID", Value: tripID}
	}
	userObjID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		log.Error().Err(err).Msgf("error parsing userID: %s", userID)
		return nil, application.ErrInvalidID{Name: "userID", Value: userID}
	}
	filter := bson.M{"tripid": tripObjID, "requesterid": userObjID}
	var conversationDAO dao.ConversationDAO
	err = collection.FindOne(ctx, filter).Decode(&conversationDAO)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		log.Error().Err(err).Msgf("error getting trip with tripID: %s", tripID)
		return nil, application.ErrConversationNotFound{}
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
		return nil, application.ErrInvalidID{Name: "conversationID", Value: conversationID}
	}
	var conversationDAO dao.ConversationDAO
	err = collection.FindOne(ctx, bson.M{"_id": convObjID}).Decode(&conversationDAO)
	if err != nil {
		log.Error().Err(err).Msgf("error getting conversation with conversationID: %s", conversationID)
		return nil, application.ErrConversationNotFound{}
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
		return nil, application.ErrInvalidID{Name: "tripID", Value: tripID}
	}
	filter := bson.M{"tripid": tripObjID}
	var conversations []dao.ConversationDAO
	opts := options.Find().SetProjection(bson.M{"messages": 0})
	cursor, err := collection.Find(ctx, filter, opts)
	if err != nil {
		log.Error().Err(err).Msgf("error getting trip with tripID: %s", tripID)
		return nil, application.ErrTripNotFound{}
	}
	defer cursor.Close(ctx)
	if err = cursor.All(ctx, &conversations); err != nil {
		log.Error().Err(err).Msgf("error getting conversations with tripID: %s", tripID)
		return nil, application.ErrConversationNotFound{}
	}
	return mappers.MapConversationDAOs2Conversations(conversations), nil
}

// MarkConversationsRead marks all the messages of a conversation as read depending on the direction
func (cr ConversationRepository) MarkConversationsRead(conversationID string, direction string) error {
	collection := cr.dbClient.Database(cr.dbName).Collection(viper.GetString("ConversationsCollection"))
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	convoObjID, err := primitive.ObjectIDFromHex(conversationID)
	if err != nil {
		log.Error().Err(err).Msgf("error parsing conversationID: %s", conversationID)
		return application.ErrInvalidID{Name: "conversationID", Value: conversationID}
	}
	filter := bson.M{"_id": convoObjID, "messages.direction": direction}
	update := bson.M{"$set": bson.M{"messages.$.read": true}}
	_, err = collection.UpdateMany(ctx, filter, update)
	if err != nil {
		log.Error().Err(err).Msgf("error updating conversation: %v", conversationID)
		return application.ErrMessageNotUpdated{}
	}
	return nil
}

// UpdateApproval updates the approval of a conversation
func (cr ConversationRepository) UpdateApproval(conversationID string, supplierApprove string, requesterApprove string) error {
	collection := cr.dbClient.Database(cr.dbName).Collection(viper.GetString("ConversationsCollection"))
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	convoObjID, err := primitive.ObjectIDFromHex(conversationID)
	if err != nil {
		log.Error().Err(err).Msgf("error parsing conversationID: %s", conversationID)
		return application.ErrInvalidID{Name: "conversationID", Value: conversationID}
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
		log.Error().Err(err).Msgf("error updating conversation: %v", conversationID)
		return application.ErrConversationNotUpdated{}
	}
	if result.MatchedCount == 0 {
		log.Error().Err(err).Msgf("conversation not found: %v", conversationID)
		return application.ErrConversationNotFound{}
	}
	return nil
}
