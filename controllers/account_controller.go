package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/muyi2905/models"
	"gorm.io/gorm"
)

var db *gorm.DB

func CreateAccount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "encoding/json")
	var account models.Account
	err := json.NewDecoder(r.Body).Decode(&account)
	if err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
	}
	result := db.Create(&account)
	if result.Error != nil {
		http.Error(w, "error creating account", http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(account)
}

func GetAccounts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "encoding/json")
	var accounts []models.Account
	results := db.Find(&accounts); results.Error!= nil{
		http.Error(w, results.Error.Error(), http.StatusInternalServerError)
		return 
	}
	json.NewEncoder(w).Encode(&accounts)

}
