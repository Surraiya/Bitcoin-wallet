package utils_test

import (
	"bitcoin-wallet/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetExchangeRate_RealAPI(t *testing.T) {
	// This test accesses the real API, so it may fail due to network issues or changes in the API response.

	// Call the function to get the exchange rate
	rate, err := utils.GetExchangeRate()

	// Check if there was an error accessing the API
	if err != nil {
		t.Skipf("Skipping test: unable to access the API. Error: %v", err)
	}

	// Assert that the rate is a positive number
	assert.Greater(t, rate, 0.0)
}
