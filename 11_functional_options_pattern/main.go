package main

import (
	builder "11_functional_options_pattern/builder_pattern"
	functional "11_functional_options_pattern/functional_options_pattern"
)

func main() {
	NewServer("localhost", config{}) // bad

	// builder pattern
	b := builder.ConfigBuilder{}
	b.Port(8080)
	cfg, err := b.Build()
	if err != nil {
		panic(err)
	}
	builder.NewServer("localhost", cfg) // look good

	// functional options pattern
	/* s, err := functional.NewServer("localhost") */
	s, err := functional.NewServer("localhost", functional.WithPort(8084))
	if err != nil {
		panic(err)
	}
	_ = s
}

// common behavior
type config struct {
	port int
}

func NewServer(addr string, cfg config) {
	println("new server")
}
