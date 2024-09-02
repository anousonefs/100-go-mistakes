package main

import "sync"

func main() {
	counter := NewCounter()
	go func() {
		counter.Increment("foo")
	}()
	go func() {
		counter.Increment("bar")
	}()
}

type Counter struct {
	mu       sync.Mutex
	counters map[string]int
}

func NewCounter() Counter {
	return Counter{counters: map[string]int{}}
}

/*
sync types shouldn’t be copied. This rule applies to the following types:
    sync.Cond
    sync.Map
    sync.Mutex
    sync.RWMutex
    sync.Once
    sync.Pool
    sync.WaitGroup
*/
/* it performs a copy of the Counter struct, which also copies the mutex */
func (c Counter) Increment(name string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counters[name]++
}

// use pointer for prevent copy
func (c *Counter) Increment2(name string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counters[name]++
}
