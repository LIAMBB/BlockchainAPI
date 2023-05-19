package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
)

const electrsURL = "127.0.0.1:50001"

type HeaderSubscription struct {
	ID     int          `json:"id"`
	Result HeaderResult `json:"result"`
}

type HeaderResult struct {
	Height int    `json:"height"`
	Hex    string `json:"hex"`
}

func Electrsinterface(method string, params []interface{}) string {
	// Establish a TCP connection to the server
	conn, err := net.Dial("tcp", electrsURL)
	if err != nil {
		return ""
	}
	defer conn.Close()
	data := map[string]interface{}{
		"method": method,
		"params": params,
		"id":     0,
	}

	payload, err := json.Marshal(data)
	if err != nil {
		return ""
	}
	// Send the request message to the server
	_, err = fmt.Fprintf(conn, "%s\n", payload)
	if err != nil {
		return ""
	}

	// Read the response message from the server
	responseBytes, err := bufio.NewReader(conn).ReadBytes('\n')
	if err != nil {
		return ""
	}
	return string(responseBytes)
}

func BlockWatcher(blocks *[]HeaderResult) {
	conn, err := net.Dial("tcp", electrsURL)
	if err != nil {
		//Error Out TODO
	}
	defer conn.Close()
	data := map[string]interface{}{
		"method": "blockchain.headers.subscribe",
		"params": []interface{}{},
		"id":     0,
	}
	payload, err := json.Marshal(data)
	if err != nil {
		//Error Handling TODO
	}
	// Send the request message to the server
	_, err = fmt.Fprintf(conn, "%s\n", payload)
	if err != nil {
		//Error Handling TODO
	}

	for {
		responseBytes, err := bufio.NewReader(conn).ReadBytes('\n')
		if err != nil {
			//Error Handling TODO
		}

		fmt.Println("===========================================")
		fmt.Println(string(responseBytes))
		fmt.Println("===========================================")

		var response HeaderSubscription
		err = json.Unmarshal(responseBytes, &response)
		if err != nil {
			// Error Handling TODO
		}
		temp := *blocks
		temp = append(temp, response.Result)
		*blocks = temp
	}

}
