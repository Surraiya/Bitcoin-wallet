package services

import (
	"bitcoin-wallet/models"
	"bitcoin-wallet/repositories"
	"bitcoin-wallet/utils"
	"fmt"
)

type CurrentBalanceService interface {
	GetCurrentBalance() (models.Balance, error)
}

type currentBalanceService struct {
	repo repositories.TransactionRepository
}

func NewCurrentBalanceService(repo repositories.TransactionRepository) *currentBalanceService {
	return &currentBalanceService{repo: repo}
}

func (b *currentBalanceService) GetCurrentBalance() (models.Balance, error) {
	//get all the transaction
	transactions, err := b.repo.GetAll()
	if err != nil {
		return models.Balance{}, fmt.Errorf("error getting all the transactions")
	}

	btc := utils.CalculateCurrentBalance(transactions)
	rate, err := utils.GetExchangeRate()
	euro := btc * rate
	if err != nil {
		return models.Balance{}, fmt.Errorf("error in getting the exchange rate")
	}
	return models.Balance{
		BTC:  btc,
		Euro: euro,
	}, err
}
