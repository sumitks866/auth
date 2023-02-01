package models

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserAddress struct {
	Country string
	State   string
	City    string
	PIN     int
}

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Firstname string             `json:"firstname" bson:"firstname" binding:"required"`
	Lastname  string             `json:"lastname" bson:"lastname"`
	Username  string             `json:"username" bson:"username" binding:"required"`
	Password  string             `json:"password" bson:"password" binding:"required"`
}

func GetUser(username string, password string) (*User, error) {
	var res User
	filter := bson.M{"username": username}

	err := client.Database("final-project").Collection("users").FindOne(ctx, filter).Decode(&res)
	return &res, err
}

// Adds an user to database, returns ObjectId on success else returns error
func AddUser(user User) (*mongo.InsertOneResult, error) {

	usersColl := client.Database("final-project").Collection("users")
	res, err := usersColl.InsertOne(ctx, user)
	return res, err
}
