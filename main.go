package main

import (
	"bitcoin-wallet/controllers"
	"bitcoin-wallet/database"
	"bitcoin-wallet/repositories"
	"bitcoin-wallet/services"
	"log"

	"github.com/gin-gonic/gin"
)

const (
	Port           = ":8083"
	Transactions   = "/transactions"
	Transfer       = "/money-transfers"
	CurrentBalance = "/current-balance"
)

func main() {
	err := database.InitializeDB("bitcoin_wallet.db")
	if err != nil {
		log.Fatalf("Failed to initialize database connection: %v", err)
	}
	defer database.GetDB().Close()

	//Initialize repositories
	transactionRepo := repositories.NewTransactionRepository()

	//Initialize services and pass repositories
	transactionService := services.NewTransactionService(transactionRepo)
	currentBalanceService := services.NewCurrentBalanceService(transactionRepo)
	moneyTransferService := services.NewMoneyTransferService(transactionRepo)

	//Initialize API handlers and pass services
	transactionController := controllers.NewTransactionController(transactionService)
	currentBalanceController := controllers.NewCurrentBalanceController(currentBalanceService)
	moneyTransferController := controllers.NewMoneyTransferController(moneyTransferService)

	//Initialize Gin router
	server := gin.Default()

	server.GET(Transactions, transactionController.GetAllTransactions)
	server.POST(Transactions, transactionController.CreateNewTransaction)
	server.DELETE(Transactions, transactionController.DeleteAllTransactions)
	server.GET(CurrentBalance, currentBalanceController.GetCurrentBalance)
	server.POST(Transfer, moneyTransferController.CreateNewMoneyTransfer)

	server.Run(Port)
}
