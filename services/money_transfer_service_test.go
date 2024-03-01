package services_test

import (
	"bitcoin-wallet/models"
	"bitcoin-wallet/services"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreateNewMoneyTransfer_InsufficientBalance(t *testing.T) {
	expectedTransactions := []models.Transaction{
		{Id: uuid.New().String(), Amount: 10},
		{Id: uuid.New().String(), Amount: 20},
	}
	mockRepo := &mockTransactionRepo{
		mockGetAll: func() ([]models.Transaction, error) {
			return expectedTransactions, nil
		},
	}
	service := services.NewMoneyTransferService(mockRepo)

	err := service.CreateNewMoneyTransfer(5332123)
	assert.Error(t, err)
	assert.Equal(t, "insufficient balance. Unable to transfer money", err.Error())
}

func TestCreateNewMoneyTransfer_MinimumAmount(t *testing.T) {
	expectedTransactions := []models.Transaction{
		{Id: uuid.New().String(), Amount: 0.5},
		{Id: uuid.New().String(), Amount: 0.4},
	}
	mockRepo := &mockTransactionRepo{
		mockGetAll: func() ([]models.Transaction, error) {
			return expectedTransactions, nil
		},
	}
	service := services.NewMoneyTransferService(mockRepo)

	err := service.CreateNewMoneyTransfer(0.000009)

	assert.Equal(t, "amount too small. Minimum amount to transfer is 0.00001 BTC", err.Error())
}
