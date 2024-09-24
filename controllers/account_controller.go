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

}
