package main

import "fmt"

func main() {
	/* loop1() */
	/* loop2() */
	loop3()
}

// wrong
func loop1() {
	s := []int{1, 2, 3}
	for _, i := range s {
		go func() {
			fmt.Print(i)
		}()
	}
}

func loop2() {
	s := []int{1, 2, 3}
	for _, i := range s {
		val := i
		go func() {
			fmt.Print(val)
		}()
	}
}

func loop3() {
	s := []int{1, 2, 3}
	for _, i := range s {
		go func(a int) {
			fmt.Print(a)
		}(i)
	}
}
