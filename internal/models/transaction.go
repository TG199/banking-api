package models

import (
	"gorm.io/gorm"
)
type Transaction struct {
	gorm.Model
	AccountID uint		`json:"account_id"`
	Amount	  float64	`json:"amount"`
	Type	  string 	`json:"type"`
}

