package main

import (
	"fmt"
	"testing"
)

func main() {
	i := 0
	go func() { i++ }()
	fmt.Println(i)
}

func TestDataRace(t *testing.T) {
	for i := 0; i < 100; i++ {
		// Actual logic
	}
}

/*
In addition, if a specific file contains tests that lead to data races,
we can exclude it from race detection using the !race build tag

we should bear in mind that running tests with the -race flag for applications using concurrency is highly recommended, if not mandatory. This approach allows us to enable the race detector, which instruments our code to catch potential data races. While enabled, it has a significant impact on memory and performance, so it must be used in specific conditions such as local tests or CI.
*/
