package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	/* dataRace() */
	/* atomicDemo() */
	/* mutexDemo() */
	/* channelDemo() */

	/* raceConditionDemo() */
	/* demo() */
	/* demo2() */
	/* demo3() */
	demo4()
}

// data race
func dataRace() {
	i := 0
	go func() {
		i++
	}()
	go func() {
		i++
	}()
}

// fix data race: by use atomic
func atomicDemo() {
	var counter int64
	var w sync.WaitGroup
	w.Add(2)
	go func() {
		w.Add(1)
		defer w.Done()
		atomic.AddInt64(&counter, 1)
	}()
	go func() {
		w.Add(1)
		defer w.Done()
		atomic.AddInt64(&counter, 1)
	}()
	w.Wait()
	fmt.Println("Counter:", counter)
}

// fix data race: by use mutex
func mutexDemo() {
	counter := 0
	mutex := sync.Mutex{}
	var w sync.WaitGroup
	w.Add(2)
	go func() {
		defer w.Done()
		mutex.Lock()
		counter++
		mutex.Unlock()
	}()
	go func() {
		defer w.Done()
		mutex.Lock()
		counter++
		mutex.Unlock()
	}()
	w.Wait()
	fmt.Println("Counter:", counter)
}

// fix data race: by use channel
func channelDemo() {
	counter := 0
	ch := make(chan int)
	go func() {
		ch <- 1
	}()
	go func() {
		ch <- 1
	}()
	counter += <-ch
	counter += <-ch
	fmt.Println("Counter:", counter)
}

/* 1. Using atomic operations */
/* 2. Protecting a critical section with a mutex */
/* 3. Using communication and channels to ensure that a variable is updated by only one goroutine */

// output: 1 or 2
func raceConditionDemo() {
	count := 0
	mutex := sync.Mutex{}
	var w sync.WaitGroup
	w.Add(2)
	go func() {
		defer w.Done()
		mutex.Lock()
		defer mutex.Unlock()
		count = 1
	}()
	go func() {
		defer w.Done()
		mutex.Lock()
		defer mutex.Unlock()
		count = 2
	}()
	w.Wait()
	fmt.Printf("count: %v\n", count)
}

/* summary:  */
/* A data race occurs when multiple goroutines simultaneously access the same memory location and at least one of them is writing. A data race means unexpected behavior. However, a data-race-free application doesn’t necessarily mean deterministic results. An application can be free of data races but still have behavior that depends on uncontrolled events (such as goroutine execution, how fast a message is published to a channel, or how long a call to a database lasts); this is a race condition. Understanding both concepts is crucial to becoming proficient in designing concurrent applications. */

/* The goroutine is launched but remains blocked at the line <-ch waiting to receive a value from the channel. */
func demo() {
	i := 0
	ch := make(chan struct{})
	go func() {
		<-ch
		fmt.Println(i)
	}()
	i++
	ch <- struct{}{}
}

func demo2() {
	i := 0
	ch := make(chan struct{})
	go func() {
		<-ch
		fmt.Println(i)
	}()
	i++
	close(ch)
}

// ---------- go memory model -----------------

// channel buffered: data race
func demo3() {
	i := 0
	ch := make(chan struct{}, 1)
	go func() {
		i = 1
		<-ch
	}()
	ch <- struct{}{}
	fmt.Println(i)
}

// channel unbuffered
/* they represent the ordering guarantees of the Go memory model.
Because a receive from an unbuffered channel happens before a send,
the write to i will always occur before the read. */

/* Don't be clever. */
func demo4() {
	i := 0
	ch := make(chan struct{})
	go func() {
		i = 1
		<-ch
	}()
	ch <- struct{}{}
	fmt.Println(i)
}
