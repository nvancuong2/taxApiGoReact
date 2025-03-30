package main

import (
	"log"

	"github.com/nvancuong2/taxApiGoReact/backend/database"
	"github.com/nvancuong2/taxApiGoReact/backend/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Set Gin to release mode for production
	gin.SetMode(gin.ReleaseMode)

	// Connect to MongoDB
	database.ConnectToMongoDB()

	router := gin.Default()

	// Set trusted proxies (e.g., allow only localhost or specific IPs)
	err := router.SetTrustedProxies([]string{"127.0.0.1"})
	if err != nil {
		log.Fatalf("Failed to set trusted proxies: %v", err)
	}

	routes.SetupRoutes()
	log.Println("Server running on http://localhost:8080")
	log.Fatal(router.Run(":8080"))
}
