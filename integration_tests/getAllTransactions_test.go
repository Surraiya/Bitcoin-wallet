package integrationtests

import (
	"bitcoin-wallet/database"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllTransactions(t *testing.T) {
	db := database.SetupTestDB()
	defer database.CleanupTestDB(db)
	expectedTransactions := createMockTransactions(db, mockData1)

	// Create a request to pass controller
	actualTransactions, rr := GetActualTransactionsFromServer()

	//Assertions
	assert.Equal(t, http.StatusOK, rr.Code, "handler returned wrong status code")
	assert.Contains(t, rr.Header().Get("Content-Type"), "application/json", "handler returned unexpected content type")
	assert.NotEmpty(t, rr.Body.String(), "handler returned empty body")
	assert.Equal(t, len(expectedTransactions), len(actualTransactions), "Length should be equal")
	assert.Equal(t, expectedTransactions[0].Id, actualTransactions[0].Id, "Both expected and actual transaction's first id should be the same")
	assert.Equal(t, expectedTransactions[0].Amount, actualTransactions[0].Amount, "Both expected and actual transaction's first Amount should be the same")
	assert.Equal(t, expectedTransactions[0].Spent, actualTransactions[0].Spent, "Both expected and actual transaction's first Spent should be the same")
}
