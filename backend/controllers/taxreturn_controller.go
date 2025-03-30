package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nvancuong2/taxApiGoReact/backend/models"
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
	// Example data (replace with database query)
	taxReturns := []models.TaxReturn{
		{
			ID:             "1",
			BusinessNumber: "123456789",
			TaxYear:        2023,
			Revenue:        100000.00,
			Expenses:       50000.00,
			NetIncome:      50000.00,
			TaxableIncome:  45000.00,
			TaxPayable:     7500.00,
			Refund:         0.00,
			CreatedAt:      "2025-03-30T12:00:00Z",
			UpdatedAt:      "2025-03-30T12:00:00Z",
		},
	}

	// Return the tax returns as JSON
	c.JSON(http.StatusOK, taxReturns)
}
