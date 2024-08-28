package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	/* unstopGoroutineDemo() */
	/* watcherDemo() */
	watcherV2Demo()
}

func unstopGoroutineDemo() {
	ch := foo()

	// Start a goroutine to process values from the channel
	go func() {
		for v := range ch { // Loop over values received from the channel
			fmt.Println("Received:", v) // Print each received value
		}
	}()

	// Keep the main function alive until the processing is done
	time.Sleep(3 * time.Second)
}

// foo function returns a channel that emits integers
func foo() <-chan int {
	ch := make(chan int)

	// Start a goroutine that sends values to the channel
	go func() {
		defer close(ch) // Ensure the channel is closed when done
		for i := 0; i < 5; i++ {
			ch <- i                            // Send an integer to the channel
			time.Sleep(500 * time.Millisecond) // Simulate some work with a delay
		}
	}()

	return ch
}

//---------------- watcher issue --------------------

type watcher struct {
}

// watch method simulates watching for something, handling context cancellation
func (w *watcher) watch(ctx context.Context) {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done(): // Handle context cancellation
			time.Sleep(2 * time.Second) // take time for clear resource
			fmt.Println("watcher: context cancelled")
			return
		case t := <-ticker.C: // Simulate periodic work
			fmt.Println("watcher: working at", t)
		}
	}
}

// newWatcher function starts a new watcher and runs it in a goroutine
func newWatcher(ctx context.Context) {
	w := watcher{}
	go w.watch(ctx) // if ctx is cancel the parent do not have much time to wait goroutine to clear resource
}

func watcherDemo() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	newWatcher(ctx) // Run the watcher

	// Simulate some other work in the main function
	time.Sleep(5 * time.Second)
	fmt.Println("main: cancelling context")
	cancel() // Cancel the context to stop the watcher

	// Wait for a moment to ensure the watcher exits cleanly
	/* time.Sleep(1 * time.Second) */
	fmt.Println("main: exiting")
}

// ----------------- watcher fix -------------

type watcherV2 struct {
	stopChan chan struct{}
	wg       sync.WaitGroup
}

// watch method simulates watching for something
func (w *watcherV2) watch() {
	w.wg.Add(1)
	defer w.wg.Done()

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-w.stopChan:
			time.Sleep(2 * time.Second) // take time for clear resource
			fmt.Println("watcher: received stop signal")
			return
		case t := <-ticker.C:
			fmt.Println("watcher: working at", t)
		}
	}
}

// newWatcher function initializes the watcher
func newWatcherV2() *watcherV2 {
	w := &watcherV2{
		stopChan: make(chan struct{}),
	}
	go w.watch()
	return w
}

// close method signals the watcher to stop and waits for it to finish
func (w *watcherV2) close() {
	close(w.stopChan)
	w.wg.Wait()
	fmt.Println("watcher: closed")
}

// clear resource before exit program
func watcherV2Demo() {
	w := newWatcherV2()
	defer w.close() // Ensure resources are cleaned up

	// Simulate doing some work in the main function
	time.Sleep(5 * time.Second)
	fmt.Println("main: finished working")
}
