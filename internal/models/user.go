package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name string `gorm:"uniqueIndex" json:"username"`
	Email string `gorm:"uniqueIndex" json:"email"`
	Password string `json:"-"`
	Acounts []Account
}
