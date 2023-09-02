package controllers

import (
	"encoding/json"
	"go-postgress/db"
	"go-postgress/models"
	"net/http"

	"github.com/gorilla/mux"
)

func GetUserController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var users []models.User
	db.DB.Find(&users)

	json.NewEncoder(w).Encode(&users)
}

func PostUserController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user models.User
	json.NewDecoder(r.Body).Decode(&user)
	createUser := db.DB.Create(&user)

	err := createUser.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error creating user"))
	}

	json.NewEncoder(w).Encode(&user)
}

func DeleteUserController(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var user models.User
	deleteUser := db.DB.First(&user, params["id"])

	err := deleteUser.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error delete user"))
	}

	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Don't repet user"))
		return
	}

	db.DB.Unscoped().Delete(&user)
	w.WriteHeader(http.StatusOK)
}
