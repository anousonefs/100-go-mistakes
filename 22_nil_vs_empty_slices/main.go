package main

import "fmt"

func main() {
	var s []string
	log(1, s)
	s = []string(nil)
	log(2, s)
	s = []string{}
	log(3, s)
	s = make([]string, 0)
	log(4, s)
}

func log(i int, s []string) {
	fmt.Printf("%d: empty=%t\tnil=%t\n", i, len(s) == 0, s == nil)
}

/* 1 var s []string if we aren’t sure about the final length and the slice can be empty */

/* 2 []string(nil) as syntactic sugar to create a nil and empty slice */
/* s := append([]int(nil), 42) */

/* 3 []string{}, should be avoided if we initialize the slice without elements. */

/* 4 make([]string, length) if the future length is known */
