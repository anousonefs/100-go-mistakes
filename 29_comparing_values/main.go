package main

import (
	"fmt"
	"reflect"
)

type customer struct {
	id         string
	operations []float64
}

func main() {
	cust1 := customer{id: "x", operations: []float64{1.}}
	cust2 := customer{id: "x", operations: []float64{1.}}

	/*To compare types in Go, you can use the == and != operators if two types are comparable: Booleans,
	  numerals, strings, pointers, channels, and structs are composed entirely of comparable types. Otherwise, you can either use reflect.DeepEqual and pay the price of reflection or use custom implementations and libraries.*/

	// can not compare because operations field is slice
	/* fmt.Printf("use ==: %v\n", cust1 == cust2) */

	/* on average, reflect.DeepEqual is about 100 times slower than ==.  */
	fmt.Printf("use reflect.DeepEqual(): %v\n", reflect.DeepEqual(cust1, cust2))

	/* Running a local benchmark on a slice composed of 100 elements shows that our custom equal method is about 96 times faster than reflect.DeepEqual. */
	fmt.Printf("custom equal method: %v\n", cust1.equal(cust2))

	/* NOTE: bytes.Compare function to compare two slices of bytes. Before implementing a custom method, we need to make sure we don’t reinvent the wheel.*/

}

func (a customer) equal(b customer) bool {
	if a.id != b.id {
		return false
	}
	if len(a.operations) != len(b.operations) {
		return false
	}
	for i := 0; i < len(a.operations); i++ {
		if a.operations[i] != b.operations[i] {
			return false
		}
	}
	return true
}
