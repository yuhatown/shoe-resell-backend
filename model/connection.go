package model

import (
	"context"
	// "fmt"

	// "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect() (*mongo.Client, error) {
	Mongo_URL := "mongodb://127.0.0.1:27017"
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(Mongo_URL))
	if err != nil {
		panic(err)
	}

	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil { 
			panic(err)
		}
	}()

	return client, err
}