package main

import "errors"

func main() {
	f()
}

/* Error handling is omitted. */
func f() {
	notify()
}

/* ignored the error intentionally */
func f2() {
	_ = notify()
}

func notify() error {
	return errors.New("notify")
}
