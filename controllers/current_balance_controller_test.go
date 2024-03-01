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
	return models.Balance{BTC: 100.0, Euro: 500.0}, nil
}

func TestGetCurrentBalance_Success(t *testing.T) {
	controller := controllers.NewCurrentBalanceController(&mockCurrentBalanceService{})

	req, err := http.NewRequest("GET", "/current-balance", nil)
	assert.NoError(t, err)

	recorder := httptest.NewRecorder()

	context, _ := gin.CreateTestContext(recorder)
	context.Request = req

	controller.GetCurrentBalance(context)

	assert.Equal(t, http.StatusOK, recorder.Code, "Status code should be http.StatusOK")
	assert.Equal(t, `{"BTC":100,"Euro":500}`, recorder.Body.String(), "Response body should match expected JSON")
}
