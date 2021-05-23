package util

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	mongoUrlTpl   = `mongodb://%s:%s@%s:%s/%s?authSource=admin&w=majority&readPreference=primary&retryWrites=true&ssl=false`
	mongoUsername = `admin`
	mongoPassword = `passwd`
	mongoHost     = `127.0.0.1`
	mongoPort     = `27017`
	mongoDBName   = `chat`
)

var MongoClient *mongo.Client

func InitMongo() {
	ctx := context.TODO()
	clientOptions := options.Client().ApplyURI(fmt.Sprintf(
		mongoUrlTpl,
		mongoUsername,
		mongoPassword,
		mongoHost,
		mongoPort,
		mongoDBName,
	))

	mongoClient, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		panic(err)
	}

	err = mongoClient.Ping(ctx, nil)
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to MongoDB!")
	MongoClient = mongoClient
}
