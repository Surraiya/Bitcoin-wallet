package utils_test

import (
	"bitcoin-wallet/models"
	"bitcoin-wallet/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateCurrentBalance(t *testing.T) {
	transactions := []models.Transaction{
		{Amount: 10, Spent: false},
		{Amount: 20, Spent: false},
		{Amount: 5, Spent: true},
	}

	balance := utils.CalculateCurrentBalance(transactions)

	expectedBalance := 30.0

	assert.Equal(t, expectedBalance, balance, "balance should be equal to 30.0")
}

func TestCalculateCurrentBalance_NoTransactions(t *testing.T) {
	var transactions []models.Transaction

	balance := utils.CalculateCurrentBalance(transactions)

	expectedBalance := 0.0

	assert.Equal(t, expectedBalance, balance, "balance should be equal to 0.0")
}

func TestCalculateCurrentBalance_AllUnspent(t *testing.T) {
	transactions := []models.Transaction{
		{Amount: 10, Spent: false},
		{Amount: 20, Spent: false},
		{Amount: 5, Spent: false},
	}

	balance := utils.CalculateCurrentBalance(transactions)

	// Expected balance calculation: 10 + 20 + 5 = 35
	expectedBalance := 35.0

	assert.Equal(t, expectedBalance, balance, "balance should be equal to 35")
}

func TestCalculateCurrentBalance_AllSpent(t *testing.T) {
	// Create test transactions where all are spent
	transactions := []models.Transaction{
		{Amount: 10, Spent: true},
		{Amount: 20, Spent: true},
		{Amount: 5, Spent: true},
	}

	balance := utils.CalculateCurrentBalance(transactions)

	expectedBalance := 0.0

	assert.Equal(t, expectedBalance, balance, "Balance should be 0.0 as all spent is true")
}
