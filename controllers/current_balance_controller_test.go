package controllers_test

import (
	"bitcoin-wallet/controllers"
	"bitcoin-wallet/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type mockCurrentBalanceService struct{}

func (m *mockCurrentBalanceService) GetCurrentBalance() (models.Balance, error) {
	// Mocking a successful response with balances of 100.0 BTC and 500.0 Euro
	return models.Balance{BTC: 100.0, Euro: 500.0}, nil
}

func TestGetCurrentBalance_Success(t *testing.T) {
	// Create a new instance of the controller with the mock service
	controller := controllers.NewCurrentBalanceController(&mockCurrentBalanceService{})

	// Create a new HTTP request
	req, err := http.NewRequest("GET", "/current-balance", nil)
	assert.NoError(t, err)

	// Create a response recorder to record the response
	recorder := httptest.NewRecorder()

	// Create a new Gin context from the response recorder and request
	context, _ := gin.CreateTestContext(recorder)
	context.Request = req

	// Call the controller method
	controller.GetCurrentBalance(context)

	// Assert the status code is OK
	assert.Equal(t, http.StatusOK, recorder.Code)

	// Assert the response body contains the correct balances
	assert.Equal(t, `{"BTC":100,"Euro":500}`, recorder.Body.String())
}
