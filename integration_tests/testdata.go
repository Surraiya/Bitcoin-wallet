package integrationtests

import (
	"bitcoin-wallet/models"
	"bitcoin-wallet/utils"

	"github.com/jinzhu/gorm"
)

var mockData1 = []models.Transaction{
	{Id: utils.RandomUniqueStringGenerator(), Amount: 2, Spent: false},
	{Id: utils.RandomUniqueStringGenerator(), Amount: 3, Spent: false},
}

var mockData2 = []models.Transaction{}

var transferData1 = models.Transfer{
	AmountInEuro: 54326.45,
}

var transferData2 = models.Transfer{
	AmountInEuro: 0.5,
}

var transferData3 = models.Transfer{
	AmountInEuro: 6654326.45,
}

func createMockTransactions(db *gorm.DB, transactions []models.Transaction) []models.Transaction {
	for _, tx := range transactions {
		if err := db.Create(&tx).Error; err != nil {
			panic(err)
		}
	}
	return transactions
}
