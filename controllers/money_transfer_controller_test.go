package controllers_test

import (
	"bitcoin-wallet/controllers"
	"bitcoin-wallet/models"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type mockMoneyTransferService struct {
	mockCreateNewMoneyTransfer func(amount float64) error
}

func (m *mockMoneyTransferService) CreateNewMoneyTransfer(amount float64) error {
	if m.mockCreateNewMoneyTransfer != nil {
		return m.mockCreateNewMoneyTransfer(amount)
	}
	return nil
}

func TestCreateNewMoneyTransfer_Success(t *testing.T) {
	mockService := &mockMoneyTransferService{
		mockCreateNewMoneyTransfer: func(amount float64) error {
			return nil
		},
	}
	controller := controllers.NewMoneyTransferController(mockService)

	// Create a test request body
	requestBody := models.Transfer{AmountInEuro: 100.0}
	jsonBody, _ := json.Marshal(requestBody)

	// Create a test HTTP request
	req, _ := http.NewRequest("POST", "/money-transfer", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	// Create a test HTTP response recorder
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	// Call the controller method
	controller.CreateNewMoneyTransfer(c)

	// Check if the response status code is 201 (Created)
	assert.Equal(t, http.StatusCreated, w.Code)

	// Check if the response body contains the expected message
	expectedResponse := `{"message":"New Money Transfer Created!","transfer Amount":{"amount_in_euro":100}}`
	assert.Equal(t, expectedResponse, w.Body.String())
}
