package model

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Person struct {
	Name string    `bson:"name"` 
	Age int        `bson:"age"`
	Pnum string    `bson:"pnum"`
}

func main() {
	Mongo_URL := "mongodb://127.0.0.1:27017"
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(Mongo_URL))
	if err != nil {
		panic(err)
	}

	coll := client.Database("shoe-resell").Collection("tPerson")

	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil { 
			panic(err)
		}
	}()

}