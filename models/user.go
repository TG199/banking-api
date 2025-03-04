package models

import "gorm.i/gorm"

type User struct {
	gorm.Model
	Name string `json:"name"`
	Email sring `json: "email"`
}
