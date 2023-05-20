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

type GetTxResponse struct {
	ID      int         `json:"id"`
	JSONRPC string      `json:"jsonrpc"`
	Result  Transaction `json:"result"`
}

type Transaction struct {
	ID            string   `json:"txid"`
	Hash          string   `json:"hash"`
	BlockHash     string   `json:"blockhash"`
	BlockTime     int64    `json:"blocktime"`
	Confirmations int      `json:"confirmations"`
	Size          int      `json:"size"`
	VirtualSize   int      `json:"vsize"`
	Weight        int      `json:"weight"`
	Version       int      `json:"version"`
	LockTime      int      `json:"locktime"`
	Time          int64    `json:"time"`
	InActiveChain bool     `json:"in_active_chain"`
	Hex           string   `json:"hex"`
	Vin           []Input  `json:"vin"`
	Vout          []Output `json:"vout"`
}

type Input struct {
	Coinbase    string   `json:"coinbase"`
	Sequence    uint32   `json:"sequence"`
	TxInWitness []string `json:"txinwitness"`
}

type Output struct {
	N            int     `json:"n"`
	Value        float64 `json:"value"`
	ScriptPubKey PubKey  `json:"scriptPubKey"`
}

type PubKey struct {
	Address string `json:"address"`
	Asm     string `json:"asm"`
	Desc    string `json:"desc"`
	Hex     string `json:"hex"`
	Type    string `json:"type"`
}

// Variable of all the block headers occuring since the the server starts up
var blocks []HeaderResult

// Main initializes the API endpoints and runs the API server on port 8080
func main() {
	blocks = make([]HeaderResult, 0)
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

	// GetTransaction endpoint
	router.GET("/transaction/get", getTransaction)

	// GetNewestBlockHeader endpoint
	router.GET("/block/header/new", getNewestBlockHeader)

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

	spew.Dump(requestBody)

	jsonResponse := Electrsinterface("blockchain.transaction.get", []interface{}{requestBody.TxID, true})
	var response GetTxResponse
	err := json.Unmarshal([]byte(jsonResponse), &response)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("========================================")
	fmt.Println(jsonResponse)
	fmt.Println("========================================")
	spew.Dump(response)
	fmt.Println("========================================")

	c.JSON(http.StatusOK, response.Result)
}

// GetNewestBlockHeader endpoint handler
func getNewestBlockHeader(c *gin.Context) {

	c.JSON(http.StatusOK, blocks[len(blocks)-1])
}
