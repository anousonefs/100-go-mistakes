package main

import (
	"errors"
	"fmt"
)

func main() {
	/* closureDemo() */
	/* f() */ // status is empty string
	/* f2()      // pointer */
	/* f3()      // closure */
	/* deferMethodDemo() */
	deferMethodPointerDemo()
}

const (
	StatusSuccess  = "success"
	StatusErrorFoo = "error_foo"
	StatusErrorBar = "error_bar"
)

func f() error {
	var status string
	defer notify(status)           // status is empty string
	defer incrementCounter(status) // status is empty string
	if err := foo(); err != nil {
		status = StatusErrorFoo
		return err
	}
	if err := bar(); err != nil {
		status = StatusErrorBar
		return err
	}
	status = StatusSuccess
	return nil
}

func f2() error {
	var status string
	defer notifyPointer(&status)           // status is empty string
	defer incrementCounterPointer(&status) // status is empty string
	if err := foo(); err != nil {
		status = StatusErrorFoo
		return err
	}
	if err := bar(); err != nil {
		status = StatusErrorBar
		return err
	}
	status = StatusSuccess
	return nil
}

func f3() error {
	var status string
	defer func() {
		notify(status)           // status is latest
		incrementCounter(status) // status is latest
	}()
	if err := foo(); err != nil {
		status = StatusErrorFoo
		return err
	}
	if err := bar(); err != nil {
		status = StatusErrorBar
		return err
	}
	status = StatusSuccess
	return nil
}

func notify(status string) {
	fmt.Printf("notify: %v\n", status)
}

func incrementCounter(status string) {
	fmt.Printf("incrementCounter: %v\n", status)
}

func notifyPointer(status *string) {
	fmt.Printf("notify: %s\n", *status)
}

func incrementCounterPointer(status *string) {
	fmt.Printf("incrementCounter: %s\n", *status)
}

func foo() error {
	return errors.New("foo err")
}
func bar() error {
	return nil
}

func closureDemo() {
	i := 0
	j := 0
	defer func(i int) {
		fmt.Println(i, j) // 0, 1
	}(i)
	i++
	j++
}

/* Pointer and value receivers */

type Struct struct {
	id string
}

// value receivers
func (f Struct) print() {
	fmt.Println(f.id)
}

// pointer receivers
func (f *Struct) printPointer() {
	fmt.Println(f.id)
}

func deferMethodDemo() {
	s := Struct{id: "foo"}
	defer s.print() // foo
	s.id = "bar"
}

func deferMethodPointerDemo() {
	s := Struct{id: "foo"}
	defer s.printPointer() // bar
	s.id = "bar"
}
