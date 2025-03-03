package main

import (

	"github.com/gin-gonic/gin"
	"github.com/TG199/banking-api/database"
	"net/http"
)

func main() {

	database.ConnectDB()


	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	r.Run(":8080")
}
