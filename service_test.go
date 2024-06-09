package go_library_starter_kit

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/CeoFred/go_library_starter_kit/lib"

	"github.com/stretchr/testify/assert"
)

// MockClient is a mock HTTP client to replace the real one during tests.

func TestGetData(t *testing.T) {
	// Mock response
	mockResponse := UserDataResponse{
		BaseResponse: BaseResponse{
			Code:    "200",
			Message: "Success",
		},
		User: struct {
			FirstName string `json:"first_name"`
		}{
			FirstName: "John",
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
	service := &Service{
		apiKey:          "test-api-key",
		client:          mockClient,
		authCredentials: "?access_token=test-api-key",
	}

	// Call GetData
	resp, err := service.GetData(context.Background())

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "John", resp.User.FirstName)
}

func TestPostData(t *testing.T) {
	// Mock response
	mockResponse := Response{
		BaseResponse: BaseResponse{
			Code:    "200",
			Message: "Success",
		},
		Data: struct {
			User string `json:"user"`
		}{
			User: "test-user",
		},
	}

	// Convert mock response to JSON
	mockResponseBody, _ := json.Marshal(mockResponse)

	// Setup mock client
	mockClient := httpclient.NewMockClient()
	mockClient.SetPostFunc(func(ctx context.Context, path string, payload interface{}) (*http.Response, error) {
		rec := httptest.NewRecorder()
		rec.WriteHeader(http.StatusOK)
		rec.Write(mockResponseBody)
		return rec.Result(), nil
	})

	// Initialize service with mock client
	service := &Service{
		apiKey:          "test-api-key",
		client:          mockClient,
		authCredentials: "?access_token=test-api-key",
	}

	// Payload for PostData
	payload := map[string]interface{}{
		"first_name": "John",
	}

	// Call PostData
	resp, err := service.PostData(context.Background(), payload)

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "test-user", resp.Data.User)
}
