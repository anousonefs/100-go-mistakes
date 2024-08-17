package main

import "fmt"

type account struct {
	balance float32
}

func main() {

	/* NOTE: In Go, everything we assign is a copy */

	accounts := []account{
		{balance: 100.},
		{balance: 200.},
		{balance: 300.},
	}
	// wrong
	for _, a := range accounts {
		a.balance += 1000
	}
	fmt.Printf("acounts1: %v\n", accounts)

	accounts2 := []account{
		{balance: 100.},
		{balance: 200.},
		{balance: 300.},
	}
	for i := range accounts2 {
		accounts2[i].balance += 1000
	}
	fmt.Printf("accounts2: %v\n", accounts2)

	accounts3 := []account{
		{balance: 100.},
		{balance: 200.},
		{balance: 300.},
	}
	for i := 0; i < len(accounts3); i++ {
		accounts3[i].balance += 1000
	}
	fmt.Printf("accounts3: %v\n", accounts3)

	accountsPointer := []*account{
		{balance: 100.},
		{balance: 200.},
		{balance: 300.},
	}

	/* However, this option has two main downsides. First, it requires updating the slice type, which may not always be possible. Second, if performance is important, we should note that iterating over a slice of pointers may be less efficient for a CPU because of the lack of predictability (we will discuss this point in mistake #91, “Not understanding CPU caches”). */
	for _, a := range accountsPointer {
		a.balance += 1000
	}

	for _, a := range accountsPointer {
		fmt.Printf("accountPointer: %+v\n", a)
	}
}
