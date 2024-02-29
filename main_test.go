package main

import (
	"bitcoin-wallet/controllers"
	"bitcoin-wallet/repositories"
	"bitcoin-wallet/services"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGetAllTransactions(t *testing.T) {
	// Initialize repositories
	transactionRepo := repositories.NewTransactionRepository()

	// Initialize services and pass repositories
	transactionService := services.NewTransactionService(transactionRepo)

	// Initialize API handlers and pass services
	transactionController := controllers.NewTransactionController(transactionService)

	// Set up your test database and configure your application to use it
	// This might involve calling database.InitializeDB with your test database

	// Initialize your application, including setting up the Gin router
	// This should be similar to what you do in your main function, but with the test database

	// Create a request to pass to our handler.
	req, err := http.NewRequest("GET", "/transactions", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response.
	rr := httptest.NewRecorder()

	// Create a new Gin engine and pass it to the handler.
	router := gin.Default()
	router.GET("/transactions", transactionController.GetAllTransactions)

	// Perform the request
	router.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Add more assertions here to check the response body

	// Clean up the test database
}

// func TestShowBalance(t *testing.T) {
// 	// Similar to TestListTransactions, but for the /balance endpoint
// }

// func TestCreateTransfer(t *testing.T) {
// 	// Create a JSON payload for the transfer request
// 	transferPayload := map[string]interface{}{
// 		"amount": 10.0, // Example amount
// 	}
// 	jsonPayload, _ := json.Marshal(transferPayload)

// 	req, err := http.NewRequest("POST", "/transfer", bytes.NewBuffer(jsonPayload))
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	req.Header.Set("Content-Type", "application/json")

// 	rr := httptest.NewRecorder()
// 	handler := http.HandlerFunc(createTransfer)
// 	handler.ServeHTTP(rr, req)

// 	if status := rr.Code; status != http.StatusOK {
// 		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
// 	}

// 	// Add more assertions here to check the response body and the state of transactions
// }
