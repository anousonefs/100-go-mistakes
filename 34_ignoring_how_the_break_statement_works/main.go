package main

import (
	"context"
	"fmt"
)

func main() {
	// wrong
	for i := 0; i < 5; i++ {
		fmt.Printf("%d ", i)
		switch i {
		default:
		case 2:
			break
		}
	}
	// solution
loop:
	for i := 0; i < 5; i++ {
		fmt.Printf("%d ", i)
		switch i {
		default:
		case 2:
			break loop
		}
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	ch := make(chan int)

	// solution
loop2:
	for {
		select {
		case <-ch:
			// Do something
		case <-ctx.Done():
			break loop2
		}
	}

	// wrong
	for {
		select {
		case <-ch:
			// Do something
		case <-ctx.Done():
			break
		}
	}

}
