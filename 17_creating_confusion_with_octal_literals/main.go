package main

import "fmt"

func main() {
	sum := 100 + 0o10
	fmt.Println(sum) // output: 108

	b := 0b100
	fmt.Printf("b: %v\n", b)
	hex := 0xF
	fmt.Printf("hex: %v\n", hex)

	num := 4_402_998_080
	fmt.Printf("num: %v\n", num)
}
