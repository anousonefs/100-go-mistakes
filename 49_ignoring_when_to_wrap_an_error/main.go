package main

import (
	"errors"
	"fmt"
)

/*
1. If we need to mark an error, we should create a custom error type.
2. if we just want to add extra context, we should use fmt.Errorf with the %w directive as it doesn’t require creating a new error type.
3. error wrapping creates potential coupling as it makes the source error available for the caller. If we want to prevent it, we shouldn’t use error wrapping but error transformation, for example, using fmt.Errorf with the %v directive.
*/

func main() {

}

func bar() error {
	return errors.New("bar err")
}

func Foo() error {
	err := bar()
	if err != nil {
		return err
	}
	return nil
}

func Foo2() error {
	err := bar()
	if err != nil {
		return fmt.Errorf("Foo2: %w", err) // wrap error: add extra context, the caller can get source error
	}
	return nil
}

func Foo3() error {
	err := bar()
	if err != nil {
		return fmt.Errorf("Foo2: %v", err) // transform error
	}
	return nil
}

func Foo4() error {
	err := bar()
	if err != nil {
		return BarError{Err: err} // custom error: add extra context, marking an error and source error available
	}
	return nil
}

type BarError struct {
	Err error
}

func (b BarError) Error() string {
	return "bar failed:" + b.Err.Error()
}
