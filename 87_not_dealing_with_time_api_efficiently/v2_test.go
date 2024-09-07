package main

import (
	"fmt"
	"testing"
	"time"
)

// tests would no longer be isolated because they would all depend on a shared variable
// var now = time.Now

func TestCache_TrimOlderThanV2(t *testing.T) {
	events := []Event{
		{Timestamp: parseTime(t, "2020-01-01T12:00:00.04Z")}, // -20
		{Timestamp: parseTime(t, "2020-01-01T12:00:00.05Z")}, // -10
		{Timestamp: parseTime(t, "2020-01-01T12:00:00.06Z")}, // current time
	}
	cache := &CacheV2{now: func() time.Time {
		return parseTime(t, "2020-01-01T12:00:00.06Z")
	}}
	cache.Add(events)
	cache.TrimOlderThan(15 * time.Millisecond)

	got := cache.GetAll()
	fmt.Printf("got: %#v\n", got)
	expected := 2
	if len(got) != expected {
		t.Fatalf("expected %d, got %d", expected, len(got))
	}
}

func parseTime(t *testing.T, timestamp string) time.Time {
	layout := time.RFC3339 // RFC3339 is the format used by ISO 8601 timestamps like "2020-01-01T12:00:00.04Z"
	parsedTime, err := time.Parse(layout, timestamp)
	if err != nil {
		t.Fatalf("failed to parse time: %v", err)
	}
	return parsedTime
}
