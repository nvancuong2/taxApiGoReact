package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nvancuong2/taxApiGoReact/backend/database"
	"github.com/nvancuong2/taxApiGoReact/backend/models"
	"go.mongodb.org/mongo-driver/bson"
)

// TaxReturnController gère les opérations liées aux déclarations fiscales.
// @Summary Récupérer toutes les déclarations fiscales
// @Description Retourne la liste des déclarations fiscales stockées dans la base de données
// @Tags TaxReturn
// @Produce json
// @Success 200 {array} models.TaxReturn
// @Failure 500 {object} gin.H "Erreur interne"
// @Router /taxreturns [get]
func GetTaxReturns(c *gin.Context) {
	// Get the MongoDB collection
	collection := database.MongoClient.Database("taxdb").Collection("taxreturns") // Replace "taxdb" with your database name

	// Set a timeout for the database query
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Query the database
	var taxReturns []models.TaxReturn
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch tax returns"})
		return
	}
	defer cursor.Close(ctx)

	// Decode the results into the taxReturns slice
	if err := cursor.All(ctx, &taxReturns); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse tax returns"})
		return
	}

	// Return the tax returns as JSON
	c.JSON(http.StatusOK, taxReturns)
}

// CreateTaxReturn ajoute une nouvelle déclaration fiscale
// @Summary Ajouter une déclaration fiscale
// @Description Enregistre une nouvelle déclaration fiscale dans la base de données
// @Tags TaxReturn
// @Accept json
// @Produce json
// @Param taxReturn body models.TaxReturn true "Données de la déclaration fiscale"
// @Success 201 {object} models.TaxReturn
// @Failure 400 {object} gin.H "Données invalides"
// @Router /taxreturns [post]
func CreateTaxReturn(c *gin.Context) {
	var newTaxReturn struct {
		Name string `json:"name" binding:"required"`
	}

	if err := c.ShouldBindJSON(&newTaxReturn); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Données invalides"})
		return
	}

	// Logique d'ajout en base (simulée ici)
	c.JSON(http.StatusCreated, gin.H{"message": "Déclaration fiscale ajoutée", "data": newTaxReturn})
}

// DeleteTaxReturn supprime une déclaration fiscale existante
// @Summary Supprimer une déclaration fiscale
// @Description Supprime une déclaration fiscale par son ID
// @Tags TaxReturn
// @Param id path string true "ID de la déclaration fiscale"
// @Success 200 {object} gin.H "Message de succès"
// @Failure 404 {object} gin.H "Déclaration non trouvée"
// @Router /taxreturns/{id} [delete]
func DeleteTaxReturn(c *gin.Context) {
	id := c.Param("id")

	// Logique de suppression simulée
	c.JSON(http.StatusOK, gin.H{"message": "Déclaration fiscale supprimée", "id": id})
}

// @Summary Get a tax return by ID
// @Description Retrieve a tax return by its ID
