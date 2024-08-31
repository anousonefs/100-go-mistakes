package main

import (
	"fmt"
	"time"
)

type DonationV2 struct {
	balance int
	ch      chan int
}

// use channel
func donationV2() {
	donation := &DonationV2{ch: make(chan int)}

	// Listener goroutines
	f := func(goal int) {
		for balance := range donation.ch {
			if balance >= goal {
				fmt.Printf("$%d goal reached\n", balance)
				return
			}
		}
	}
	go f(10) // output: 11
	go f(15) // output: 15

	// Updater goroutine
	for {
		time.Sleep(time.Second)
		donation.balance++
		donation.ch <- donation.balance
	}

}
