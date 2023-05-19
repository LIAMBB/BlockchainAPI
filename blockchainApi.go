package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/checksum0/go-electrum/electrum"
	"github.com/davecgh/go-spew/spew"
	"github.com/gin-gonic/gin"
)

type History struct {
	Height int    `json:"height"`
	TxHash string `json:"tx_hash"`
}

type HistoryResponse struct {
	ID      int       `json:"id"`
	JSONRPC string    `json:"jsonrpc"`
	Result  []History `json:"result"`
}

type UnspentResponse struct {
	ID      int           `json:"id"`
	JSONRPC string        `json:"jsonrpc"`
	Result  []ListUnspent `json:"result"`
}

type ListUnspent struct {
	Height int    `json:"height"`
	TxHash string `json:"tx_hash"`
	TxPos  int    `json:"tx_pos"`
	Value  int    `json:"value"`
}

type BalanceResponse struct {
	ID      int     `json:"id"`
	JSONRPC string  `json:"jsonrpc"`
	Result  Balance `json:"result"`
}

type Balance struct {
	Confirmed   int64 `json:"confirmed"`
	Unconfirmed int64 `json:"unconfirmed"`
}

type TxBroadcastResponse struct {
	ID      int    `json:"id"`
	JSONRPC string `json:"jsonrpc"`
	Result  string `json:"result"`
}

type AddrRequestBody struct {
	Address string `json:"address"`
}

type TxBroadcastRequestBody struct {
	RawTransaction string `json:"rawTx"`
}

type GetTxRequestBody struct {
	TxID string `json:"txid"`
}

var blocks []BlockData

func main() {
	blocks = make([]BlockData, 0)
	go BlockWatcher(&blocks)
	router := gin.Default()

	// AddressHistory endpoint
	router.GET("/address/history", getAddressHistory)

	// AddressBalance endpoint
	router.GET("/address/balance", getAddressBalance)

	// AddressListUnspent endpoint
	router.GET("/address/unspent", getAddressListUnspent)

	// TransactionBroadcast endpoint
	router.POST("/transaction/broadcast", postTransactionBroadcast)

	// // GetTransaction endpoint
	// router.GET("/transaction/:id", getTransaction)

	// GetNewestBlockHeader endpoint
	router.GET("/block/header", getNewestBlockHeader)

	// Run the server
	router.Run(":8080")
}

// AddressHistory endpoint handler
func getAddressHistory(c *gin.Context) {
	// Get address from query parameter
	var requestBody AddrRequestBody

	// Bind the JSON content to the RequestBody struct
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// TODO: Implement AddressBalance logic
	scriphash, _ := electrum.AddressToElectrumScriptHash(requestBody.Address)
	jsonResponse := Electrsinterface("blockchain.scripthash.get_history", []interface{}{scriphash})
	fmt.Println("=========================================")

	var response HistoryResponse
	err := json.Unmarshal([]byte(jsonResponse), &response)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	c.JSON(http.StatusOK, response.Result)
}

// AddressBalance endpoint handler
func getAddressBalance(c *gin.Context) {
	// Get address from query parameter
	var requestBody AddrRequestBody

	// Bind the JSON content to the RequestBody struct
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Converts Address to Electrum ScriptHash then makes a request through the ElectrsInterface
	scriphash, _ := electrum.AddressToElectrumScriptHash(requestBody.Address)
	jsonResponse := Electrsinterface("blockchain.scripthash.get_balance", []interface{}{scriphash})

	var response BalanceResponse
	err := json.Unmarshal([]byte(jsonResponse), &response)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	c.JSON(http.StatusOK, response.Result)
}

// AddressListUnspent endpoint handler
func getAddressListUnspent(c *gin.Context) {
	// Get address from query parameter
	var requestBody AddrRequestBody

	// Bind the JSON content to the RequestBody struct
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// TODO: Implement AddressBalance logic
	scriphash, _ := electrum.AddressToElectrumScriptHash(requestBody.Address)
	jsonResponse := Electrsinterface("blockchain.scripthash.listunspent", []interface{}{scriphash})

	var response UnspentResponse
	err := json.Unmarshal([]byte(jsonResponse), &response)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	c.JSON(http.StatusOK, response.Result)
}

// TransactionBroadcast endpoint handler
func postTransactionBroadcast(c *gin.Context) {
	// Get address from query parameter
	var requestBody TxBroadcastRequestBody

	// Bind the JSON content to the RequestBody struct
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	jsonResponse := Electrsinterface("blockchain.transaction.broadcast", []interface{}{requestBody.RawTransaction})
	fmt.Println("======================================================")
	spew.Dump(jsonResponse)
	fmt.Println("======================================================")

	var response TxBroadcastResponse
	err := json.Unmarshal([]byte(jsonResponse), &response)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	c.JSON(http.StatusOK, response.Result)

}

// GetTransaction endpoint handler
func getTransaction(c *gin.Context) {
	// TODO: Implement GetTransaction logic
	// Get TxID from query parameter
	var requestBody GetTxRequestBody

	// Bind the JSON content to the RequestBody struct
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// TODO: Implement AddressBalance logic
	jsonResponse := Electrsinterface("blockchain.transaction.get", []interface{}{requestBody.TxID, true})

	var response UnspentResponse
	err := json.Unmarshal([]byte(jsonResponse), &response)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	c.JSON(http.StatusOK, response.Result)
}

// GetNewestBlockHeader endpoint handler
func getNewestBlockHeader(c *gin.Context) {
	// TODO: Implement GetNewestBlockHeader logic
	fmt.Println(blocks)
	c.JSON(http.StatusOK, blocks[:len(blocks)-1])

}
