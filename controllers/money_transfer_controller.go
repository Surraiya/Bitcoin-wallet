package controllers

import (
	"bitcoin-wallet/models"
	"bitcoin-wallet/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MoneyTransferController struct {
	service services.MoneyTransferService
}

func NewMoneyTransferController(service services.MoneyTransferService) *MoneyTransferController {
	return &MoneyTransferController{service: service}
}

func (c *MoneyTransferController) CreateNewMoneyTransfer(context *gin.Context) {
	var transfer models.Transfer
	err := context.ShouldBindJSON(&transfer)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data", "error": err.Error()})
		return
	}

	err = c.service.CreateNewMoneyTransfer(transfer.AmountInEuro)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "New Money Transfer Created!", "transfer Amount": transfer})
}
