package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"
)

// MockHTTPClient is a mock implementation of HTTPClient for testing
type MockHTTPClient struct {
	Response *http.Response
	Err      error
}

func (m *MockHTTPClient) Post(url string, contentType string, body []byte) (*http.Response, error) {
	return m.Response, m.Err
}

func TestGetBlockNumber(t *testing.T) {
	// Mock response
	mockResponse := RPCResponse{
		Jsonrpc: "2.0",
		ID:      2,
		Result:  "0x123456",
	}
	responseBody, _ := json.Marshal(mockResponse)
	mockClient := &MockHTTPClient{
		Response: &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewBuffer(responseBody)),
		},
		Err: nil,
	}

	client := NewBlockchainClient(mockClient)
	blockNumber, err := client.GetBlockNumber()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if blockNumber != "0x123456" {
		t.Fatalf("Expected block number 0x123456, got %v", blockNumber)
	}
}

func TestGetBlockByNumber(t *testing.T) {
	// Mock response
	mockResponse := RPCResponse{
		Jsonrpc: "2.0",
		ID:      2,
		Result: map[string]interface{}{
			"number": "0x123456",
			"hash":   "0xabc123",
		},
	}
	responseBody, _ := json.Marshal(mockResponse)
	mockClient := &MockHTTPClient{
		Response: &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewBuffer(responseBody)),
		},
		Err: nil,
	}

	client := NewBlockchainClient(mockClient)
	block, err := client.GetBlockByNumber("0x123456")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if block["number"] != "0x123456" {
		t.Fatalf("Expected block number 0x123456, got %v", block["number"])
	}
}