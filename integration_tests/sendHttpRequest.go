package integrationtests

import (
	"bitcoin-wallet/controllers"
	"bitcoin-wallet/models"
	"bitcoin-wallet/repositories"
	"bitcoin-wallet/services"
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

func getTransactionController() *controllers.TransactionController {
	transactionRepo := repositories.NewTransactionRepository()
	transactionService := services.NewTransactionService(transactionRepo)
	return controllers.NewTransactionController(transactionService)

}

func getCurrentBalanceController() *controllers.CurrentBalanceController {
	transactionRepo := repositories.NewTransactionRepository()
	balanceService := services.NewCurrentBalanceService(transactionRepo)
	return controllers.NewCurrentBalanceController(balanceService)

}

func getMoneyTransferController() *controllers.MoneyTransferController {
	transactionRepo := repositories.NewTransactionRepository()
	transferService := services.NewMoneyTransferService(transactionRepo)
	return controllers.NewMoneyTransferController(transferService)

}

func performRequest(req *http.Request, handler http.Handler) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	return rr
}

func unmarshalResponse(rr *httptest.ResponseRecorder, target interface{}) {
	err := json.Unmarshal(rr.Body.Bytes(), target)
	if err != nil {
		log.Fatal(err)
	}
}

func jsonMarshal(target interface{}) []byte {
	jsonPayload, err := json.Marshal(target)
	if err != nil {
		log.Fatal(err)
	}
	return jsonPayload
}

func GetActualTransactionsFromServer() ([]models.Transaction, *httptest.ResponseRecorder) {
	transactionController := getTransactionController()

	req, err := http.NewRequest("GET", "/transactions", nil)
	if err != nil {
		log.Fatal(err)
	}

	router := gin.Default()
	router.GET("/transactions", transactionController.GetAllTransactions)

	rr := performRequest(req, router)

	var actualTransactions []models.Transaction
	unmarshalResponse(rr, &actualTransactions)
	return actualTransactions, rr
}

func GetActualBalanceFromServer() (models.Balance, *httptest.ResponseRecorder) {
	balanceController := getCurrentBalanceController()

	req, err := http.NewRequest("GET", "/current-balance", nil)
	if err != nil {
		log.Fatal(err)
	}

	router := gin.Default()
	router.GET("/current-balance", balanceController.GetCurrentBalance)

	rr := performRequest(req, router)

	var actualBalance models.Balance
	unmarshalResponse(rr, &actualBalance)
	return actualBalance, rr
}

func CreateMoneyTransferRequest(transfer models.Transfer) *httptest.ResponseRecorder {
	transferController := getMoneyTransferController()
	jsonPayload := jsonMarshal(transfer)

	req, err := http.NewRequest("POST", "/money-transfers", bytes.NewBuffer(jsonPayload))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Serve HTTP
	router := gin.Default()
	router.POST("/money-transfers", transferController.CreateNewMoneyTransfer)
	rr := performRequest(req, router)

	return rr
}
