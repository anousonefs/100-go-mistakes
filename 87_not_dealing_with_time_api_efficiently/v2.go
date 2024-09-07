package main

import (
	"sync"
	"time"
)

type now func() time.Time

type CacheV2 struct {
	mu     sync.RWMutex
	events []Event
	now    now
}

func NewCacheV2(now now) *CacheV2 {
	return &CacheV2{
		events: make([]Event, 0),
		now:    now,
	}
}

func (c *CacheV2) GetAll() []Event {
	return c.events
}

func (c *CacheV2) Add(event []Event) {
	c.events = append(c.events, event...)
}

func (c *CacheV2) TrimOlderThan(since time.Duration) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	t := c.now().Add(-since)
	for i := 0; i < len(c.events); i++ {
		if c.events[i].Timestamp.After(t) {
			c.events = c.events[i:]
			return
		}
	}
}

func (c *Cache) TrimOlderThan2(now time.Time, since time.Duration) {
	// ...
}

func (c *Cache) TrimOlderThan3(t time.Time) {
	// ...
}

/* cache.TrimOlderThan3(time.Now().Add(time.Second)) */
