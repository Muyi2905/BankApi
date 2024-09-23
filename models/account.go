package models

import (
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model

	UserId         string  `json:"userid"`
	AccountBalance float64 `json:"accountbalance"`
}
