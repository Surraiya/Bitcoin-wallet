package controllers

import (
	"bitcoin-wallet/models"
	"bitcoin-wallet/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TransactionController struct {
	service services.TransactionService
}

func NewTransactionController(service services.TransactionService) *TransactionController {
	return &TransactionController{service: service}
}

func (c *TransactionController) GetAllTransactions(context *gin.Context) {
	transactions, err := c.service.GetAllTransactions()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	context.JSON(http.StatusOK, transactions)
}

func (c *TransactionController) CreateNewTransaction(context *gin.Context) {
	var transaction models.Transaction

	if err := context.ShouldBindJSON(&transaction); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := c.service.SaveTransaction(transaction)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Transaction Created!", "transaction": transaction})
}

func (c *TransactionController) DeleteAllTransactions(context *gin.Context) {
	err := c.service.DeleteAllTransactions()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete all transactions."})
		return
	}
}
