package routes

import (
	"github.com/nvancuong2/taxApiGoReact/backend/controllers"

	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/example", controllers.ExampleHandler).Methods("GET")
	return router
}
