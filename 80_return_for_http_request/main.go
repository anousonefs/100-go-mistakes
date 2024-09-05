package main

import (
	"errors"
	"net/http"
)

func main() {

}

func handler(w http.ResponseWriter, req *http.Request) {
	err := foo(req)
	if err != nil {
		http.Error(w, "foo", http.StatusInternalServerError)
		// Thanks to the return statement, the function will stop its
		// execution if we end in the if err != nil branch.
		return
	}
	_, _ = w.Write([]byte("all good"))
	w.WriteHeader(http.StatusCreated)
}

func foo(req *http.Request) error {
	// Example: Check for the existence of a query parameter "id"
	id := req.URL.Query().Get("id")
	if id == "" {
		return errors.New("missing 'id' parameter")
	}

	// Simulate some processing logic, such as validating the ID
	if len(id) < 3 {
		return errors.New("invalid 'id' parameter")
	}

	// Return nil if everything is fine
	return nil
}
