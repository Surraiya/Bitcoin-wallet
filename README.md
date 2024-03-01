# Bitcoin Wallet Application

This is a simple Bitcoin wallet application written in Go.

## Features

- Allows users to perform various operations related to Bitcoin transactions.
- Provides a RESTful API for interacting with the wallet.
- Supports creating new transactions, transferring money, and checking balances.

## Getting Started

To get started with the Bitcoin Wallet application, follow these steps:

1. Clone the repository:
   ```bash
   git clone https://github.com/Surraiya/Bitcoin-wallet.git
   ```
   
2. Build the Docker image:
    ```bash
    docker build -t bitcoin-wallet .
    ```
    
3. Run the Docker container:
   ```bash
   docker run -p 8083:8083 bitcoin-wallet
   ```
   
4. Access the application in your web browser at http://localhost:8083

## API Endpoints

- **GET /transactions**: Retrieve all transactions.
- **POST /transactions**: Create a new transaction.
- **POST /money-transfers**: Transfer money from one account to another.
- **GET /current-balance**: Get the current balance.


## Dependencies

- Go: The programming language used to develop the application.
- Docker: Containerization platform used to build, ship, and run the application in containers.
- Gin Framework: Web framework used to handle HTTP requests and responses in the application.
- GORM: Object-Relational Mapping (ORM) library used for database interactions in the application.

   
