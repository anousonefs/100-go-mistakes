package main

import "fmt"

func main() {
	/* append1() */
	/* append2() */
	append3()
}

// if length == capacity then create a new backing array
func append1() {
	s := make([]int, 1)
	go func() {
		s1 := append(s, 1)
		fmt.Println(s1) // [0,1]
	}()
	go func() {
		s2 := append(s, 1)
		fmt.Println(s2) // [0,1]
	}()
}

// data race
func append2() {
	s := make([]int, 0, 1)
	go func() {
		s1 := append(s, 1)
		fmt.Println(s1) // try to update index 1
	}()
	go func() {
		s2 := append(s, 1)
		fmt.Println(s2) // try to update index 1
	}()
}

// create a copy of s
func append3() {
	s := make([]int, 0, 1)
	go func() {
		sCopy := make([]int, len(s), cap(s))
		copy(sCopy, s)
		s1 := append(sCopy, 1)
		fmt.Println(s1)
	}()
	go func() {
		sCopy := make([]int, len(s), cap(s))
		copy(sCopy, s)
		s2 := append(sCopy, 1)
		fmt.Println(s2)
	}()
}
