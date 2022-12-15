package model

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Person struct {
	Name string       `bson:"name"` 
	Email int         `bson:"email"`
	Password string   `bson:"password"`
	Created_at string `bson:"created_at"`
	Updated_at string `bson:"updated_at"`
}

func (a Person) PostPerson() {
	client, err := Connect()
	coll := client.Database("shoe-resell").Collection("tPerson")

	// example
	newPerson := Person{Name: "test", Email: 20, Password: "asd23", Created_at: "0318894561", Updated_at: "114314" }

	result, err := coll.InsertOne(context.TODO(), newPerson)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Document inserted with ID: %s\n", result.InsertedID)
}