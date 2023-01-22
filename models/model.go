package models

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Model struct {
	ID string `json:"id"`
}

var client *mongo.Client
var ctx context.Context

// Initialize connection with MongoDB database
func Setup() {
	var err error
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file")
	}

	dburl, exists := os.LookupEnv("MONGO_xURI")
	fmt.Println("dburi", dburl, exists)

	if !exists {
		log.Fatal("MongoDB URI not present")
	}

	client, err = mongo.Connect(ctx, options.Client().ApplyURI(dburl))
	if err != nil {
		log.Fatalln(err)
	}

	if err = client.Ping(ctx, nil); err != nil {
		log.Fatalln("err")
	} else {
		log.Println("Connected to Database")
	}

	_, err = client.Database("final-project").Collection("users").Indexes().CreateOne(
		ctx, mongo.IndexModel{Keys: bson.M{"username": 1}, Options: options.Index().SetUnique(true)})
	if err != nil {
		log.Println(err)
	} else {
		log.Println("created index on username")
	}

	// defer func() {
	// 	if err = client.Disconnect(ctx); err != nil {
	// 		log.Fatalln(err)
	// 	}
	// 	fmt.Println("disconnecting db")
	// }()

}
