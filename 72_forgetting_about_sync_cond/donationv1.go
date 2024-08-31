package main

import (
	"fmt"
	"sync"
	"time"
)

type Donation struct {
	mu      sync.RWMutex
	balance int
}

// use mutex: makes the CPU usage gigantic
func donationV1() {
	donation := &Donation{}

	// Listener goroutines
	f := func(goal int) {
		donation.mu.RLock()
		for donation.balance < goal {
			donation.mu.RUnlock()
			donation.mu.RLock()
		}
		fmt.Printf("$%d goal reached\n", donation.balance)
		donation.mu.RUnlock()
	}

	go f(10)
	go f(15)

	// Updater goroutine
	go func() {
		for {
			time.Sleep(time.Second)
			donation.mu.Lock()
			donation.balance++
			donation.mu.Unlock()
		}
	}()
	time.Sleep(20 * time.Second)
}
