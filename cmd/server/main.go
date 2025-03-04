package main

import (

	"github.com/gin-gonic/gin"
	"github.com/TG199/banking-api/database"
)

func main() {

	database.ConnectDB()


	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})
	
	r.GET("/me", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hi, my name is Kelechi Ebiri"})
	})

	r.Run(":8080")
}
