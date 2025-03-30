package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client

func ConnectToMongoDB() (*mongo.Client, error) {
	// Remplacez cette URI par l'URI de votre MongoDB Atlas
	uri := "mongodb+srv://nvancuong:PvxBhYFBgD55r0km@cluster0.qz2fljz.mongodb.net/"

	clientOptions := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	log.Println("Connected to MongoDB!")
	return client, nil
}
