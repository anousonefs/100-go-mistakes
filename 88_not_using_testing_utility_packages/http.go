package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("X-API-VERSION", "1.0")
	b, _ := io.ReadAll(r.Body)
	_, _ = w.Write(append([]byte("hello "), b...))
	w.WriteHeader(http.StatusCreated)
}

// test

type DurationClient struct {
	client *http.Client
}

// buildRequestBody creates the JSON request body for the API call.
func buildRequestBody(lat1, lng1, lat2, lng2 float64) *bytes.Buffer {
	req := map[string]interface{}{
		"start": map[string]float64{
			"lat": lat1,
			"lng": lng1,
		},
		"end": map[string]float64{
			"lat": lat2,
			"lng": lng2,
		},
	}
	body, _ := json.Marshal(req)
	return bytes.NewBuffer(body)
}

// parseResponseBody reads the response body and parses the duration.
func parseResponseBody(body io.ReadCloser) (time.Duration, error) {
	defer body.Close()

	var response struct {
		Duration int64 `json:"duration"` // duration is expected to be in seconds in the JSON response
	}

	if err := json.NewDecoder(body).Decode(&response); err != nil {
		return 0, fmt.Errorf("failed to parse response: %w", err)
	}

	// Convert seconds to time.Duration
	return time.Duration(response.Duration) * time.Second, nil
}

// GetDuration interacts with an API to get the travel duration between two coordinates.
func (c DurationClient) GetDuration(url string,
	lat1, lng1, lat2, lng2 float64) (time.Duration, error) {

	// Make the HTTP POST request to the provided `url`
	resp, err := c.client.Post(
		url, "application/json",
		buildRequestBody(lat1, lng1, lat2, lng2),
	)
	if err != nil {
		return 0, fmt.Errorf("failed to make API request: %w", err)
	}
	defer resp.Body.Close()

	// Check if the request was successful.
	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	// Parse the response body into a time.Duration
	return parseResponseBody(resp.Body)
}

func httpDemo() {
	client := DurationClient{client: &http.Client{Timeout: 10 * time.Second}}

	// Example usage: Call the `GetDuration` method with sample coordinates and a mock URL.
	lat1, lng1 := 40.7128, -74.0060  // New York City coords
	lat2, lng2 := 34.0522, -118.2437 // Los Angeles coords
	duration, err := client.GetDuration("http://example.com/get_duration", lat1, lng1, lat2, lng2)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print the returned duration (in hours and minutes).
	fmt.Printf("Duration between the two points: %v\n", duration)
}
