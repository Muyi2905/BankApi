package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	AccountNumber string    `json:"AccountNumber"`
	PhoneNo       int       `json:"PhoneNo"`
	Account       []Account `gorm:"foreignKey:UserID" json:"accounts"`
}
