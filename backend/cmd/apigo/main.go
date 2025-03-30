package main

import (
	"context"
	"log"
	"time"

	"github.com/nvancuong2/taxApiGoReact/backend/database"
	"github.com/nvancuong2/taxApiGoReact/backend/routes"

	"github.com/gin-gonic/gin"
	_ "github.com/nvancuong2/taxApiGoReact/backend/cmd/apigo/docs" // Import Swagger docs
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.mongodb.org/mongo-driver/bson"
)

// @title Tax API
// @version 1.0
// @description This is a sample API for tax management.
// @host localhost:8080
// @BasePath /
func main() {
	// Set Gin to release mode for production
	gin.SetMode(gin.ReleaseMode)

	// Connect to MongoDB
	database.ConnectToMongoDB()

	// Get the MongoDB collection
	collection := database.MongoClient.Database("taxdb").Collection("taxreturns")

	// Insert data
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := collection.InsertMany(ctx, []interface{}{
		bson.M{
			"business_number": "123456789",
			"tax_year":        2023,
			"revenue":         100000,
			"expenses":        50000,
			"net_income":      50000,
			"taxable_income":  45000,
			"tax_payable":     7500,
			"refund":          0,
			"created_at":      time.Now(),
			"updated_at":      time.Now(),
		},
		bson.M{
			"business_number": "987654321",
			"tax_year":        2022,
			"revenue":         200000,
			"expenses":        100000,
			"net_income":      100000,
			"taxable_income":  90000,
			"tax_payable":     15000,
			"refund":          0,
			"created_at":      time.Now(),
			"updated_at":      time.Now(),
		},
	})
	if err != nil {
		log.Fatalf("Failed to insert data: %v", err)
	}

	log.Println("Data inserted successfully!")

	router := gin.Default()

	// Set trusted proxies (e.g., allow only localhost or specific IPs)
	err = router.SetTrustedProxies([]string{"127.0.0.1"})
	if err != nil {
		log.Fatalf("Failed to set trusted proxies: %v", err)
	}

	// Swagger route
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	routes.SetupRoutes(router)
	log.Println("Server running on http://localhost:8080")
	log.Fatal(router.Run(":8080"))
}
