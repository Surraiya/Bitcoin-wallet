package services_test

import (
	"bitcoin-wallet/models"
	"bitcoin-wallet/services"
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

type mockTransactionRepo struct {
	mockGetAll        func() ([]models.Transaction, error)
	mockSave          func(transaction models.Transaction) error
	mockDeleteAll     func() error
	mockUpdate        func(transaction models.Transaction) error
	mockGetAllUnspent func() ([]models.Transaction, error)
}

func (m *mockTransactionRepo) GetAll() ([]models.Transaction, error) {
	if m.mockGetAll != nil {
		return m.mockGetAll()
	}
	return nil, errors.New("mock function not set")
}

func (m *mockTransactionRepo) Save(transaction models.Transaction) error {
	if m.mockSave != nil {
		return m.mockSave(transaction)
	}
	return errors.New("mock function not set")
}

func (m *mockTransactionRepo) DeleteAll() error {
	if m.mockDeleteAll != nil {
		return m.mockDeleteAll()
	}
	return errors.New("mock function not set")
}

func (m *mockTransactionRepo) Update(transaction models.Transaction) error {
	if m.mockUpdate != nil {
		return m.mockUpdate(transaction)
	}
	return errors.New("mock function not set")
}

func (m *mockTransactionRepo) GetAllUnspentTransactions() ([]models.Transaction, error) {
	if m.mockGetAllUnspent != nil {
		return m.mockGetAllUnspent()
	}
	return nil, errors.New("mock function not set")
}

func TestSaveTransaction_Success(t *testing.T) {
	mockRepo := &mockTransactionRepo{
		mockSave: func(transaction models.Transaction) error {
			return nil
		},
	}
	service := services.NewTransactionService(mockRepo)

	transaction := models.Transaction{
		Amount: 10,
		Spent:  false,
	}
	err := service.SaveTransaction(transaction)
	assert.NoError(t, err)
}

func TestSaveTransaction_ValidationFailure(t *testing.T) {
	mockRepo := &mockTransactionRepo{}
	service := services.NewTransactionService(mockRepo)

	transaction := models.Transaction{
		Amount: 0, // Amount is less than or equal to 0, should fail validation
	}
	err := service.SaveTransaction(transaction)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "transaction amount should be more than 0")
}

func TestGetAllTransactions_Success(t *testing.T) {
	expectedTransactions := []models.Transaction{
		{Id: uuid.New().String(), Amount: 100},
		{Id: uuid.New().String(), Amount: 200},
	}
	mockRepo := &mockTransactionRepo{
		mockGetAll: func() ([]models.Transaction, error) {
			return expectedTransactions, nil
		},
	}
	service := services.NewTransactionService(mockRepo)

	transactions, err := service.GetAllTransactions()
	assert.NoError(t, err)
	assert.Equal(t, expectedTransactions, transactions)
}

func TestGetAllTransactions_Error(t *testing.T) {
	mockRepo := &mockTransactionRepo{
		mockGetAll: func() ([]models.Transaction, error) {
			return nil, errors.New("error getting transactions")
		},
	}
	service := services.NewTransactionService(mockRepo)

	_, err := service.GetAllTransactions()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "error getting transactions")
}
