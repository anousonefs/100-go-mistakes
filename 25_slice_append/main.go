package main

import "fmt"

func main() {
	s1 := []int{1, 2, 3}
	s2 := s1[1:2]
	s3 := append(s2, 10)
	// s1 = 1,2,10, s2 = 2, s3 = 2, 10
	fmt.Printf("s1: %v, s2: %v, s3: %v\n\n", s1, s2, s3)

	// If the resulting slice has a length smaller than its capacity
	// append can mutate the original slice
	updateSlice() // wrong

	updateSlice2() // correct: use slice copy
	updateSlice3() // correct: use full slice expression
	updateSlice4()
}

func updateSlice() {
	s := []int{1, 2, 3}
	f(s[0:2])                                    // len: 2, cap: 3
	fmt.Printf("mutate original slice: %v\n", s) // [9 2 10]
}

// update just first 2 elements
func updateSlice2() {
	s := []int{1, 2, 3, 4, 5}
	sCopy := make([]int, 2)
	copy(sCopy, s) // len: 2, cap: 2
	f(sCopy)
	result := append(sCopy, s[2:]...)
	fmt.Printf("slice copy: %v\n", result) // [9 2 3 4 5]
}

func updateSlice3() {
	s := []int{1, 2, 3}
	f(s[:2:2])                                   // len: 2, cap: 2
	fmt.Printf("full slice expression: %v\n", s) // [9 2 3]
}

func f(s []int) {
	fmt.Printf("s: %v, cap: %v\n", s, cap(s))
	s[0] = 9
	_ = append(s, 10)
}

func updateSlice4() {
	s := []int{1, 2, 3}
	sCopy := make([]int, 2)
	copy(sCopy, s)
	sCopy = f3(sCopy) // whan f3() append effect
	result := append(sCopy, s[2])
	fmt.Printf("update v4: %v\n", result) // [9, 2, 10, 3]
}

func f3(s []int) []int {
	s[0] = 9
	return append(s, 10)
}
