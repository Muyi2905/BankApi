package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/muyi2905/models"
	"gorm.io/gorm"
)

var db *gorm.DB

func CreateAccount(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var account models.Account
		err := json.NewDecoder(r.Body).Decode(&account)
		if err != nil {
			http.Error(w, "invalid request", http.StatusBadRequest)
			return
		}
		result := db.Create(&account)
		if result.Error != nil {
			http.Error(w, "error creating account", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(account)
	}
}

func GetAccounts(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var accounts []models.Account
		if results := db.Find(&accounts); results.Error != nil {
			http.Error(w, results.Error.Error(), http.StatusInternalServerError)
			return
		}

		if err := json.NewEncoder(w).Encode(&accounts); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func GetAccountById(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(&account); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func UpdateAccount(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		vars := mux.Vars(r)
		id := vars["id"]

		var account models.Account
		if result := db.First(&account, "id=?", id); result.Error != nil {
			http.Error(w, "Account not found", http.StatusNotFound)
			return
		}

		var UpdatedAccount models.Account
		if err := json.NewDecoder(r.Body).Decode(&UpdatedAccount); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		account.AccountBalance = UpdatedAccount.AccountBalance
		if result := db.Save(&account); result.Error != nil {
			http.Error(w, result.Error.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(account)
	}
}

func DeleteAccount(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		vars := mux.Vars(r)
		id := vars["id"]

		var account models.Account
		if result := db.First(&account, "id=?", id); result.Error != nil {
			if result.Error == gorm.ErrRecordNotFound {
				http.Error(w, "Account not found", http.StatusNotFound)
			} else {
				http.Error(w, result.Error.Error(), http.StatusInternalServerError)
			}
			return
		}

		if result := db.Delete(&account); result.Error != nil {
			http.Error(w, result.Error.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
