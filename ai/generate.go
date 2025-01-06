// Copyright 2016 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package ai

import (
	"fmt"
	"os"
	"net/url"
	"net/http"
	"strconv"

	"github.com/venusgalstar/go-ethereum/core/types"
	"github.com/venusgalstar/go-ethereum/core"
	"github.com/venusgalstar/go-ethereum/common"
)

func generate(tx *types.Transaction, msg *core.Message) {
	
	server := os.Getenv("AI_SERVER_IP")
	port := os.Getenv("AI_SERVER_PORT")	

	if port == "" {
		port = "3000" // Default value
	}
	if server == "" {
		server = "127.0.0.1"
	}

	url :=  "https://" + server + ":" + port + "/generate"

	data := map[string]string{
		"hash": tx.hash.Hex(),
		"from": msg.From.Hex(),
		"to": msg.To.Hex(),
		"nonce": strconv.FormatUint(msg.Nonce, 10),
		"value": strconv.FormatUint(msg.Value, 10),
		"data": string(msg.Data),
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("Error marshaling JSON: %v", err)
	}
	
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatalf("Error making POST request: %v", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}

	// Print the response
	fmt.Printf("Response Status: %s\n", resp.Status)
	fmt.Printf("Response Body: %s\n", string(body))
}
