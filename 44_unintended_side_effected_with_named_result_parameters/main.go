package main

import (
	"context"
	"errors"
)

func main() {

}

type loc struct {
}

func (f loc) validateAddress(_ string) bool {
	return true
}

func (l loc) getCoordinates(ctx context.Context, address string) (
	lat, lng float32, err error) {
	isValid := l.validateAddress(address)
	if !isValid {
		return 0, 0, errors.New("invalid address")
	}
	// wrong
	/* if ctx.Err() != nil { */
	/* 	return 0, 0, err */
	/* } */
	/* err in this example shadows the result variable. */
	if err := ctx.Err(); err != nil {
		return 0, 0, err
	}
	// using named result parameters doesn’t necessarily mean using naked returns.
	// sometimes we can just use named result parameters to make a signature clearer.
	/* if err = ctx.Err(); err != nil { */
	/* 	return */
	/* } */
	return
}
