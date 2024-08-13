package main

import "fmt"

func main() {
	s1 := make([]int, 3, 6)
	fmt.Printf("s1: %v, len: %d, cap: %d\n", s1, len(s1), cap(s1))

	s2 := s1[1:3]
	fmt.Printf("s2: %v, len: %d, cap: %d\n", s2, len(s2), cap(s2))
	s2 = append(s2, 4)

	println("=> s2 = append(s2, 4)")
	fmt.Printf("s1: %v, len: %d, cap: %d\n", s1, len(s1), cap(s1))
	fmt.Printf("s2: %v, len: %d, cap: %d\n", s2, len(s2), cap(s2))
	s1 = append(s1, 2)
	s2 = append(s2, 3)
	s2 = append(s2, 9)
	s2 = append(s2, 8)
	println("=> s2 full capacity")
	fmt.Printf("s1: %v, len: %d, cap: %d\n", s1, len(s1), cap(s1))
	fmt.Printf("s2: %v, len: %d, cap: %d\n", s2, len(s2), cap(s2))
}
