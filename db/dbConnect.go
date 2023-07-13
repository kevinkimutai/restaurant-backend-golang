package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DbConnect establishes a connection to a MongoDB server
func DbConnect() (*mongo.Client, error) {
	// Set client options
	clientOptions := options.Client().ApplyURI("YOUR_ATLAS_CONNECTION_STRING")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	// Ping the MongoDB server to verify the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	log.Println("Connected to MongoDB!")

	return client, nil
}
