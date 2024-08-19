package main

import "io"

func main() {

}

type locator interface {
	getCoordinates(address string) (lat, lng float32, err error) // lat first then lng
}

type Customer struct {
	name string
}

// in this case we should favor not using named result parameters.
func StoreCustomer(customer Customer) (err error) {
	return
}

/*
	naked returns (returns without arguments):
	they are considered acceptable in short functions;

otherwise, they can harm readability because the reader must remember the outputs throughout the entire function.
We should also be consistent within the scope of a function,
using either only naked returns or only returns with arguments.
*/
func ReadFull(r io.Reader, buf []byte) (n int, err error) {
	for len(buf) > 0 && err == nil {
		var nr int
		nr, err = r.Read(buf)
		n += nr
		buf = buf[nr:]
	}
	return
}
