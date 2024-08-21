package main

import (
	"database/sql"
	"errors"
	"fmt"
)

/*1. Expected errors should be designed as error values (sentinel errors): var ErrFoo = errors.New("foo"). */
/*2. Unexpected errors should be designed as error types: type BarError struct { ... }, with BarError implementing the error interface. */

var ErrFoo = errors.New("foo")

func main() {
	/* getUser() */
	getUser2()
}

func getUser() {
	err := query()
	if err != nil {
		if err == sql.ErrNoRows { // alway false if error is wraped
			println("sql err no rows")
		} else {
		}
	}
}

func getUser2() {
	err := query()
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) { // alway true if error is wraped or not
			println("sql err no rows")
		} else {
		}
	}
}

func query() error {
	/* return sql.ErrNoRows */
	return fmt.Errorf("getUser: %w", sql.ErrNoRows)
}
