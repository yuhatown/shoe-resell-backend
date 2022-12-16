package model

import (
	"context"
	"fmt"
	"time"

)

type User struct {
	Name string       `bson:"name"` 
	Email int         `bson:"email"`
	Password string   `bson:"password"`
	Created_at time.Time `bson:"created_at"`
	Updated_at *time.Time `bson:"updated_at"`
}

func (a *User) PostUserModel() {
	client, err := Connect()
	coll := client.Database("shoe-resell").Collection("tPerson")

	// example
	result, err := coll.InsertOne(context.TODO(), a)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Document inserted with ID: %s\n", result.InsertedID)
}