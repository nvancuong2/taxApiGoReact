package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nvancuong2/taxApiGoReact/backend/controllers"
)

func SetupRoutes(router *gin.Engine) {
	// Define the /taxreturns route
	router.GET("/taxreturns", controllers.GetTaxReturns)
}
