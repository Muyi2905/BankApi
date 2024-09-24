package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/muyi2905/models"
	"gorm.io/gorm"
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
	vars := mux.Vars(r)
	id := vars["id"]

	var user models.User
	if details := db.Find(&user, "id=?", id); details.Error != nil {
		if details.Error == gorm.ErrRecordNotFound {
			http.Error(w, "user not found", http.StatusNotFound)
		} else {
			http.Error(w, details.Error.Error(), http.StatusInternalServerError)
		}
		return
	}
	w.Header().Set("Content-Type", "encoding/json")
	json.NewEncoder(w).Encode(user)

}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var user models.User
	if details := db.Find(&user, "id=?", id); details.Error != nil {
		if details.Error == gorm.ErrRecordNotFound {
			http.Error(w, "user not found", http.StatusNotFound)
		} else {
			http.Error(w, details.Error.Error(), http.StatusInternalServerError)
		}
		return
	}

	var UpdatedUser models.User
	if err := json.NewDecoder(r.Body).Decode(&UpdatedUser); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if details := db.Save(&user); details.Error != nil {
		http.Error(w, "Error saving updated information", http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(user)

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "encoding/json")
	vars := mux.Vars(r)
	id := vars["id"]

	var user models.User
	if result := db.First(&user, "id=?", id); result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			http.Error(w, "user not found", http.StatusNotFound)
		} else {
			http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		}
	}

	if result := db.Delete(&user); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
}
