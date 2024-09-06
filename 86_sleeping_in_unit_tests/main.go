package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

}

type Handler struct {
	n         int
	publisher publisher
}

type publisher interface {
	Publish([]Foo)
}

func (h Handler) getBestFoo(someInputs int) Foo {
	foos := getFoos(someInputs)
	best := foos[0]
	go func() {
		if len(foos) > h.n {
			foos = foos[:h.n]
		}
		time.Sleep(5 * time.Millisecond)
		h.publisher.Publish(foos)
	}()
	return best
}

type Foo struct {
	ID   int
	Name string
}

func getFoos(someInputs int) []Foo {
	foos := []Foo{}
	for i := 0; i < someInputs; i++ {
		foos = append(foos, Foo{
			ID:   i,
			Name: fmt.Sprintf("Foo%d", i),
		})
	}
	return foos
}

// test

type publisherMock struct {
	mu  sync.RWMutex
	got []Foo
}

func (p *publisherMock) Publish(got []Foo) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.got = got
}

func (p *publisherMock) Get() []Foo {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return p.got
}

type publisherMock2 struct {
	ch chan []Foo
}

func (p *publisherMock2) Publish(got []Foo) {
	time.Sleep(500 * time.Millisecond)
	p.ch <- got
}
