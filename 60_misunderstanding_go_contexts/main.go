package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/fsnotify/fsnotify"
)

func main() {
	cancelDemo()
}

/* When main returns, it calls the cancel function to cancel the context passed to CreateFileWatcher so that the file descriptor is closed gracefully. */
func cancelDemo() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go func() {
		if err := CreateFileWatcher(ctx, "foo.txt"); err != nil {
			if errors.Is(err, context.Canceled) {
				fmt.Printf("ctx err: %v\n", err)
			}
		}
	}()
	time.Sleep(4 * time.Second)
}

type Position struct {
	x float32
	y float32
}

type publisher interface {
	Publish(ctx context.Context, position Position) error
}

type publishHandler struct {
	pub publisher
}

// deadline 4 second
func (h publishHandler) publishPosition(position Position) error {
	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second) // go routine(context) will live for 4 seconds
	defer cancel()                                                          // if parent function return then destoy go routine(context)
	return h.pub.Publish(ctx, position)
}

/* 1. Deadline */

func CreateFileWatcher(ctx context.Context, fileName string) error {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	// Start watching the file
	err = watcher.Add(fileName)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Started watching file: %s\n", fileName)

	for {
		select {
		case event := <-watcher.Events:
			if event.Op&fsnotify.Write == fsnotify.Write {
				fmt.Printf("Modified file: %s\n", event.Name)
			}
		case err := <-watcher.Errors:
			fmt.Printf("Error: %s\n", err)
		case <-ctx.Done():
			fmt.Println("File watcher stopped")
			return ctx.Err()
		}
	}
}

// -- context values ----

type key string

/* The isValidHostKey constant is unexported. Hence, there’s no risk that another package using the same context could override the value that is already set. Even if another package creates the same myCustomKey based on a key type as well, it will be a different key. */
const isValidHostKey key = "isValidHost"

func checkValid(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		validHost := r.Host == "acme"
		ctx := context.WithValue(r.Context(), isValidHostKey, validHost)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// ctx.Err()
/* 1. context.Canceled */
/* 2. context.DeadlineExceeded */
/* 3. nil if the Done channel isn’t yet closed */
type Message struct {
}

func handler(ctx context.Context, ch chan Message) error {
	for {
		select {
		case msg := <-ch: // Do something with msg
			fmt.Print(msg)
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}
