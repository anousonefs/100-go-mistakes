package main

import (
	"fmt"
	"strings"
)

func main() {

	s := []string{"anousone", "sone", "freestyle"}
	/* BenchmarkConcatV1-4    16      72291485 ns/op */
	fmt.Printf("concat v1: %v\n", concatV1(s))

	/* BenchmarkConcatV2-4  1188        878962 ns/op */
	fmt.Printf("concat v2: %v\n", concatV2(s))

	/* BenchmarkConcatV3-4  5922        190340 ns/op */
	/* strings.Builder solution is faster from the moment we have to concatenate more than about five strings.  */
	fmt.Printf("concat v3: %v\n", concatV3(s))
}

/* we forget one of the core characteristics of a string: its immutability */
/* this function reallocates a new string in memory, which significantly impacts the performance*/
func concatV1(values []string) string {
	s := ""
	for _, value := range values {
		s += value
	}
	return s
}

/* we constructed the resulting string by calling the WriteString method that appends the content of value to its internal buffer, hence minimizing memory copying. */
/* 1. A byte slice using Write */
/* 2. A single byte using WriteByte  */
/* 3. A single rune using WriteRune */

/*
	two impacts.

1.  this struct shouldn’t be used concurrently, as the calls to append would lead to race conditions.
2.  something that we saw in mistake #21, “Inefficient slice initialization”: if the future length of a slice is already known, we should preallocate it. For that purpose, strings.Builder exposes a method Grow(n int) to guarantee space for another n bytes.
*/
//without preallocation
func concatV2(values []string) string {
	sb := strings.Builder{}
	for _, value := range values {
		_, _ = sb.WriteString(value)
	}
	return sb.String()
}

// preallocation
func concatV3(values []string) string {
	total := 0
	for i := 0; i < len(values); i++ {
		total += len(values[i])
	}
	sb := strings.Builder{}
	sb.Grow(total)
	for _, value := range values {
		_, _ = sb.WriteString(value)
	}
	return sb.String()
}
