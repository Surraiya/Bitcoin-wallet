package repositories

import (
	"bitcoin-wallet/database"
	"bitcoin-wallet/models"
	"fmt"
)

type TransactionRepository interface {
	GetAll() ([]models.Transaction, error)
	Save(transaction models.Transaction) error
	DeleteAll() error
	Update(transaction models.Transaction) error
	GetAllUnspentTransactions() ([]models.Transaction, error)
}

type transactionRepository struct{}

func NewTransactionRepository() TransactionRepository {
	return &transactionRepository{}
}

func (r *transactionRepository) GetAll() ([]models.Transaction, error) {
	var transactions []models.Transaction
	if err := database.GetDB().Find(&transactions).Error; err != nil {
		return nil, fmt.Errorf("error retrieving all transactions: %v", err)
	}
	return transactions, nil
}

func (r *transactionRepository) GetAllUnspentTransactions() ([]models.Transaction, error) {
	var transactions []models.Transaction
	if err := database.GetDB().Model(&models.Transaction{}).Where("spent = ?", false).Find(&transactions).Error; err != nil {
		return nil, fmt.Errorf("error retrieving unspent transactions: %v", err)
	}
	return transactions, nil
}

func (r *transactionRepository) Save(transaction models.Transaction) error {
	if err := database.GetDB().Create(&transaction).Error; err != nil {
		return fmt.Errorf("error creating a new transaction: %v", err)
	}
	return nil
}

func (r *transactionRepository) Update(transaction models.Transaction) error {
	if err := database.GetDB().Model(&models.Transaction{}).Where("id = ?", transaction.Id).Updates(transaction).Error; err != nil {
		return fmt.Errorf("error updating a transaction: %v", err)
	}
	return nil
}

func (r *transactionRepository) DeleteAll() error {
	return database.GetDB().Delete(&models.Transaction{}).Error
}
