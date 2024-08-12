package main

import (
	"fmt"
	"math"
)

/* Go provides a package to deal with large numbers: math/big. */

func main() {
	var counter int32 = math.MaxInt32
	fmt.Printf("max=%d\n", counter)
	counter++
	/* counter = Inc32(counter) */
	fmt.Printf("min=%d\n", counter)

	// add
	var a uint16 = 30000
	var b uint16 = 35000
	result := AddUint16(a, b)
	fmt.Printf("add result: %v\n", result)

	// multiply
	var x int16 = math.MinInt16
	var y int16 = 2
	res := MultiplyInt16(x, y)
	fmt.Printf("multiply result: %v\n", res)
}

func Inc32(counter int32) int32 {
	if counter == math.MaxInt32 {
		panic("int32 overflow")
	}
	return counter + 1
}

func AddUint16(a, b uint16) uint16 {
	if a > math.MaxUint16-b {
		panic("int overflow")
	}
	return a + b
}

func MultiplyInt16(a, b int16) int16 {
	if a == 0 || b == 0 {
		return 0
	}
	result := a * b
	if a == 1 || b == 1 {
		return result
	}
	if a == math.MinInt16 || b == math.MinInt16 {
		panic("integer overflow")
	}
	if result/b != a {
		panic("integer overflow")
	}
	return result
}
