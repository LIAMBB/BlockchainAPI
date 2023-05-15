# GoLang Portfolio Project - HTTP Server for Electrs Interface

This repository contains a GoLang component that serves as an HTTP server acting as an interface to an Electrs server. Electrs is a popular open-source server that provides a fast and efficient index for the Bitcoin blockchain.

## Installation

1. Ensure you have Go installed on your system. You can download and install Go from the official website: [https://golang.org/](https://golang.org/)

2. Clone this repository to your local machine:

   ```bash
   git clone https://github.com/LIAMBB/BlockchainAPI.git

Install the required dependencies:

bash
Copy code
    
    go mod download
Configuration
Before running the server, make sure to configure the following parameters:

Electrs server address: Set the address of the Electrs server you want to connect to. Update the electrsURL constant in the electrsInterface.go file with the appropriate address.
Usage
To start the HTTP server, run the following command:

    go run .

By default, the server listens on port 8080. You can access it by opening http://localhost:8080 in your web browser.

API Endpoints
The HTTP server provides the following API endpoints:

GET /address/history: Retrieves the transaction history for a specific Bitcoin address.
GET /address/balance: Retrieves the balance for a specific Bitcoin address.
GET /address/unspent: Retrieves the list of unspent transaction outputs (UTXOs) for a specific Bitcoin address.
POST /transaction/broadcast: Broadcasts a Bitcoin transaction to the network.
GET /transaction/:id: Retrieves information about a specific transaction identified by its ID.
GET /block/header: Retrieves the header of the newest block in the Bitcoin blockchain.
Contributing
Contributions to this project are welcome. If you find any issues or want to add new features, please follow these steps:

Fork the repository.
Create a new branch.
Make your changes.
Test your changes.
Commit your changes.
Push the branch to your forked repository.
Open a pull request in this repository, describing your changes.
License
This project is licensed under the MIT License. Feel free to use, modify, and distribute the code as per the terms of the license.

Disclaimer
This software is provided as-is with no warranty or support. Use at your own risk. The authors of this project are not responsible for any damages or loss of funds resulting from the use of this software.

If you have any questions or need further assistance, please don't hesitate to contact us at your-email@example.com.