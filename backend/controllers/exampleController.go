package controllers

import (
	"encoding/json"
	"net/http"
)

// ExampleHandler handles example requests
// @Summary Example endpoint
// @Description This is an example endpoint that returns a simple message.
// @Tags Example
// @Accept json
// @Produce json
// @Success 200 {object} map[string]string
// @Router /example [get]
func ExampleHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"message": "Hello, World!"}
	json.NewEncoder(w).Encode(response)
}
