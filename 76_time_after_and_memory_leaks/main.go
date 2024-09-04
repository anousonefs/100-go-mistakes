package main

import (
	"context"
	"log"
	"strconv"
	"sync"
	"time"
)

type Event struct {
	ID   int
	Name string
}

func handle(event Event) {
	log.Printf("Handling event: %v\n", event)
}

// use time.After
func consumer(ch <-chan Event) {
	for {
		select {
		case event, ok := <-ch:
			if !ok {
				log.Println("channel closed, exiting consumer")
				return
			}
			handle(event)
		/* time.After returns a channel
		  In Go 1.15, about 200 bytes of memory are used per call to time.After.
			If we receive a significant volume of messages, such as 5 million per hour,
			our application will consume 1 GB of memory to store the time.After resources.
			Note: we can not close channel. because it is a receive-only
		*/
		case <-time.After(time.Second * 2):
			log.Println("warning: no messages received")
		}
	}
}

// use context
func consumerV2(ch <-chan Event) {
	for {
		//re-create a context during every single loop iteration
		//Creating a context isn’t the most lightweight operation in Go
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
		select {
		case event, ok := <-ch:
			cancel()
			if !ok {
				log.Println("channel closed, exiting consumer")
				return
			}
			handle(event)
		case <-ctx.Done():
			log.Println("warning: no messages received")
		}
	}
}

// use time.NewTimer
func consumerV3(ch <-chan Event) {
	timerDuration := 2 * time.Second
	timer := time.NewTimer(timerDuration)
	defer timer.Stop()
	for {
		timer.Reset(timerDuration)
		select {
		case event := <-ch:
			handle(event)
		case <-timer.C:
			log.Println("warning: no messages received")
		}
	}
}

// clear resource
func consumerV4(ch <-chan Event, done <-chan struct{}) {
	timerDuration := 2 * time.Second
	timer := time.NewTimer(timerDuration)
	defer timer.Stop()

	for {
		timer.Reset(timerDuration)
		select {
		case event := <-ch:
			handle(event)

		case <-timer.C:
			log.Println("warning: no messages received")

		case <-done:
			log.Println("consumer stopped")
			return
		}
	}
}

func main() {
	eventCh := make(chan Event)
	doneCh := make(chan struct{})

	go consumerV4(eventCh, doneCh)

	var w sync.WaitGroup
	w.Add(1)
	go func() {
		defer w.Done()
		for i := 1; i <= 7; i++ {
			eventCh <- Event{ID: i, Name: "Event " + strconv.Itoa(i)}
			time.Sleep(time.Second * 3)
			if i == 3 {
				close(doneCh)
				close(eventCh)
				break
			}
		}
	}()

	w.Wait()

}
