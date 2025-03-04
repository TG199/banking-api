package database

import "github.com/TG199/banking-api/models"

func Migrate() {
	DB.AutoMigrate(&models.User{})
}
