package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	rpcURL = "https://polygon-rpc.com/"
)

// HTTPClient interface for mocking HTTP requests
type HTTPClient interface {
	Post(url string, contentType string, body []byte) (*http.Response, error)
}

// DefaultHTTPClient implements HTTPClient using the real http.Client
type DefaultHTTPClient struct{}

func (c *DefaultHTTPClient) Post(url string, contentType string, body []byte) (*http.Response, error) {
	return http.Post(url, contentType, bytes.NewBuffer(body))
}

type RPCRequest struct {
	Jsonrpc string        `json:"jsonrpc"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
	ID      int           `json:"id"`
}

type RPCResponse struct {
	Jsonrpc string      `json:"jsonrpc"`
	ID      int         `json:"id"`
	Result  interface{} `json:"result"`
}

type BlockchainClient struct {
	HTTPClient HTTPClient
}

func NewBlockchainClient(client HTTPClient) *BlockchainClient {
	return &BlockchainClient{HTTPClient: client}
}

func (bc *BlockchainClient) sendRPCRequest(method string, params []interface{}) ([]byte, error) {
	request := RPCRequest{
		Jsonrpc: "2.0",
		Method:  method,
		Params:  params,
		ID:      2,
	}

	requestBody, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	resp, err := bc.HTTPClient.Post(rpcURL, "application/json", requestBody)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (bc *BlockchainClient) GetBlockNumber() (string, error) {
	body, err := bc.sendRPCRequest("eth_blockNumber", nil)
	if err != nil {
		return "", err
	}

	var response RPCResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return "", err
	}

	return response.Result.(string), nil
}

func (bc *BlockchainClient) GetBlockByNumber(blockNumber string) (map[string]interface{}, error) {
	params := []interface{}{blockNumber, true}
	body, err := bc.sendRPCRequest("eth_getBlockByNumber", params)
	if err != nil {
		return nil, err
	}

	var response RPCResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, err
	}

	return response.Result.(map[string]interface{}), nil
}

func main() {
	client := NewBlockchainClient(&DefaultHTTPClient{})

	http.HandleFunc("/blockNumber", func(w http.ResponseWriter, r *http.Request) {
		blockNumber, err := client.GetBlockNumber()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"blockNumber": blockNumber})
	})

	http.HandleFunc("/block", func(w http.ResponseWriter, r *http.Request) {
		blockNumber := r.URL.Query().Get("number")
		if blockNumber == "" {
			http.Error(w, "Missing block number", http.StatusBadRequest)
			return
		}

		block, err := client.GetBlockByNumber(blockNumber)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(block)
	})

	log.Println("Starting server on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
