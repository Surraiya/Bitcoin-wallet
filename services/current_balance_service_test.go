package services_test

// type mockTransactionRepository struct {
// 	Transactions []models.Transaction
// 	Err          error
// }

// func (m *mockTransactionRepository) GetAll() ([]models.Transaction, error) {
// 	return m.Transactions, m.Err
// }

// func (m *mockTransactionRepository) Save(transaction models.Transaction) error {
// 	return nil // Implement according to your test requirements
// }

// func (m *mockTRepository) DeleteAll() error {
// 	return m.Err //
// }

// func TestGetCurrentBalance_Success(t *testing.T) {
// 	mockRepo := &mockTransactionRepository{
// 		Transactions: []models.Transaction{
// 			{Amount: 0.1},
// 			{Amount: 0.2},
// 		},
// 	}

// 	//service := services.NewCurrentBalanceService(mockRepo)

// 	expectedBalance := 0.1 + 0.2
// 	expectedBalanceRounded := round(expectedBalance, 2) // Round to 2 decimal places

// 	balance, err := service.GetCurrentBalance()
// 	assert.NoError(t, err)
// 	assert.Equal(t, expectedBalanceRounded, round(balance.BTC, 2))
// }

// // round rounds a float64 number to the specified number of decimal places
// func round(num float64, places int) float64 {
// 	shift := math.Pow(10, float64(places))
// 	return math.Round(num*shift) / shift
// }
