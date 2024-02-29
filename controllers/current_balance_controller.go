package controllers

import (
	"bitcoin-wallet/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CurrentBalanceController struct {
	service services.CurrentBalanceService
}

func NewCurrentBalanceController(service services.CurrentBalanceService) *CurrentBalanceController {
	return &CurrentBalanceController{service: service}
}

func (c *CurrentBalanceController) GetCurrentBalance(context *gin.Context) {
	balance, err := c.service.GetCurrentBalance()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}

	context.JSON(http.StatusOK, balance)
}
