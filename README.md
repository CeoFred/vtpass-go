# Go Library Starter Kit

Welcome to the Go Library Starter Kit! This starter kit is designed to help Go developers quickly start writing code to interact with third-party services. It includes a basic structure for service initialization, making HTTP requests, and handling responses. The package also includes a mock HTTP client for easy and effective testing.

## Features

- **Service Initialization**: Easily create service instances with API keys and authentication credentials.
- **HTTP Client**: Pre-configured HTTP client for making GET, POST, PUT, PATCH, and DELETE requests.
- **Response Handling**: Structured response handling for JSON data.
- **Mock Client**: Mock HTTP client for unit testing.

## Installation

To install the Go Library Starter Kit, run:

```sh
go get github.com/CeoFred/go_library_starter_kit
```

## Usage

### Service Initialization

Initialize a new service instance by providing your API key:

```go
package main

import (
	"context"
	"fmt"
	"github.com/CeoFred/go_library_starter_kit"
)

func main() {
	service := go_library_starter_kit.NewService("your-api-key")

	// Example usage
	response, err := service.GetData(context.Background())
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Balance:", response.Data.Balance)
}
```

### HTTP Client

The package provides a pre-configured HTTP client for making various types of HTTP requests. You can use the client directly or through the service methods.

### Response Handling

Responses are structured to handle JSON data with custom types:

```go
type BalanceResponse struct {
	BaseResponse
	Data struct {
		Balance  string `json:"balance"`
		Currency string `json:"currency"`
	} `json:"data"`
}

type BaseResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type ErrorResponse struct {
	BaseResponse
	Error string `json:"error"`
}
```

### Mock Client

For testing, a mock HTTP client is provided. This allows you to simulate HTTP responses without making actual network requests.

Example test for `GetData`:

```go
package go_library_starter_kit_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/CeoFred/go_library_starter_kit"
	"github.com/CeoFred/go_library_starter_kit/lib/httpclient"
	"github.com/stretchr/testify/assert"
)

func TestGetData(t *testing.T) {
	// Mock response
	mockResponse := go_library_starter_kit.BalanceResponse{
		BaseResponse: go_library_starter_kit.BaseResponse{
			Code:    "200",
			Message: "Success",
		},
		Data: struct {
			Balance  string `json:"balance"`
			Currency string `json:"currency"`
		}{
			Balance:  "100.00",
			Currency: "USD",
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
	service := &go_library_starter_kit.Service{
		apiKey:          "test-api-key",
		client:          mockClient,
		authCredentials: "?access_token=test-api-key",
	}

	// Call GetData
	resp, err := service.GetData(context.Background())

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "100.00", resp.Data.Balance)
}
```

## Contributing

Contributions are welcome! Please fork the repository and create a pull request with your changes.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Authors

- CeoFred

## Acknowledgements

- Thanks to the God almighty.

---

Happy coding!# go-library-starter-kit
