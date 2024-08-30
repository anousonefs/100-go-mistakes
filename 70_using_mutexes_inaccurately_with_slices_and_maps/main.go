package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	cache := Cache{
		balances: make(map[string]float64),
	}

	go func() {
		cache.AddBalance("user1", 100.50)
		cache.AddBalance("user2", 200.50)
		cache.AddBalance("user3", 300.50)
	}()

	go func() {
		fmt.Println(cache.AverageBalanceV2()) // output: 100.5 or 200.5 or 150.5
	}()

	time.Sleep(100 * time.Millisecond)

}

type Cache struct {
	mu       sync.RWMutex
	balances map[string]float64
}

func (c *Cache) AddBalance(id string, balance float64) {
	c.mu.Lock()
	c.balances[id] = balance
	c.mu.Unlock()
}

// balances is being use in loop but another goroutine try to update balances
func (c *Cache) AverageBalance() float64 {
	c.mu.RLock()
	balances := c.balances // does't copy the actual data
	c.mu.RUnlock()
	sum := 0.
	for _, balance := range balances {
		sum += balance
	}
	return sum / float64(len(balances))
}

// lock until end process
func (c *Cache) AverageBalanceV2() float64 {
	c.mu.RLock()
	defer c.mu.RUnlock()
	balances := c.balances // does't copy the actual data
	sum := 0.
	for _, balance := range balances {
		sum += balance
	}
	return sum / float64(len(balances))
}

// copies the map
func (c *Cache) AverageBalanceV3() float64 {
	c.mu.RLock()
	m := make(map[string]float64, len(c.balances))
	for k, v := range c.balances {
		m[k] = v
	}
	c.mu.RUnlock()
	sum := 0.
	for _, balance := range m {
		sum += balance
	}
	return sum / float64(len(m))
}
