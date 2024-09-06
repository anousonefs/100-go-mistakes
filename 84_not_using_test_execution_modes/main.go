package main

import "testing"

// go test -parallel 16 .
// By default, the maximum number of tests that can run simultaneously equals the GOMAXPROCS value
func main() {

}

func TestA(t *testing.T) {
	t.Parallel()
	// ...
}
func TestB(t *testing.T) {
	t.Parallel()
	// ...
}

func TestC(t *testing.T) {
	// ...
}

// we can use the - shuffle flag to randomize tests.
// go test -shuffle=on -v .
// go test -shuffle=1636399552801504000 -v .   run same order
