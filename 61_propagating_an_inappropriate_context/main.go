package main

import (
	"context"
	"log"
	"net/http"
	"time"
)

func main() {

}

// demo: send message to kafka
/* request can cancel in different conditions: */
/* 1. When the client’s connection closes */
/* 2. In the case of an HTTP/2 request, when the request is canceled */
/* 3. When the response has been written back to the client */
func handler(w http.ResponseWriter, r *http.Request) {
	response, err := doSomeTask(r.Context(), r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	go func() {
		// should not use the same context
		err := sendMessageToKafka(r.Context(), response)
		if err != nil {
		}
	}()
	writeResponse(w, response)
}

func handlerV2(w http.ResponseWriter, r *http.Request) {
	response, err := doSomeTask(r.Context(), r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	go func() {
		// use context.Background()
		err := sendMessageToKafka(context.Background(), response)
		if err != nil {
		}
	}()
	writeResponse(w, response)
}

/*
Now the context passed to publish will never expire or be canceled,
but it will carry the parent context’s values.
*/
func handlerV3(w http.ResponseWriter, r *http.Request) {
	response, err := doSomeTask(r.Context(), r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	go func() {
		// use custom context for retrieve a value from parent context
		err := sendMessageToKafka(detach{ctx: r.Context()}, response)
		if err != nil {
		}
	}()
	writeResponse(w, response)
}

func handlerV4(w http.ResponseWriter, r *http.Request) {
	response, err := doSomeTask(r.Context(), r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	go func() {
		// if client close connection then stop send message to kafka
		/* if err := r.Context().Err(); err != nil { */
		/* } */
		err := sendMessageToKafka(detach{ctx: r.Context()}, response)
		if err != nil {
		}
	}()
	writeResponse(w, response)
}

func doSomeTask(ctx context.Context, _ *http.Request) (string, error) {
	// Simulate processing the request
	select {
	case <-time.After(2 * time.Second): // Simulate a delay
		// Return a response after processing
		return "Task completed successfully", nil
	case <-ctx.Done(): // Handle context cancellation
		return "", ctx.Err()
	}
}

func sendMessageToKafka(ctx context.Context, response string) error {
	// Simulate publishing the response
	select {
	case <-time.After(3 * time.Second): // Simulate a delay
		// Publishing successful
		log.Println("Response published:", response)
		return nil
	case <-ctx.Done(): // Handle context cancellation
		return ctx.Err()
	}
}

func writeResponse(w http.ResponseWriter, response string) {
	// Write the response to the HTTP client
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/plain")
	_, err := w.Write([]byte(response))
	if err != nil {
		log.Printf("Error writing response: %v", err)
	}
}

type detach struct {
	ctx context.Context
}

func (d detach) Deadline() (time.Time, bool) {
	return time.Time{}, false
}
func (d detach) Done() <-chan struct{} {
	return nil
}
func (d detach) Err() error {
	return nil
}

/* Except for the Value method that calls the parent context to retrieve a value, the other methods return a default value so the context is never considered expired or canceled. */
func (d detach) Value(key any) any {
	return d.ctx.Value(key)
}
