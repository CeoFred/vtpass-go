package vtupass_go

import (
	"context"
	"net/http"
)

// HttpClient is the interface for the HTTP client.
type HttpClient interface {
	Get(ctx context.Context, path string,headers ...map[string]string) (*http.Response, error)
	Post(ctx context.Context, path string, payload interface{},headers ...map[string]string) (*http.Response, error)
	Delete(ctx context.Context, path string, payload interface{},headers ...map[string]string) (*http.Response, error)
	Put(ctx context.Context, path string, payload interface{},headers ...map[string]string) (*http.Response, error)
	Patch(ctx context.Context, path string, payload interface{},headers ...map[string]string) (*http.Response, error)
}
