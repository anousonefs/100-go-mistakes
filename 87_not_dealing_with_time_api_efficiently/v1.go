package main

import (
	"sync"
	"time"
)

type Cache struct {
	mu     sync.RWMutex
	events []Event
}

func (c *Cache) GetAll() []Event {
	return c.events
}

func (c *Cache) Add(event []Event) {
	c.events = append(c.events, event...)
}

type Event struct {
	Timestamp time.Time
	Data      string
}

func (c *Cache) TrimOlderThan(since time.Duration) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	t := time.Now().Add(-since)
	for i := 0; i < len(c.events); i++ {
		if c.events[i].Timestamp.After(t) {
			c.events = c.events[i:]
			return
		}
	}
}
