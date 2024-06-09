package utils

import (
	"reflect"
	"testing"
)

func TestMergeHeaders(t *testing.T) {
	// Test cases
	testCases := []struct {
		name        string
		baseHeaders map[string]string
		headers     map[string]string
		expected    map[string]string
	}{
		{
			name:        "Basic Test",
			baseHeaders: map[string]string{"Content-Type": "application/json"},
			headers:     map[string]string{"Authorization": "Bearer token"},
			expected:    map[string]string{"Content-Type": "application/json", "Authorization": "Bearer token"},
		},
		{
			name:        "Overlapping Keys",
			baseHeaders: map[string]string{"Content-Type": "application/json"},
			headers:     map[string]string{"Content-Type": "text/html"},
			expected:    map[string]string{"Content-Type": "text/html"},
		},
		{
			name:        "Empty Inputs",
			baseHeaders: map[string]string{},
			headers:     map[string]string{},
			expected:    map[string]string{},
		},
		{
			name:        "Empty baseHeaders",
			baseHeaders: map[string]string{},
			headers:     map[string]string{"Authorization": "Bearer token"},
			expected:    map[string]string{"Authorization": "Bearer token"},
		},
		{
			name:        "Empty headers",
			baseHeaders: map[string]string{"Content-Type": "application/json"},
			headers:     map[string]string{},
			expected:    map[string]string{"Content-Type": "application/json"},
		},
		{
			name:        "Nil Inputs",
			baseHeaders: nil,
			headers:     nil,
			expected:    map[string]string{},
		},
	}

	// Run tests
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := MergeHeaders(tc.baseHeaders, tc.headers)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected: %v, but got: %v", tc.expected, result)
			}
		})
	}
}

func TestHeadersToMap(t *testing.T) {
	// Test cases
	testCases := []struct {
		name     string
		headers  []map[string]string
		expected map[string]string
	}{
		{
			name:     "Basic Test",
			headers:  []map[string]string{{"Content-Type": "application/json"}, {"Authorization": "Bearer token"}, {"Cache-Control": "no-cache"}},
			expected: map[string]string{"Content-Type": "application/json", "Authorization": "Bearer token", "Cache-Control": "no-cache"},
		},
		{
			name:     "Empty Input",
			headers:  []map[string]string{{}},
			expected: map[string]string{},
		},
		{
			name:     "Empty Maps in Inputs",
			headers:  []map[string]string{{"Content-Type": "application/json"}, {}, {"Authorization": "Bearer token"}},
			expected: map[string]string{"Content-Type": "application/json", "Authorization": "Bearer token"},
		},
		{
			name:     "Nil Inputs",
			headers:  nil,
			expected: map[string]string{},
		},
		{
			name:     "Duplicate Keys",
			headers:  []map[string]string{{"Content-Type": "application/json"}, {"Content-Type": "text/html"}},
			expected: map[string]string{"Content-Type": "text/html"},
		},
	}

	// Run tests
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := HeadersToMap(tc.headers...)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected: %v, but got: %v", tc.expected, result)
			}
		})
	}
}
