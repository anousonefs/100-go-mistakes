package main

import (
	"fmt"
	"sync"
	"time"
	"unsafe"
)

// Worker function that does some work until notified to stop
func worker(id int, stopCh <-chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker %d starting\n", id)

	for {
		select {
		case <-stopCh:
			fmt.Printf("Worker %d stopping\n", id)
			return
		default:
			// Simulate work
			fmt.Printf("Worker %d is working...\n", id)
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	var wg sync.WaitGroup

	var s struct{}
	fmt.Printf("struct size: %v\n", unsafe.Sizeof(s))
	var i interface{}
	fmt.Printf("interface size: %v\n", unsafe.Sizeof(i))
	var b bool
	fmt.Printf("bool size: %v\n", unsafe.Sizeof(b))

	stopCh := make(chan struct{}) // use struct{} because it occupies zero bytes of storage

	// Start 3 workers
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, stopCh, &wg)
	}

	// Let the workers run for 2 seconds
	time.Sleep(2 * time.Second)

	// Notify all workers to stop
	close(stopCh)

	// Wait for all workers to finish
	wg.Wait()

	fmt.Println("All workers have stopped.")
}
