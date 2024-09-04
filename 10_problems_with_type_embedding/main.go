package main

import "fmt"

type Foo struct {
	Bar
	// Bar Bar  (solution)
}

type Bar struct {
	Baz int
}

type Foo2 struct {
	b Bar // we want to hide from the outside
}

// use oop subclassing
func (f Foo2) Baz() int {
	return f.b.Baz
}

func callEmbedType() {
	a := Bar{
		Baz: 44,
	}
	b := Foo{
		a,
	}

	// problem
	fmt.Printf("value1: %v\n", b.Baz)
	fmt.Printf("value2: %v\n", b.Bar.Baz)

	x := Bar{
		Baz: 44,
	}
	y := Foo2{
		x,
	}

	// solution
	fmt.Printf("fix1: %v\n", y.Baz())
}

func main() {
	callEmbedType()
}
