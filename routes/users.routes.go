package routes

import (
	"go-postgress/controllers"

	"github.com/gorilla/mux"
)

func UserRoutes(routes *mux.Router) {
	routes.HandleFunc("/users", controllers.GetUserController).Methods("GET")
	routes.HandleFunc("/users", controllers.PostUserController).Methods("POST")
	routes.HandleFunc("/users/{id}", controllers.DeleteUserController).Methods("DELETE")
}
