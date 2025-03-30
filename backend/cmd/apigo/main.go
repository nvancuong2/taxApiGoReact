package main

import (
	"context"
	"log"

	"github.com/nvancuong2/taxApiGoReact/backend/database" // Ensure this matches your module path
)

func main() {
	// Connect to MongoDB using the database package
	client, err := database.ConnectToMongoDB()
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			log.Fatalf("Failed to disconnect from MongoDB: %v", err)
		}
	}()

	log.Println("MongoDB connection established successfully!")
}
