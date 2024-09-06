package main

import (
	"fmt"
	"testing"
	"time"
)

// use mutex
func TestGetBestFoo(t *testing.T) {
	mock := publisherMock{}
	h := Handler{
		publisher: &mock,
		n:         4}
	foo := h.getBestFoo(42)
	fmt.Printf("foo: %v\n", foo)

	/* time.Sleep(10 * time.Millisecond) // Sleeps for 10 milliseconds before checking the arguments passed to Publish */

	assert(t, func() bool {
		return len(mock.Get()) == 4
	}, 30, time.Millisecond)

	published := mock.Get()
	fmt.Printf("published: %v\n", published)
}

func assert(t *testing.T, assertion func() bool,
	maxRetry int, waitTime time.Duration) {
	for i := 0; i < maxRetry; i++ {
		fmt.Printf("i: %v\n", i)
		if assertion() {
			return
		}
		time.Sleep(waitTime)
	}
	t.Fail()
}

// use channel
func TestGetBestFoo2(t *testing.T) {
	mock := publisherMock2{
		ch: make(chan []Foo),
	}
	defer close(mock.ch)
	h := Handler{
		publisher: &mock,
		n:         3,
	}
	foo := h.getBestFoo(42)
	_ = foo
	// Check foo
	if v := len(<-mock.ch); v != 3 {
		t.Fatalf("expected 2, got %d", v)
	}
}
