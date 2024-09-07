package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestHandler(t *testing.T) {
	// Mock an HTTP GET request with a body containing "foo"
	req := httptest.NewRequest(http.MethodGet, "http://localhost", strings.NewReader("foo"))

	// Create a ResponseRecorder to record the response
	w := httptest.NewRecorder()

	// Call the handler under test
	Handler(w, req)

	// Check that the "X-API-VERSION" header is set correctly
	result := w.Result() // store result once

	fmt.Printf("header: %v\n", result)
	if got := result.Header.Get("X-API-VERSION"); got != "1.0" {
		t.Errorf("api version: expected 1.0, got %s", got)
	}

	// Read and check the response body
	body, err := io.ReadAll(result.Body)
	if err != nil {
		t.Fatalf("failed to read response body: %v", err)
	}

	// Check that the response body matches "hello foo"
	if got := string(body); got != "hello foo" {
		t.Errorf("body: expected hello foo, got %s", got)
	}

	// Assert that the status code is 200 OK
	if result.StatusCode != http.StatusOK {
		t.Fatalf("expected status OK, got %d", result.StatusCode)
	}
}

func TestDurationClientGet(t *testing.T) {
	srv := httptest.NewServer(
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				_, _ = w.Write([]byte(`{"duration": 314}`))
			},
		))
	defer srv.Close()
	client := NewDurationClient()
	duration, err :=
		client.GetDuration(srv.URL, 51.551261, -0.1221146, 51.57, -0.13)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("duration: %v\n", duration)
	if duration != 314*time.Second {
		t.Errorf("expected 314 seconds, got %v", duration)
	}
}

func NewDurationClient() *DurationClient {
	return &DurationClient{
		client: &http.Client{
			Timeout: 10 * time.Second, // You can adjust this timeout as necessary.
		},
	}
}
