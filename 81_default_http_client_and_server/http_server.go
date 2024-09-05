package main

import (
	"fmt"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(2100 * time.Millisecond)
	fmt.Fprintln(w, "Request processed successfully")
}

func startServer() {
	/*
			   Note that if http.Server.IdleTimeout isn’t set, the value of http.Server .ReadTimeout is used for the idle timeout. If
		neither is set, there won’t be any timeouts, and connections will remain open until they are closed by clients.
	*/

	s := &http.Server{
		Addr:              ":8080",
		ReadHeaderTimeout: 500 * time.Millisecond,
		ReadTimeout:       500 * time.Millisecond,                                                           // entire request
		Handler:           http.TimeoutHandler(http.HandlerFunc(handler), 2*time.Second, "Request timeout"), // Wrap the handler with TimeoutHandler
		IdleTimeout:       30 * time.Second,                                                                 // default is 3 minute
	}

	fmt.Println("Starting server on :8080...")
	if err := s.ListenAndServe(); err != nil {
		fmt.Println("Server error:", err)
	}
}
