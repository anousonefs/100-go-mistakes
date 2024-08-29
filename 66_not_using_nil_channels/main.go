package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int) // unbuffered channel
	ch2 := make(chan int) // unbuffered channel

	go func() {
		ch1 <- 2
		close(ch1)
	}()

	go func() {
		ch2 <- 3
		time.Sleep(800 * time.Millisecond)
		ch2 <- 4
		close(ch2)
	}()

	ch := mergeV4(ch1, ch2)

	// loop until ch is closed
	for v := range ch {
		fmt.Printf("Received: %v\n", v)
	}

	ch3 := make(chan int)
	close(ch3)
	fmt.Print(<-ch3, <-ch3) // 0, 0
}

/*
	The main issue with this first version is that we receive from ch1 and then we receive from ch2.

It means we won’t receive from ch2 until ch1 is closed.
This doesn’t fit our use case, as ch1 may be open forever,
so we want to receive from both channels simultaneously.
*/
func merge(ch1, ch2 <-chan int) <-chan int {
	ch := make(chan int, 1)
	go func() {
		for v := range ch1 {
			ch <- v
		}
		for v := range ch2 {
			ch <- v
		}
		close(ch)
	}()
	return ch
}

// loop forever
func mergeV2(ch1, ch2 <-chan int) <-chan int {
	ch := make(chan int, 1)
	go func() {
		for {
			select {
			case v := <-ch1:
				ch <- v
			case v := <-ch2:
				ch <- v
			}
		}
		close(ch) // unreachable code
	}()
	return ch
}

/*
	The first condition, ch1 is closed, will always be valid.

Therefore, as long as we don’t receive a message in ch2 and this channel isn’t closed,
we will keep looping over the first case.
This will lead to wasting CPU cycles and must be avoided. Therefore,
our solution isn’t viable.
*/
func mergeV3(ch1, ch2 <-chan int) <-chan int {
	ch := make(chan int, 1)
	ch1Closed := false
	ch2Closed := false
	go func() {
		for {
			select {
			case v, open := <-ch1:
				fmt.Println("ch1")
				if !open {
					ch1Closed = true
					break
				}
				ch <- v
			case v, open := <-ch2:
				fmt.Println("ch2")
				if !open {
					ch2Closed = true
					break
				}
				ch <- v
			}
			if ch1Closed && ch2Closed {
				close(ch)
				return
			}
		}
	}()
	return ch
}

/* receiving from a nil channel will block forever */
func mergeV4(ch1, ch2 <-chan int) <-chan int {
	ch := make(chan int)

	go func() {
		defer close(ch)
		for ch1 != nil || ch2 != nil {
			select {
			case v, ok := <-ch1:
				fmt.Println("ch1")
				if !ok {
					ch1 = nil // Mark ch1 as closed
					continue
				}
				ch <- v
			case v, ok := <-ch2:
				fmt.Println("ch2")
				if !ok {
					ch2 = nil // Mark ch2 as closed
					continue
				}
				ch <- v
			}
		}
	}()

	return ch
}
