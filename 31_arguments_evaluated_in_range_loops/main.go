package main

import "fmt"

func main() {
	/* rangeDemo() */
	/* channelDemo() */
	arrayDemo()
}

func rangeDemo() {
	s := []int{0, 1, 2}
	// iterate 3 times
	for range s {
		s = append(s, 10)
	}
	fmt.Printf("s: %v\n", s)

	// loop never end
	s1 := []int{0, 1, 2}
	for i := 0; i < len(s1); i++ {
		s1 = append(s1, 10)
		fmt.Printf("s1: %v\n", s1)
	}
}

func channelDemo() {
	ch1 := make(chan int, 3)
	go func() {
		println("go 1 start")
		ch1 <- 0
		ch1 <- 1
		ch1 <- 2
		println("go 1 close")
		close(ch1)
	}()

	ch2 := make(chan int, 3)
	go func() {
		println("go 2 start")
		ch2 <- 10
		ch2 <- 11
		ch2 <- 12
		println("go 2 close")
		close(ch2)
	}()

	println("ch := ch1")
	ch := ch1
	for v := range ch {
		fmt.Println(v) // 0, 1, 2
		ch = ch2
	}
	/* The ch = ch2 statement isn’t without effect, though. Because we assigned ch to the second variable, if we call close(ch) */
	/* following this code, it will close the second channel, not the first. */
	for v := range ch {
		fmt.Printf("ch = %v\n", v)
	}

	println("end")
}

func arrayDemo() {
	a := [3]int{0, 1, 2}
	for i, v := range a {
		a[2] = 10
		if i == 2 {
			fmt.Println(v) // 2
		}
	}
	fmt.Printf("a: %v\n", a) // 0,1,10

	/* If we want to print the actual value of the last element */
	/* By accessing the element from its index */
	for i := range a {
		a[2] = 10
		if i == 2 {
			fmt.Println(a[2])
		}
	}

	/* Using an array pointer */
	for i, v := range &a {
		a[2] = 10
		if i == 2 {
			fmt.Println(v)
		}
	}

}
