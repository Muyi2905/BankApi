package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/muyi2905/models"
	"gorm.io/gorm"
)

var DB *gorm.DB

func CreateAccount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "encoding/json")
	var account models.Account
	json.NewDecoder(r.Body).Decode(&account)
	if err!=

}
