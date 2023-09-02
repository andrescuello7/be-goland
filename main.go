package main

import (
	"go-postgress/db"
	"go-postgress/models"
	"go-postgress/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	// "github.com/joho/godotenv"
)

func main() {
	//.ENV
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	//Database
	db.PostgresConnect()
	db.DB.AutoMigrate(models.User{})

	//Routes
	router := mux.NewRouter().PathPrefix("/api").Subrouter()
	routes.UserRoutes(router)

	//Server
	log.Printf("Server listening running in port 3000")
	http.ListenAndServe(":3000", router)
}
