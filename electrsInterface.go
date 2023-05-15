package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
)

const electrsURL = "127.0.0.1:50001"

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
