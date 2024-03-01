package integrationtests

import (
	"bitcoin-wallet/database"
	"bitcoin-wallet/utils"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTransferMoney(t *testing.T) {
	db := database.SetupTestDB()
	defer database.CleanupTestDB(db)

	transactionsBeforeTransfer := createMockTransactions(db, mockData1)
	balanceBeforeTransfer := utils.CalculateCurrentBalance(transactionsBeforeTransfer)

	rr := CreateMoneyTransferRequest(transferData1)

	balanceAfterTransfer, _ := GetActualBalanceFromServer()
	transactionsAfterTransfer, _ := GetActualTransactionsFromServer()

	assert.Equal(t, http.StatusCreated, rr.Code, "handler returned wrong status code")
	assert.NotEqual(t, balanceBeforeTransfer, balanceAfterTransfer.BTC, "BTC balance before and after transfer should not be the same.")
	assert.GreaterOrEqualf(t, balanceBeforeTransfer, balanceAfterTransfer.BTC, "After transfer balance would be less")
	assert.NotEqual(t, transactionsBeforeTransfer[0].Spent, transactionsAfterTransfer[0].Spent, "After transfer transaction should update and spent should be true if used.")
	assert.NotEqual(t, len(transactionsBeforeTransfer), len(transactionsAfterTransfer), "Transactions before and after money transfer should not be the same as a new transaction of remaining should create")
}

func TestTransferMoney_AmountSmall(t *testing.T) {
	db := database.SetupTestDB()
	defer database.CleanupTestDB(db)

	transactionsBeforeTransfer := createMockTransactions(db, mockData1)
	balanceBeforeTransfer := utils.CalculateCurrentBalance(transactionsBeforeTransfer)

	rr := CreateMoneyTransferRequest(transferData2)

	balanceAfterTransfer, _ := GetActualBalanceFromServer()
	transactionsAfterTransfer, _ := GetActualTransactionsFromServer()

	assert.Equal(t, http.StatusBadRequest, rr.Code, "Transfer amount is less than minimum amount which is 0.00001 so api should reject the request.")
	assert.Contains(t, rr.Body.String(), "amount too small. Minimum amount to transfer is 0.00001 BTC", "Error should mention that amount too small")
	assert.Equal(t, balanceBeforeTransfer, balanceAfterTransfer.BTC, "BTC balance before and after transfer should be the same as transaction didnt happen.")
	assert.Equal(t, transactionsBeforeTransfer[0].Spent, transactionsAfterTransfer[0].Spent, "Spent will remain false")
	assert.Equal(t, len(transactionsBeforeTransfer), len(transactionsAfterTransfer), "Transactions before and after money transfer should be the same.")
}

func TestTransferMoney_InsufficientBalance(t *testing.T) {
	db := database.SetupTestDB()
	defer database.CleanupTestDB(db)

	transactionsBeforeTransfer := createMockTransactions(db, mockData1)
	balanceBeforeTransfer := utils.CalculateCurrentBalance(transactionsBeforeTransfer)

	rr := CreateMoneyTransferRequest(transferData3)

	balanceAfterTransfer, _ := GetActualBalanceFromServer()
	transactionsAfterTransfer, _ := GetActualTransactionsFromServer()

	assert.Equal(t, http.StatusBadRequest, rr.Code, "Transfer amount is greater than balance so api should reject the request.")
	assert.Contains(t, rr.Body.String(), "insufficient balance. Unable to transfer money", "Error should mention insifficient balance")
	assert.Equal(t, balanceBeforeTransfer, balanceAfterTransfer.BTC, "BTC balance before and after transfer should be the same as transaction didnt happen.")
	assert.Equal(t, transactionsBeforeTransfer[0].Spent, transactionsAfterTransfer[0].Spent, "Spent will remain false")
	assert.Equal(t, len(transactionsBeforeTransfer), len(transactionsAfterTransfer), "Transactions before and after money transfer should be the same.")
}
