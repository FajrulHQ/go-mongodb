package config

import (
	"context"
	"log"
	"os"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongoDBConnection() (*mongo.Database, error) {
	godotenv.Load()

	// Set client options
	uri := os.Getenv("MONGODB_URI")
	clientOptions := options.Client().ApplyURI(uri)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal("⛒ Connection Failed to Database")
		log.Fatal(err)
	}
	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal("⛒ Connection Failed to Database")
		log.Fatal(err)
	}

	color.Green("⛁ Connected to Database")

	return client.Database("portfolio"), nil
}
