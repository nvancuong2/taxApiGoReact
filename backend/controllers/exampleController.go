package controllers

import (
	"encoding/json"
	"net/http"
)

func ExampleHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"message": "Hello, World!"}
	json.NewEncoder(w).Encode(response)
}
