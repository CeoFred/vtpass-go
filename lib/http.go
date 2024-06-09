package httpclient

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	utils "github.com/CeoFred/vtupass_go/utils"
)

var (
	// defaultClient is the default HTTP client for the package.
	defaultClient = &http.Client{
		Timeout: time.Duration(15) * time.Second,
		Transport: &http.Transport{
			MaxIdleConns:        10,
			MaxIdleConnsPerHost: 100,
			IdleConnTimeout:     30 * time.Second,
		},
	}
)

// APIClient is a wrapper for making HTTP requests to the API.
type APIClient struct {
	baseURL string
	apiKey  string
	client  *http.Client
}

// NewAPIClient creates a new instance of APIClient.
func NewAPIClient(baseURL, apiKey string) *APIClient {
	return &APIClient{
		baseURL: baseURL,
		apiKey:  apiKey,
		client:  defaultClient,
	}
}

// Helper function to convert variadic headers to a map


func (c *APIClient) Put(ctx context.Context, endpoint string, payload interface{}, headers ...map[string]string) (*http.Response, error) {
	url := c.baseURL + endpoint

	// Convert payload to JSON
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPut, url, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	for key, value := range utils.HeadersToMap(headers...) {
		req.Header.Set(key, value)
	}

	return c.client.Do(req)
}

func (c *APIClient) Patch(ctx context.Context, endpoint string, payload interface{}, headers ...map[string]string) (*http.Response, error) {
	url := c.baseURL + endpoint

	// Convert payload to JSON
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPatch, url, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	for key, value := range utils.HeadersToMap(headers...) {
		req.Header.Set(key, value)
	}

	return c.client.Do(req)
}

// Post sends a POST request to the specified endpoint with the given payload.
func (c *APIClient) Post(ctx context.Context, endpoint string, payload interface{}, headers ...map[string]string) (*http.Response, error) {
	url := c.baseURL + endpoint

	// Convert payload to JSON
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	for key, value := range utils.HeadersToMap(headers...) {
		req.Header.Set(key, value)
	}

	return c.client.Do(req)
}

func (c *APIClient) Delete(ctx context.Context, endpoint string, payload interface{}, headers ...map[string]string) (*http.Response, error) {
	url := c.baseURL + endpoint

	// Convert payload to JSON
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, url, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	for key, value := range utils.HeadersToMap(headers...) {
		req.Header.Set(key, value)
	}

	return c.client.Do(req)
}

// Get sends a GET request to the specified endpoint, appending id as a path parameter
func (c *APIClient) Get(ctx context.Context, path string, headers ...map[string]string) (*http.Response, error) {
	url := fmt.Sprintf("%s%s", c.baseURL, path)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	for key, value := range utils.HeadersToMap(headers...) {
		req.Header.Set(key, value)
	}

	return c.client.Do(req)
}
