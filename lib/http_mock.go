package httpclient

import (
	"context"
	"net/http"
)

// MockClient is a mock HTTP client to replace the real one during tests.
type MockClient struct {
	GetFunc    func(ctx context.Context, path string) (*http.Response, error)
	PostFunc   func(ctx context.Context, path string, payload interface{}) (*http.Response, error)
	DeleteFunc func(ctx context.Context, path string, payload interface{}) (*http.Response, error)
	PutFunc    func(ctx context.Context, path string, payload interface{}) (*http.Response, error)
	PatchFunc  func(ctx context.Context, path string, payload interface{}) (*http.Response, error)
}

func (m *MockClient) Get(ctx context.Context, path string) (*http.Response, error) {
	return m.GetFunc(ctx, path)
}

func (m *MockClient) Post(ctx context.Context, path string, payload interface{}) (*http.Response, error) {
	return m.PostFunc(ctx, path, payload)
}

func (m *MockClient) Delete(ctx context.Context, path string, payload interface{}) (*http.Response, error) {
	return m.DeleteFunc(ctx, path, payload)
}

func (m *MockClient) Put(ctx context.Context, path string, payload interface{}) (*http.Response, error) {
	return m.PutFunc(ctx, path, payload)
}

func (m *MockClient) Patch(ctx context.Context, path string, payload interface{}) (*http.Response, error) {
	return m.PatchFunc(ctx, path, payload)
}

func NewMockClient() (*MockClient) {
	return &MockClient{}
}

// SetGetFunc sets the mock function for Get requests.
func (m *MockClient) SetGetFunc(f func(ctx context.Context, path string) (*http.Response, error)) {
	m.GetFunc = f
}

// SetPostFunc sets the mock function for Post requests.
func (m *MockClient) SetPostFunc(f func(ctx context.Context, path string, payload interface{}) (*http.Response, error)) {
	m.PostFunc = f
}

// SetDeleteFunc sets the mock function for Delete requests.
func (m *MockClient) SetDeleteFunc(f func(ctx context.Context, path string, payload interface{}) (*http.Response, error)) {
	m.DeleteFunc = f
}

// SetPutFunc sets the mock function for Put requests.
func (m *MockClient) SetPutFunc(f func(ctx context.Context, path string, payload interface{}) (*http.Response, error)) {
	m.PutFunc = f
}

// SetPatchFunc sets the mock function for Patch requests.
func (m *MockClient) SetPatchFunc(f func(ctx context.Context, path string, payload interface{}) (*http.Response, error)) {
	m.PatchFunc = f
}