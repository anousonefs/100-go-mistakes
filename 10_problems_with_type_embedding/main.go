package main

import "fmt"

type Foo struct {
	Bar
}
type Bar struct {
	Baz int
}

func callEmbedType() {
	a := Bar{
		Baz: 44,
	}
	b := Foo{
		a,
	}

	fmt.Printf("value1: %v\n", b.Baz)
	fmt.Printf("value2: %v\n", b.Bar.Baz)
}
