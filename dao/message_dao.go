package dao

import (
	"context"
	"github.com/jasonkayzk/web-chat/model"
	"github.com/jasonkayzk/web-chat/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func InsertMessage(message *model.Message) (string, error) {
	collection := util.MongoClient.Database("chat").Collection("chat_log")

	res, err := collection.InsertOne(context.TODO(), model.Message{
		UUID:        message.UUID,
		IP:          message.IP,
		ToUUID:      message.ToUUID,
		MessageType: message.MessageType,
		Username:    message.Username,
		Content:     message.Content,
		MessageTime: message.MessageTime,
		InsertTime:  time.Now().UnixNano(),
	})
	if err != nil {
		return "", err
	}

	return res.InsertedID.(primitive.ObjectID).Hex(), nil
}

func GetHistoryMessage(insertTime, messageCount int64) ([]model.MessageResponse, error) {
	collection := util.MongoClient.Database("chat").Collection("chat_log")

	filter := bson.M{
		"insert_time": bson.M{
			"$lt": insertTime - 500,
		},
	}

	option := options.Find()
	option.SetSort(bson.D{
		{"insert_time", -1},
	})
	option.SetLimit(messageCount)

	cursor, err := collection.Find(context.TODO(), filter, option)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err = cursor.Close(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()

	var res []model.MessageResponse
	if err := cursor.All(context.TODO(), &res); err != nil {
		return nil, err
	}

	return res, nil
}
