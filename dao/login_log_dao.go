package dao

import (
	"context"
	"github.com/jasonkayzk/web-chat/model"
	"github.com/jasonkayzk/web-chat/util"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func InsertLoginLog(loginLog *model.LoginLog) (string, error) {
	collection := util.MongoClient.Database("chat").Collection("login_log")

	res, err := collection.InsertOne(context.TODO(), model.LoginLog{
		UUID:       loginLog.UUID,
		IP:         loginLog.IP,
		LogType:    loginLog.LogType,
		Username:   loginLog.Username,
		InsertTime: time.Now().UnixNano(),
	})
	if err != nil {
		return "", err
	}

	return res.InsertedID.(primitive.ObjectID).Hex(), nil
}
