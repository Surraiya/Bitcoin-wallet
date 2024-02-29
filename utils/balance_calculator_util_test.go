package utils_test

import (
	"bitcoin-wallet/models"
	"bitcoin-wallet/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateCurrentBalance(t *testing.T) {
	// Create test transactions
	transactions := []models.Transaction{
		{Amount: 10, Spent: false},
		{Amount: 20, Spent: false},
		{Amount: 5, Spent: true},
	}

	// Call the function to calculate the balance
	balance := utils.CalculateCurrentBalance(transactions)

	// Expected balance calculation: (10 + 20) - 5 = 25
	expectedBalance := 25.0

	// Assert that the calculated balance matches the expected balance
	assert.Equal(t, expectedBalance, balance)
}

func TestCalculateCurrentBalance_NoTransactions(t *testing.T) {
	// No transactions provided
	var transactions []models.Transaction

	// Call the function to calculate the balance
	balance := utils.CalculateCurrentBalance(transactions)

	// Expected balance calculation: 0 (no transactions)
	expectedBalance := 0.0

	// Assert that the calculated balance matches the expected balance
	assert.Equal(t, expectedBalance, balance)
}

func TestCalculateCurrentBalance_AllUnspent(t *testing.T) {
	// Create test transactions where all are unspent
	transactions := []models.Transaction{
		{Amount: 10, Spent: false},
		{Amount: 20, Spent: false},
		{Amount: 5, Spent: false},
	}

	// Call the function to calculate the balance
	balance := utils.CalculateCurrentBalance(transactions)

	// Expected balance calculation: 10 + 20 + 5 = 35
	expectedBalance := 35.0

	// Assert that the calculated balance matches the expected balance
	assert.Equal(t, expectedBalance, balance)
}

func TestCalculateCurrentBalance_AllSpent(t *testing.T) {
	// Create test transactions where all are spent
	transactions := []models.Transaction{
		{Amount: 10, Spent: true},
		{Amount: 20, Spent: true},
		{Amount: 5, Spent: true},
	}

	// Call the function to calculate the balance
	balance := utils.CalculateCurrentBalance(transactions)

	// Expected balance calculation: 0 (all transactions are spent)
	expectedBalance := 0.0

	// Assert that the calculated balance matches the expected balance
	assert.Equal(t, expectedBalance, balance)
}
