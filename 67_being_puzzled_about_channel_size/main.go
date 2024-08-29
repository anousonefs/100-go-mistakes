package main

import (
	"fmt"
	"time"
)

/*
1. An unbuffered channel enables synchronization.
	 We have the guarantee that two goroutines will be in a known state:
   one receiving and another sending a message.

2. A buffered channel doesn’t provide any strong synchronization.
	Indeed, a producer goroutine can send a message and then continue its execution if the channel isn’t full.
	The only guarantee is that a goroutine won’t receive a message before it is sent.
	But this is only a guarantee because of causality (you don’t drink your coffee before you prepare it).

Note:
  1. buffered channels can lead to obscure deadlocks that would be immediately apparent with unbuffered channels.
  2. notification channel should not use buffered channel
*/

func main() {
	/* ch3 := make(chan int, 1) // buffered channel */
	/* ch3 <- 1                 // Non-blocking */
	/* ch3 <- 2                 // Blocking */

	/* ch4 := make(chan int) */
	ch4 := make(chan int, 2)
	go func() {
		ch4 <- 1
		ch4 <- 4
		println("----")
		ch4 <- 5 // blocking
		ch4 <- 6
	}()
	fmt.Printf("ch4: %v\n", <-ch4)
	fmt.Printf("ch4: %v\n", <-ch4)
	fmt.Printf("ch4: %v\n", <-ch4)
	fmt.Printf("ch4: %v\n", <-ch4)
	time.Sleep(400 * time.Millisecond)
}
