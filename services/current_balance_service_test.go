package services_test

import (
	"bitcoin-wallet/models"
	"bitcoin-wallet/services"
	"math"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestGetCurrentBalance_Success(t *testing.T) {
	expectedTransactions := []models.Transaction{
		{Id: uuid.New().String(), Amount: 0.1},
		{Id: uuid.New().String(), Amount: 0.2},
	}
	mockRepo := &mockTransactionRepo{
		mockGetAll: func() ([]models.Transaction, error) {
			return expectedTransactions, nil
		},
	}
	service := services.NewCurrentBalanceService(mockRepo)

	expectedBalance := 0.1 + 0.2
	expectedBalanceRounded := round(expectedBalance, 2) // Round to 2 decimal places

	balance, err := service.GetCurrentBalance()
	assert.NoError(t, err)
	assert.Equal(t, expectedBalanceRounded, round(balance.BTC, 2))
}

func round(num float64, places int) float64 {
	shift := math.Pow(10, float64(places))
	return math.Round(num*shift) / shift
}
