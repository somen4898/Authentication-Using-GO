package mongoDB

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

var client *mongo.Client
var primaryDatabase *mongo.Database

var UserCollection *mongo.Collection

func Connect() error {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	_client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return err
	}
	setClient(_client)
	return nil
}

func setClient(c *mongo.Client) {
	client = c
	primaryDatabase = c.Database("go-auth")
	UserCollection = primaryDatabase.Collection("user")
}

func GetClient() *mongo.Client {
	return client
}

func GetCollection(name string) *mongo.Collection {
	return primaryDatabase.Collection(name)
}
