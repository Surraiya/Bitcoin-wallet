package integrationtests

import (
	"bitcoin-wallet/database"
	"bitcoin-wallet/utils"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShowCurrentBalance(t *testing.T) {
	db := database.SetupTestDB()
	defer database.CleanupTestDB(db)
	expectedTransactions := createMockTransactions(db, mockData1)

	expectedBalanceBTC := utils.CalculateCurrentBalance(expectedTransactions)
	rate, err := utils.GetExchangeRate()
	if err != nil {
		t.Fatal(err)
	}
	expectedBalanceEuro := expectedBalanceBTC * rate

	actualBalance, rr := GetActualBalanceFromServer()

	tolerance := 0.001
	assert.Equal(t, http.StatusOK, rr.Code, "handler returned wrong status code")
	assert.Equal(t, expectedBalanceBTC, actualBalance.BTC, "Balance in BTC should be equal")
	assert.InDelta(t, expectedBalanceEuro, actualBalance.Euro, tolerance, "Euro balance should be almost equal to the expected value with a tolerance")
}

func TestShowCurrentBalance_Null(t *testing.T) {
	db := database.SetupTestDB()
	defer database.CleanupTestDB(db)
	expectedTransactions := createMockTransactions(db, mockData2)

	expectedBalanceBTC := utils.CalculateCurrentBalance(expectedTransactions)
	rate, err := utils.GetExchangeRate()
	if err != nil {
		t.Fatal(err)
	}
	expectedBalanceEuro := expectedBalanceBTC * rate

	actualBalance, rr := GetActualBalanceFromServer()

	tolerance := 0.001
	assert.Equal(t, http.StatusOK, rr.Code, "handler returned wrong status code")
	assert.Equal(t, expectedBalanceBTC, actualBalance.BTC, "Balance in BTC should be equal")
	assert.InDelta(t, expectedBalanceEuro, actualBalance.Euro, tolerance, "Euro balance should be almost equal to the expected value with a tolerance")
}
