package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/muyi2905/models"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "encoding/json")
	var user []models.User

	if details := db.Find(&user); details.Error != nil {
		http.Error(w, details.Error.Error(), http.StatusInternalServerError)
	}

	if err := json.NewEncoder(w).Encode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "encoding/json")

	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
	}

	if details := db.Create(&user); details.Error != nil {
		http.Error(w, details.Error.Error(), http.StatusInternalServerError)
	}
}

func GetUserById(w http.ResponseWriter, r *http.Request) {

}
