package services

import (
	"bitcoin-wallet/models"
	"bitcoin-wallet/repositories"
	"bitcoin-wallet/utils"
	"fmt"
)

type TransactionService interface {
	GetAllTransactions() ([]models.Transaction, error)
	SaveTransaction(transaction models.Transaction) error
	DeleteAllTransactions() error
}

type transactionService struct {
	repo repositories.TransactionRepository
}

func NewTransactionService(repo repositories.TransactionRepository) TransactionService {
	return &transactionService{repo: repo}
}

func (s *transactionService) SaveTransaction(transaction models.Transaction) error {
	if err := validator(transaction); err != nil {
		//	return err.Error()
		return err
	}

	transaction.Id = utils.RandomUniqueStringGenerator()

	return s.repo.Save(transaction)
}

func (s *transactionService) GetAllTransactions() ([]models.Transaction, error) {
	return s.repo.GetAll()
}

func (s *transactionService) DeleteAllTransactions() error {
	return s.repo.DeleteAll()
}

func validator(transaction models.Transaction) error {
	if transaction.Amount <= 0 {
		return fmt.Errorf("transaction amount should be more than 0")
	}
	return nil
}
