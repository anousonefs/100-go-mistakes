package main

import (
	"errors"
	"fmt"
)

func main() {
	s := "Hello, World!"
	x := s[:5] // Hello
	println(x)

	s2 := "Hêllo, World!"
	y := string([]rune(s2)[:5]) // Hêllo
	println(y)
	demo()
}

type store struct {
}

func (f store) myStore(uuid string) {

}

func (s store) handleLog(log string) error {
	if len(log) < 36 {
		return errors.New("log is not correctly formatted")
	}
	/*
		GoLand, the Go JetBrains IDE, warns about a redundant type conversion.
		it prevents the new string from being backed by the same array as uuid. We need to be aware that the warnings raised by IDEs or linters may sometimes be inaccurate. */

	/* log[:36] will create a new string referencing the same backing array. Therefore, each uuid string that we store in */
	/* memory will contain not just 36 bytes but the number of bytes in the initial log string: potentially, thousands of bytes. */
	/* uuid := log[:36] */

	uuid := string([]byte(log[:36])) // preventing a memory leak
	/* uuid := strings.Clone(log[:36]) */
	/* uuid := log[:36:36] */ // referencing same backing array

	s.myStore(uuid)
	// Do something
	return nil
}

/* NOTE Because a string is mostly a pointer, calling a function to pass a string doesn’t result in a deep copy of the bytes. The copied string will still reference the same backing array. */

func demo() {
	x := "anousone"
	y := x[:3]
	x = "worlakoumman"
	fmt.Printf("x: %v\n", x)
	fmt.Printf("y: %v\n", y)
}
