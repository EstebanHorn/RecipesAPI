package database

import (
	"context"
	"fmt"
	//"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Connect establishes a connection to the MongoDB database and returns the client.
func Connect() (*mongo.Client, error) {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().
		ApplyURI("mongodb+srv://estebanhxrn:4o6H3xrdKrvop0RY@recipes1.ras5f.mongodb.net/?retryWrites=true&w=majority&appName=Recipes1").
		SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %v", err)
	}

	// Check the connection to MongoDB
	if err := client.Ping(context.TODO(), nil); err != nil {
		return nil, fmt.Errorf("failed to ping MongoDB: %v", err)
	}

	fmt.Println("Successfully connected to MongoDB!")
	return client, nil
}
