package models

import (
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model

	UserID         uint    `json:"userId"`
	AccountBalance float64 `json:"accountBalance"`
}
