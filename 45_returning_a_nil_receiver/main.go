package main

import (
	"errors"
	"fmt"
	"log"
	"strings"
)

func main() {
	/* pointerNil() */
	customer := Customer{Age: 33, Name: "John"}
	/* if err := customer.ValidateWrong(); err != nil { */
	/* 	log.Fatalf("ValidateWrong is invalid: %v", err) */
	/* } */
	if err := customer.ValidateFixed(); err != nil {
		log.Fatalf("ValidateFixed is invalid: %v", err)
	}
	println("pass")
}

type MultiError struct {
	errs []string
}

func (m *MultiError) Add(err error) {
	m.errs = append(m.errs, err.Error())
}

// go allow nil receiver
func (m *MultiError) Error() string {
	return strings.Join(m.errs, ";")
}

type Customer struct {
	Age  int
	Name string
}

// wrong return nil pointer
func (c Customer) ValidateWrong() error {
	var m *MultiError
	if c.Age < 0 {
		m = &MultiError{}
		m.Add(errors.New("age is negative"))
	}
	if c.Name == "" {
		if m == nil {
			m = &MultiError{}
		}
		m.Add(errors.New("name is nil"))
	}
	// return nil receiver
	return m
}

// fixed
func (c Customer) ValidateFixed() error {
	var m *MultiError
	if c.Age < 0 {
		m = &MultiError{}
		m.Add(errors.New("age is negative"))
	}
	if c.Name == "" {
		if m == nil {
			m = &MultiError{}
		}
		m.Add(errors.New("name is nil"))
	}
	/* when we have to return an interface, we should return not a nil pointer but a nil value directly.
	   Generally, having a nil pointer isn’t a desirable state and means a probable bug. */
	if m != nil {
		return m
	}
	// return nil value
	return nil
}

// ----------------------------
// pointer receiver can be nil.
type Foo struct {
	a string
}

// allow nil foo
func (foo *Foo) Bar() string {
	//println(foo.a) // error
	return "bar"
}

// allow nil foo
func Bar(foo *Foo) string {
	return "bar"
}

/*
		We know that passing a nil pointer to a function is valid.
	  Therefore, using a nil pointer as a receiver is also valid.
*/
func pointerNil() {
	var foo *Foo
	fmt.Println(foo.Bar())
}
