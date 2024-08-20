package main

import (
	"errors"
	"fmt"
	"net/http"
)

type transientError struct {
	err error
}

func (t transientError) Error() string {
	return fmt.Sprintf("transient error: %v", t.err)
}

func getTransactionAmount2(transactionID string) (float32, error) {
	if len(transactionID) != 5 {
		return 0, fmt.Errorf("id is invalid: %s", transactionID)
	}
	amount, err := getTransactionAmountFromDB2(transactionID)
	if err != nil {
		return 0, fmt.Errorf("failed to get transaction %s: %w", transactionID, err) // wrap error
	}
	return amount, nil
}

func getTransactionAmount(transactionID string) (float32, error) {
	if len(transactionID) != 5 {
		return 0, fmt.Errorf("id is invalid: %s", transactionID)
	}
	amount, err := getTransactionAmountFromDB(transactionID)
	if err != nil {
		return 0, transientError{err: err} // custom error
	}
	return amount, nil
}

func getTransactionAmountFromDB2(_ string) (float32, error) {
	return 0, transientError{err: errors.New("sql error")}
}

func getTransactionAmountFromDB(_ string) (float32, error) {
	return 150000.0, nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	transactionID := r.URL.Query().Get("transaction")
	_, err := getTransactionAmount(transactionID)
	if err != nil {
		switch err := err.(type) {
		case transientError:
			http.Error(w, err.Error(), http.StatusServiceUnavailable)
		default:
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		return
	}
	return
}

func handler2(w http.ResponseWriter, r *http.Request) {
	transactionID := r.URL.Query().Get("transaction")
	_, err := getTransactionAmount2(transactionID)
	/* _, err := getTransactionAmount(transactionID) */
	if err != nil {
		if errors.As(err, &transientError{}) { // use errors.As for unwraps an error
			http.Error(w, err.Error(),
				http.StatusServiceUnavailable)
		} else {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		return
	}
	return
}

func main() {

}
