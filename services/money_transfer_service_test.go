package services_test

// type mockTRepository struct {
// 	Transactions []models.Transaction
// 	Err          error
// }

// func (m *mockTRepository) GetAll() ([]models.Transaction, error) {
// 	return m.Transactions, m.Err
// }

// func (m *mockTRepository) Save(transaction models.Transaction) error {
// 	return nil // Implement according to your test requirements
// }

// func (m *mockTRepository) DeleteAllTransactions() error {
// 	return nil // Implement according to your test requirements
// }

// func TestCreateNewMoneyTransfer_InsufficientBalance(t *testing.T) {
// 	mockRepo := &mockTRepository{
// 		Transactions: []models.Transaction{
// 			{Amount: 0.5},
// 			{Amount: 0.2},
// 		},
// 	}

// 	service := services.NewMoneyTransferService(mockRepo)

// 	err := service.CreateNewMoneyTransfer(5332123)
// 	assert.Error(t, err)
// 	assert.Equal(t, "insufficient balance. Unable to transfer money", err.Error())
// }

// func TestCreateNewMoneyTransfer_MinimumAmount(t *testing.T) {
// 	mockRepo := &mockTRepository{
// 		Transactions: []models.Transaction{
// 			{Amount: 0.5},
// 			{Amount: 0.2},
// 		},
// 	}

// 	service := services.NewMoneyTransferService(mockRepo)

// 	err := service.CreateNewMoneyTransfer(0.000009)
// 	assert.Error(t, err)
// 	assert.Equal(t, "amount too small. Minimum amount to transfer is 0.00001 BTC", err.Error())
// }

// func TestCreateNewMoneyTransfer_Success(t *testing.T) {
// 	mockRepo := &mockTRepository{
// 		Transactions: []models.Transaction{
// 			{Amount: 0.5},
// 			{Amount: 0.2},
// 		},
// 	}

// 	service := services.NewMoneyTransferService(mockRepo)

// 	err := service.CreateNewMoneyTransfer(10000)
// 	assert.NoError(t, err)
// }
