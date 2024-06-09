package go_library_starter_kit

import (
	"context"
	"net/http"
)

// HttpClient is the interface for the HTTP client.
type HttpClient interface {
	Get(ctx context.Context, path string) (*http.Response, error)
	Post(ctx context.Context, path string, payload interface{}) (*http.Response, error)
	Delete(ctx context.Context, path string, payload interface{}) (*http.Response, error)
	Put(ctx context.Context, path string, payload interface{}) (*http.Response, error)
	Patch(ctx context.Context, path string, payload interface{}) (*http.Response, error)
}
