package controllers_test

import (
	"bitcoin-wallet/controllers"
	"bitcoin-wallet/models"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type mockTransactionService struct {
	mockGetAllTransactions    func() ([]models.Transaction, error)
	mockSaveTransaction       func(transaction models.Transaction) error
	mockDeleteAllTransactions func() error
}

func (m *mockTransactionService) GetAllTransactions() ([]models.Transaction, error) {
	if m.mockGetAllTransactions != nil {
		return m.mockGetAllTransactions()
	}
	return nil, errors.New("mock function not set")
}

func (m *mockTransactionService) SaveTransaction(transaction models.Transaction) error {
	if m.mockSaveTransaction != nil {
		return m.mockSaveTransaction(transaction)
	}
	return errors.New("mock function not set")
}

func (m *mockTransactionService) DeleteAllTransactions() error {
	return nil
}

func TestGetAllTransactions_Success(t *testing.T) {
	mockService := &mockTransactionService{
		mockGetAllTransactions: func() ([]models.Transaction, error) {
			return []models.Transaction{
				{Amount: 10, Spent: false},
				{Amount: 20, Spent: false},
			}, nil
		},
	}
	controller := controllers.NewTransactionController(mockService)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	controller.GetAllTransactions(c)

	assert.Equal(t, http.StatusOK, w.Code)

	var response []map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	expectedJSON := `[{"Amount":10,"Spent":false},{"Amount":20,"Spent":false}]`

	var expected []map[string]interface{}
	err = json.Unmarshal([]byte(expectedJSON), &expected)
	assert.NoError(t, err)

	for i, item := range expected {
		assert.Equal(t, item["Amount"], response[i]["Amount"])
		assert.Equal(t, item["Spent"], response[i]["Spent"])
	}
}

func TestGetAllTransactions_Error(t *testing.T) {
	mockService := &mockTransactionService{
		mockGetAllTransactions: func() ([]models.Transaction, error) {
			return nil, errors.New("error getting transactions")
		},
	}
	controller := controllers.NewTransactionController(mockService)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	controller.GetAllTransactions(c)

	assert.Equal(t, http.StatusInternalServerError, w.Code)

	expectedResponse := `{"error":"error getting transactions"}null`
	actualResponse := strings.TrimSpace(w.Body.String()) // Trim whitespace
	assert.Equal(t, expectedResponse, actualResponse)
}

func TestCreateNewTransaction_Success(t *testing.T) {
	mockService := &mockTransactionService{
		mockSaveTransaction: func(transaction models.Transaction) error {
			return nil
		},
	}
	controller := controllers.NewTransactionController(mockService)

	requestBody := `{"Amount":10}`
	req, _ := http.NewRequest("POST", "/transactions", strings.NewReader(requestBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	controller.CreateNewTransaction(c)

	assert.Equal(t, http.StatusCreated, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	expectedResponse := `{"message":"Transaction Created!","transaction":{"Amount":10}}`

	var expected map[string]interface{}
	err = json.Unmarshal([]byte(expectedResponse), &expected)
	assert.NoError(t, err)

	assert.Equal(t, expected["message"], response["message"])
	assert.Equal(t, expected["transaction"].(map[string]interface{})["Amount"], response["transaction"].(map[string]interface{})["Amount"])
}

func TestCreateNewTransaction_InvalidAmount(t *testing.T) {
	mockService := &mockTransactionService{
		mockSaveTransaction: func(transaction models.Transaction) error {
			if transaction.Amount <= 0 {
				return errors.New("transaction amount should be more than 0")
			}
			return nil
		},
	}
	controller := controllers.NewTransactionController(mockService)

	requestBody := `{"Amount":0}`
	req, _ := http.NewRequest("POST", "/transactions", strings.NewReader(requestBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	controller.CreateNewTransaction(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	expectedResponse := `{"error":"transaction amount should be more than 0"}`
	assert.Equal(t, expectedResponse, strings.TrimSpace(w.Body.String()))
}
