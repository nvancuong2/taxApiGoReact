package database

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client

//@author nvancuong2
//@date 2023-10-01
//@description This file contains the MongoDB connection logic for the backend of the tax API application. It uses the MongoDB Go driver to establish a connection to the database and handles environment variables using godotenv.

// ConnectToMongoDB establishes a connection to MongoDB
func ConnectToMongoDB() {
	// Explicitly load the .env file
	//err := godotenv.Load("e:/source/repose/taxApiGoReact/backend/.env")
	err := godotenv.Load("backend/.env")
	if err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	// Get MongoDB URI from environment variables
	mongoURI := os.Getenv("MONGO_URI")
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
