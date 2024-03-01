package services

import (
	"bitcoin-wallet/models"
	"bitcoin-wallet/repositories"
	"bitcoin-wallet/utils"
	"fmt"
)

const MinimumAmount = 0.00001

type MoneyTransferService interface {
	CreateNewMoneyTransfer(amount float64) error
}

type moneyTransferService struct {
	repo repositories.TransactionRepository
}

func NewMoneyTransferService(repo repositories.TransactionRepository) *moneyTransferService {
	return &moneyTransferService{repo: repo}
}

func (m *moneyTransferService) CreateNewMoneyTransfer(amountEuro float64) error {
	balanceBTC := m.getCurrentBalance()
	transferAmount := m.convertEuroToBtc(amountEuro)

	if err := m.ValidateTransferAmount(balanceBTC, transferAmount); err != nil {
		return err
	}

	unspentTransactions, err := m.repo.GetAllUnspentTransactions()
	if err != nil {
		return err
	}

	if err := m.processTransactions(unspentTransactions, transferAmount); err != nil {
		return err
	}

	return nil
}

func (m *moneyTransferService) processTransactions(transactions []models.Transaction, transferAmount float64) error {
	remainingAmount := 0.0

	for _, transaction := range transactions {

		if transaction.Amount >= transferAmount {

			transaction.Spent = true
			remainingAmount = transaction.Amount - transferAmount

			if err := m.updateTransaction(transaction); err != nil {
				return err
			}

			break
		}

		transaction.Spent = true
		transferAmount -= transaction.Amount

		if err := m.updateTransaction(transaction); err != nil {
			return err
		}
	}

	if remainingAmount >= 0 {
		if err := m.createTransaction(remainingAmount, false); err != nil {
			return err
		}
	}
	return nil
}

func (m *moneyTransferService) ValidateTransferAmount(balanceBTC float64, amountBTC float64) error {
	if amountBTC > balanceBTC {
		return fmt.Errorf("insufficient balance. Unable to transfer money")
	}

	if amountBTC < MinimumAmount {
		return fmt.Errorf("amount too small. Minimum amount to transfer is 0.00001 BTC")
	}

	return nil
}

func (m *moneyTransferService) createTransaction(amount float64, spent bool) error {
	transaction := models.Transaction{
		Id:     utils.RandomUniqueStringGenerator(),
		Amount: amount,
		Spent:  spent,
	}

	if err := m.repo.Save(transaction); err != nil {
		return err
	}

	return nil
}

func (m *moneyTransferService) updateTransaction(transaction models.Transaction) error {
	return m.repo.Update(transaction)
}

func (m *moneyTransferService) getCurrentBalance() float64 {
	transactions, err := m.repo.GetAll()
	if err != nil {
		fmt.Println(err)
	}

	return utils.CalculateCurrentBalance(transactions)
}

func (m *moneyTransferService) convertEuroToBtc(amount float64) float64 {
	rate, err := utils.GetExchangeRate()
	if err != nil {
		fmt.Println("Error getting the exchange rate")
	}
	return amount / rate
}
