package main

import "fmt"

func main() {
	/* 1. A receiver must be a pointer */
	/*  - If the method needs to mutate the receiver. This rule is also valid if the receiver is a slice and a method needs to append elements: */
	/*  - If the method receiver contains a field that cannot be copied: for example, a type part of the sync package (we will discuss this point in mistake #74, “Copying a sync type”). */
	/* 2. A receiver should be a pointer */
	/*  - If the receiver is a large object. Using a pointer can make the call more efficient, as doing so prevents making an extensive copy. When in doubt about how large is large, benchmarking can be the solution; it’s pretty much impossible to state a specific size, because it depends on many factors. */
	/* 3. A receiver must be a value */
	/* - If we have to enforce a receiver’s immutability. */
	/* - If the receiver is a map, function, or channel. Otherwise, a compilation error occurs. */
	/* 4. A receiver should be a value */
	/* - If the receiver is a slice that doesn’t have to be
	mutated.*/
	/* - If the receiver is a small array or struct that is naturally a value type without mutable fields, such as time.Time. */
	/* - If the receiver is a basic type such as int, float64, or string. */

	structNested()
}

type customer struct {
	data *data
}
type data struct {
	balance float64
}

func (c customer) add(operation float64) {
	c.data.balance += operation
}

/* In this case, we don’t need the receiver to be a pointer to mutate balance. However, for clarity, we may favor using a pointer receiver to highlight that customer as a whole object is mutable. */
func structNested() {
	c := customer{data: &data{
		balance: 100,
	}}
	c.add(50.)
	fmt.Printf("balance: %.2f\n", c.data.balance) // 150.00
}

/* Mixing receiver types */
/* Are we allowed to mix receiver types, such as a struct containing multiple methods, some of which have pointer receivers and others of which have value receivers? The consensus tends toward forbidding it. However, there are some counterexamples in the standard library, for example, time.Time. */
/* The designers wanted to enforce that a time.Time struct is immutable. Hence, most methods such as After, IsZero, and UTC have a value receiver. But to comply with existing interfaces such as encoding.TextUnmarshaler, time.Time has to implement the UnmarshalBinary([]byte) error method, which mutates the receiver given a byte slice. Thus, this method has a pointer receiver. */
/* Consequently, mixing receiver types should be avoided in general but is not forbidden in 100% of cases. */
