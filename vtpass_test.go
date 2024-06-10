package vtupass_go

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	httpclient "github.com/CeoFred/vtpass_go/lib"

	"github.com/stretchr/testify/assert"
)

func TestPing(t *testing.T) {
	// Mock response
	mockResponse := WalletBalance{
		BaseResponse: BaseResponse{
			Code: "020",
		},
	}

	// Convert mock response to JSON
	mockResponseBody, _ := json.Marshal(mockResponse)

	// Setup mock client
	mockClient := httpclient.NewMockClient()

	mockClient.SetGetFunc(func(ctx context.Context, path string) (*http.Response, error) {
		rec := httptest.NewRecorder()
		rec.WriteHeader(http.StatusOK)
		rec.Write(mockResponseBody)
		return rec.Result(), nil
	})

	// Initialize service with mock client
	service := &VTService{
		apiKey: "test-api-key",
		client: mockClient,
		authCredentials: map[string]string{
			"api-key": "test-api-key",
		},
	}

	// Call Balance
	resp, err := service.Ping(context.Background())

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

// func TestPostData(t *testing.T) {
// 	// Mock response
// 	mockResponse := Response{
// 		BaseResponse: BaseResponse{
// 			Code:    020,
// 		},
// 		Data: struct {
// 			User string `json:"user"`
// 		}{
// 			User: "test-user",
// 		},
// 	}

// 	// Convert mock response to JSON
// 	mockResponseBody, _ := json.Marshal(mockResponse)

// 	// Setup mock client
// 	mockClient := httpclient.NewMockClient()
// 	mockClient.SetPostFunc(func(ctx context.Context, path string, payload interface{}) (*http.Response, error) {
// 		rec := httptest.NewRecorder()
// 		rec.WriteHeader(http.StatusOK)
// 		rec.Write(mockResponseBody)
// 		return rec.Result(), nil
// 	})

// 	// Initialize service with mock client
// 	service := &Service{
// 		apiKey:          "test-api-key",
// 		client:          mockClient,
// 	authCredentials: map[string]string{
// 			"api-key": "test-api-key",
// 		},
// 	}

// 	// Payload for PostData
// 	payload := map[string]interface{}{
// 		"first_name": "John",
// 	}

// 	// Call PostData
// 	resp, err := service.PostData(context.Background(), payload)

// 	// Assertions
// 	assert.NoError(t, err)
// 	assert.NotNil(t, resp)
// 	assert.Equal(t, "test-user", resp.Data.User)
// }
