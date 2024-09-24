package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
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
	if results := db.Find(&accounts); results.Error != nil {
		http.Error(w, results.Error.Error(), http.StatusInternalServerError)
	}

	if err := json.NewEncoder(w).Encode(&accounts); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func GetAccountById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var account models.Account
	if result := db.First(&account, "id=?", id); result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			http.Error(w, "account not found", http.StatusNotFound)
		} else {
			http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		}
		return
	}
	w.Header().Set("Content-Type", "encoding/json")
	if err := json.NewEncoder(w).Encode(&account); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func UpdateAccount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]

	var account models.Account
	if result := db.First(&account, "id=?", id); result.Error != nil {
		http.Error(w, "Account not found", http.StatusNotFound)
	}

	var UpdatedAccount models.Account
	if err := json.NewDecoder(r.Body).Decode(&UpdatedAccount); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	account.AccountBalance = UpdatedAccount.AccountBalance
	if result := db.Save(&account); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(account)
}

func DeleteAccount(w http.ResponseWriter, r *http.Request) {

}
