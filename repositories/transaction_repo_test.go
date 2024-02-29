package repositories_test

import (
	"bitcoin-wallet/database"
	"bitcoin-wallet/models"
	"bitcoin-wallet/repositories"
	"bitcoin-wallet/utils"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTransactionRepository_GetAll(t *testing.T) {
	testDb := database.SetupTestDB()
	repo := repositories.NewTransactionRepository()

	// Test empty database
	transactions, err := repo.GetAll()
	assert.NoError(t, err)
	assert.Empty(t, transactions)

	// Test saving transaction
	testTransaction := models.Transaction{
		Id:     utils.RandomUniqueStringGenerator(),
		Amount: 1,
		Spent:  false,
	}
	err = repo.Save(testTransaction)
	assert.NoError(t, err)

	// Test getting all transactions after saving
	transactions, err = repo.GetAll()
	log.Println(transactions)
	assert.NoError(t, err)
	assert.NotEmpty(t, transactions)
	assert.Equal(t, 1, len(transactions))
	assert.Equal(t, testTransaction.Amount, transactions[0].Amount)
	assert.Equal(t, testTransaction.Spent, transactions[0].Spent)

	database.CleanupTestDB(testDb)
}

func TestTransactionRepository_Save(t *testing.T) {
	testDb := database.SetupTestDB()

	repo := repositories.NewTransactionRepository()

	testTransaction := models.Transaction{
		Id:     utils.RandomUniqueStringGenerator(),
		Amount: 2,
		Spent:  false,
	}
	err := repo.Save(testTransaction)
	assert.NoError(t, err)

	// Retrieve the saved transaction and check its correctness
	var retrievedTransaction models.Transaction
	err = database.GetDB().First(&retrievedTransaction).Error
	log.Println(retrievedTransaction)
	assert.NoError(t, err)
	assert.Equal(t, testTransaction.Amount, retrievedTransaction.Amount)
	assert.Equal(t, testTransaction.Spent, retrievedTransaction.Spent)

	database.CleanupTestDB(testDb)
}
