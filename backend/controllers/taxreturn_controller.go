package controllers

import (
	"github.com/gin-gonic/gin"
)

// @Summary Get all tax returns
// @Description Retrieve all tax returns from the database
// @Tags taxreturns
// @Accept json
// @Produce json
// @Success 200 {array} models.TaxReturn
// @Failure 500 {object} gin.H
// @Router /taxreturns [get]
func GetTaxReturns(c *gin.Context) {
	// ...existing code to handle tax returns...
}
