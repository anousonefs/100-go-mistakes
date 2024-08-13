package main

import "fmt"

func main() {
	copy1() // incorrect
	copy2() // correct
	copy3() // correct
}

func copy1() {
	src := []int{0, 1, 2}
	var dst []int
	copy(dst, src)
	fmt.Printf("copy1(): %v\n", dst)
}

// easy to read
func copy2() {
	src := []int{0, 1, 2}
	dst := make([]int, len(src))
	copy(dst, src)
	fmt.Printf("copy2(): %v\n", dst)
}

func copy3() {
	src := []int{0, 1, 2}
	dst := append([]int(nil), src...)
	fmt.Printf("copy3(): %v\n", dst)
}
