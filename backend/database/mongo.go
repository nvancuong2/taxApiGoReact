package database

import (
	"context"
	"log"
	"time"

	"github.com/joho/godotenv"
	"github.com/nvancuong2/taxApiGoReact/backend/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client

// ConnectToMongoDB establishes a connection to MongoDB
func ConnectToMongoDB() {
	// Explicitly load the .env file
	err := godotenv.Load(".env")

	println("Loading .env file...", err)
	if err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	// Get MongoDB URI from environment variables
	mongoURI := config.GetMongoURI()

	println("Mongo URI: ", mongoURI)

	if mongoURI == "" {
		log.Fatal("MONGO_URI environment variable is not set")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	// Ping the database to verify the connection
	if err := client.Ping(ctx, nil); err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	}

	log.Println("Connected to MongoDB!")
	MongoClient = client
}
