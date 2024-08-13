package main

func main() {

}

type Bar struct {
	value int
}

type Foo struct {
	value string
}

func fooToBar(_ Foo) Bar {
	return Bar{}
}

// slow
/* Every time the backing array is full, Go creates another array by doubling its capacity */
func convert(foos []Foo) []Bar {
	bars := make([]Bar, 0)
	for _, foo := range foos {
		bars = append(bars, fooToBar(foo))
	}
	return bars
}

// fast and easy to read
func convert2(foos []Foo) []Bar {
	bars := make([]Bar, 0, len(foos))
	for _, foo := range foos {
		bars = append(bars, fooToBar(foo))
	}
	return bars
}

// faster because do not use append function
func convert3(foos []Foo) []Bar {
	bars := make([]Bar, len(foos))
	for i, foo := range foos {
		bars[i] = fooToBar(foo)
	}
	return bars
}
