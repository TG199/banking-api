package handlers

import (
	"net/http"

	"github.com/TG199/banking-api/database"
	"github.com/TG199/banking-api/internal/models"
	"github.com/gin-gonic/gin"
)

func CreateAccount(c *gin.Context) {

}
func Deposit(c *gin.Context) {
	var txn models.Transaction
	if err := c.ShouldBindJSON(&txn); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var account models.Account
	database.DB.First(&account, txn.AccountID)

	if account.ID == 0 {
		c.JSON((http.StatusNotFound), gin.H{"Deposit error": "Account not found"})
		return
	}

	account.Balance += txn.Amount
	txn.Type = "deposit"

	database.DB.Save(&account)
	database.DB.Create(&txn)

	c.JSON(http.StatusOK, gin.H{"message": "Deposit successful", "balance": account.Balance})
}

func Withdraw(c *gin.Context) {
	var txn models.Transaction
	if err := c.ShouldBindJSON(&txn); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var account models.Account
	database.DB.First(&account, txn.AccountID)
	if account.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Account not found"})
		return
	}

	if account.Balance < txn.Amount {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Insufficient funds"})
		return
	}

	account.Balance -= txn.Amount
	txn.Type = "withdrawal"

	database.DB.Save(&account)
	database.DB.Create(&txn)

	c.JSON(http.StatusOK, gin.H{"message": "Withdrawal succesful", "balance": account.Balance})
}

func TransactionHistory(c *gin.Context) {
	userID := c.GetUint("user_id")


	var transactions []models.Transaction
	if err := database.DB.Where("user_id = ?", userID).Find(&transactions).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error":"Failed to retrieve transactions"})
		return
	}

	c.JSON(http.StatusOK, transactions)
}

func TransferFunds(c *gin.Context) {
	var transferData struct {
		FromAccountID uint 		`json:"from_account_id"`
		ToAccountID   uint 		`json:"to_account_id"`
		Amount		  float64	`json:"amount"`
	}

	if err := c.ShouldBindJSON(&transferData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	userID := c.GetUint("user_id")

	var sender models.Account

	if err := database.DB.Where("id = ? AND user_id = ?", transferData.FromAccountID, userID).First(&sender).Error; err != nil{
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Sender account not found or Unauthorized"})
		return
	}

	var recipient models.Account

	if err := database.DB.Where("id = ?", transferData.ToAccountID).First(&recipient).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Recipient account not found"})
		return
	}

	if sender.Balance < transferData.Amount {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Insufficient funds"})
		return
	}

	sender.Balance -= transferData.Amount
	recipient.Balance += transferData.Amount


	tx := database.DB.Begin()

	if err := tx.Save(&sender).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Transaction failed"})
		return
	}
	tx.Commit()

	go BroadcastBalanceUpdate(sender)
    go BroadcastBalanceUpdate(recipient)
	c.JSON(http.StatusOK, gin.H{"message": "Transfer Succcesful"})
}

func GetBalance(c *gin.Context) {
	userId := c.GetUint("user_id")

	var account models.Account

	if err := database.DB.Where("user_id = ?", userId).First(&account).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error":"Account not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"balance": account.Balance})
}