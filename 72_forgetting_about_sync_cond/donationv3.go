package main

import (
	"fmt"
	"sync"
	"time"
)

type DonationV3 struct {
	cond    *sync.Cond
	balance int
}

// use sync.Cond
func donationV3() {
	donation := &DonationV3{
		cond: sync.NewCond(&sync.Mutex{}),
	}

	// Listener goroutines
	f := func(goal int) {
		donation.cond.L.Lock()
		for donation.balance < goal {
			donation.cond.Wait() // unlock -> wait signal -> lock
		}
		fmt.Printf("%d$ goal reached\n", donation.balance)
		donation.cond.L.Unlock()
	}
	go f(10)
	go f(15)

	// Updater goroutine
	for {
		time.Sleep(time.Second)
		donation.cond.L.Lock()
		donation.balance++
		donation.cond.L.Unlock()
		donation.cond.Broadcast()
	}
}
