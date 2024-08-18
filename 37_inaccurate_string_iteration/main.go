package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "hêllo"
	// wrong
	for i := range s {
		fmt.Printf("position %d: %c\n", i, s[i])
	}
	fmt.Printf("len=%d\n", len(s))
	fmt.Println(utf8.RuneCountInString(s))

	// solution
	for i, r := range s {
		fmt.Printf("position %d: %c\n", i, r)
	}
	println("-------------")

	// solution
	/* converting a string into a slice of runes requires allocating an additional slice and converting the bytes into runes: an O(n) time complexity with n the number of bytes in the string. */
	runes := []rune(s)
	for i, r := range runes {
		fmt.Printf("position %d: %c\n", i, r)
	}
	println("------------")
	r := []rune(s)[4]
	fmt.Printf("%c\n", r) // o
}
