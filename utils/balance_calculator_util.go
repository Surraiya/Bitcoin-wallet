package utils

import "bitcoin-wallet/models"

func CalculateCurrentBalance(transactions []models.Transaction) float64 {
	balance := 0.0

	for _, transaction := range transactions {
		if !transaction.Spent {
			balance += transaction.Amount
		}
	}

	return balance
}
