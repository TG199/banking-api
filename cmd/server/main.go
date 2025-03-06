package main

import (
	"github.com/TG199/banking-api/database"
	"github.com/TG199/banking-api/routes"
)

func main() {
	database.ConnectDB()

	r := routes.SetupRouter()

	r.Run(":8080")
}
