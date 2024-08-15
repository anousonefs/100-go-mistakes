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

	/*
			1. Channels—Compare whether two channels were created by the same call to make or if both are nil.
			2. Interfaces—Compare whether two interfaces have identical dynamic types and equal dynamic values or if both are nil.
			3. Pointers—Compare whether two pointers point to the same value in memory or if both are nil.
			4. Structs and arrays—Compare whether they are composed of similar types.
		  5. Booleans, numerics, strings Compare whether two of them are equal.
	*/

	// can not compare because operations field is slice
	/* fmt.Printf("use ==: %v\n", cust1 == cust2) */

	/* on average, reflect.DeepEqual is about 100 times slower than ==.  */
	fmt.Printf("use reflect.DeepEqual(): %v\n", reflect.DeepEqual(cust1, cust2))

	/* Running a local benchmark on a slice composed of 100 elements shows that our custom equal method is about 96 times faster than reflect.DeepEqual. */
	fmt.Printf("custom equal method: %v\n", cust1.equal(cust2))
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
