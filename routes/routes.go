package routes

import (
	"github.com/TG199/banking-api/internal/handlers"
	"github.com/TG199/banking-api/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Auth routes
	r.POST("/signup", handlers.SignUp)
	r.POST("/login", handlers.Login)

	protected := r.Group("/")
	protected.Use(middleware.AuthMiddleware())
	// Account routes
	protected.POST("/deposit", handlers.Deposit)
	protected.GET("/withdraw", handlers.Withdraw)
	protected.GET("/transactions", handlers.TransactionHistory)
	protected.POST("/transfer", handlers.TransferFunds)
	protected.GET("/balance", handlers.GetBalance)

	r.GET("/ws", handlers.HandleConnections)
    go handlers.HandleMessages()




	return r
}