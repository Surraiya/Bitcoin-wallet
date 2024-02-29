package utils

import (
	"bitcoin-wallet/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

var data struct {
	Data []models.ExchangeRate `json:"data"`
}

var ExchangeRateAPIURL = "http://api-cryptopia.adca.sh/v1/prices/ticker"

func GetExchangeRate() (float64, error) {

	response, err := http.Get(ExchangeRateAPIURL)
	if err != nil {
		return 0, err
	}
	defer response.Body.Close()

	if err := json.NewDecoder(response.Body).Decode(&data); err != nil {
		return 0, err
	}

	for _, rate := range data.Data {
		if rate.Symbol == "BTC/EUR" {
			rate, err := strconv.ParseFloat(rate.Value, 64)
			if err != nil {
				return 0, err
			}
			return rate, nil
		}
	}

	return 0, fmt.Errorf("BTC/EUR rate not found in API response")
}
