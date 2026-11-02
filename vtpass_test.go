package vtupass_go

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	httpclient "github.com/CeoFred/vtpass-go/lib"

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

func TestPurchaseAirtime(t *testing.T) {
	mockResponse := PayResponse{
		Code: "000",
		ResponseDescription: "TRANSACTION SUCCESSFUL",
		RequestID: "req-123",
	}

	mockResponseBody, _ := json.Marshal(mockResponse)

	mockClient := httpclient.NewMockClient()
	mockClient.SetPostFunc(func(ctx context.Context, path string, payload interface{}) (*http.Response, error) {
		assert.Equal(t, "pay", path)
		rec := httptest.NewRecorder()
		rec.WriteHeader(http.StatusOK)
		rec.Write(mockResponseBody)
		return rec.Result(), nil
	})

	service := &VTService{
		apiKey: "test-api-key",
		client: mockClient,
	}

	resp, err := service.PurchaseAirtime(context.Background(), AirtimePurchase{
		RequestID: "req-123",
		ServiceID: "mtn",
		Amount:    100,
		Phone:     "08011111111",
	})

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "000", resp.Code)
}

func TestPurchaseData(t *testing.T) {
	mockResponse := PayResponse{
		Code: "000",
		ResponseDescription: "TRANSACTION SUCCESSFUL",
		RequestID: "req-456",
	}

	mockResponseBody, _ := json.Marshal(mockResponse)

	mockClient := httpclient.NewMockClient()
	mockClient.SetPostFunc(func(ctx context.Context, path string, payload interface{}) (*http.Response, error) {
		assert.Equal(t, "pay", path)
		rec := httptest.NewRecorder()
		rec.WriteHeader(http.StatusOK)
		rec.Write(mockResponseBody)
		return rec.Result(), nil
	})

	service := &VTService{
		apiKey: "test-api-key",
		client: mockClient,
	}

	resp, err := service.PurchaseData(context.Background(), DataPurchase{
		RequestID:     "req-456",
		ServiceID:     "mtn-data",
		BillersCode:   "08011111111",
		VariationCode: "mtn-10mb-100",
		Amount:        100,
		Phone:         "08011111111",
	})

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "000", resp.Code)
}

func TestVerifySmartCard(t *testing.T) {
	mockResponse := TVCustomerInfoResponse{
		Code: "000",
		Content: TVCustomerInfo{
			CustomerName:   "TEST CUSTOMER",
			Status:         "ACTIVE",
			DueDate:        "2025-02-06T00:00:00",
			CustomerNumber: "8061522780",
			CustomerType:   "DSTV",
			RenewalAmount:  1850,
		},
	}

	mockResponseBody, _ := json.Marshal(mockResponse)

	mockClient := httpclient.NewMockClient()
	mockClient.SetPostFunc(func(ctx context.Context, path string, payload interface{}) (*http.Response, error) {
		assert.Equal(t, "merchant-verify", path)
		rec := httptest.NewRecorder()
		rec.WriteHeader(http.StatusOK)
		rec.Write(mockResponseBody)
		return rec.Result(), nil
	})

	service := &VTService{
		apiKey: "test-api-key",
		client: mockClient,
	}

	resp, err := service.VerifySmartCard(context.Background(), "1212121212", "dstv")

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "TEST CUSTOMER", resp.CustomerName)
	assert.Equal(t, "ACTIVE", resp.Status)
}

func TestPurchaseTV(t *testing.T) {
	mockResponse := PayResponse{
		Code: "000",
		ResponseDescription: "TRANSACTION SUCCESSFUL",
		RequestID: "req-789",
	}

	mockResponseBody, _ := json.Marshal(mockResponse)

	mockClient := httpclient.NewMockClient()
	mockClient.SetPostFunc(func(ctx context.Context, path string, payload interface{}) (*http.Response, error) {
		assert.Equal(t, "pay", path)
		rec := httptest.NewRecorder()
		rec.WriteHeader(http.StatusOK)
		rec.Write(mockResponseBody)
		return rec.Result(), nil
	})

	service := &VTService{
		apiKey: "test-api-key",
		client: mockClient,
	}

	resp, err := service.PurchaseTV(context.Background(), TVPurchase{
		RequestID:        "req-789",
		ServiceID:        "dstv",
		BillersCode:      "1212121212",
		VariationCode:    "dstv-padi",
		Amount:           1850,
		Phone:            "08011111111",
		SubscriptionType: "change",
	})

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "000", resp.Code)
}

func TestCustomerInfoUnmarshalNumericFields(t *testing.T) {
	var c CustomerInfo
	err := json.Unmarshal([]byte(`{"Customer_Name":"A","Customer_Arrears":0,"Last_Purchase_Days":"12","KCT1":null,"Customer_Number":123}`), &c)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if c.CustomerArrears != "0" || c.LastPurchaseDays != "12" || c.CustomerNumber != "123" || c.KCT1 != "" || c.CustomerName != "A" {
		t.Fatalf("unexpected result: %+v", c)
	}
}
