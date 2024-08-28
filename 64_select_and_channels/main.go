package main

import (
	"fmt"
	"time"
)

func demo2() {
	messageCh := make(chan int)
	disconnectCh := make(chan struct{})

	go func() {
		for {
			select {
			case v := <-messageCh:
				fmt.Println(v)
			case <-disconnectCh:
				fmt.Println("disconnection, return")
				return
			}
		}
	}()

	for i := 0; i < 10; i++ {
		messageCh <- i
	}
	disconnectCh <- struct{}{} // Signal disconnection
	time.Sleep(1 * time.Second)
}

func demo1() {
	messageCh := make(chan int)
	disconnectCh := make(chan struct{})

	// Producer: Sends messages to messageCh
	go func() {
		for i := 0; i < 10; i++ {
			messageCh <- i
		}
		disconnectCh <- struct{}{} // Signal disconnection
		close(messageCh)           // Close the message channel when done
	}()

	// Consumer: Receives and processes messages
	for {
		select {
		case v, ok := <-messageCh:
			if !ok { // Check if the channel is closed
				return
			}
			fmt.Println(v)
		case <-disconnectCh:
			fmt.Println("disconnection, return")
			return
		}
	}
}

func main() {
	/* demo1() */
	demo2()
}

/* When using select with multiple channels, we must remember that if multiple options are possible, the first case */
/* in the source order does not automatically win. Instead, Go selects randomly, so there’s no guarantee about which option will be chosen.
To overcome this behavior, in the case of a single producer goroutine, we can use either unbuffered channels or a single channel.
In the case of multiple producer goroutines, we can use inner selects and default to handle prioritizations */

/*
the book said output is:
0
1
2
3
4
disconnection, return
*/

// may be the old book was wrong
