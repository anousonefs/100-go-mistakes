package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	increment()
}

func increment() {
	wg := sync.WaitGroup{}
	var v uint64
	for i := 0; i < 3; i++ {
		go func() {
			wg.Add(1)
			atomic.AddUint64(&v, 1)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(v)
}

/*
	When using a sync.WaitGroup, the Add operation must be done before spinning up a goroutine in the parent goroutine,
	whereas the Done operation must be done within the goroutine.
*/

func increment2() {
	wg := sync.WaitGroup{}
	var v uint64
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddUint64(&v, 1)
		}()
	}
	wg.Wait()
	fmt.Println(v)
}
