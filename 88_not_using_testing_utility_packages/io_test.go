package main

import (
	"fmt"
	"io"
	"strings"
	"testing"
	"testing/iotest"
)

func TestLowerCaseReader(t *testing.T) {
	// TestReader will expect the output to be "acegi".
	err := iotest.TestReader(
		&LowerCaseReader{reader: strings.NewReader("aBcDeFgHiJ")},
		[]byte("acegi"),
	)
	if err != nil {
		t.Fatal(err)
	}
}

func TestLowerCaseReader2(t *testing.T) {
	lcr := &LowerCaseReader{reader: strings.NewReader("aBcDeFgHiJ")}
	buf := make([]byte, 10) // Create a buffer to read into

	n, err := lcr.Read(buf)
	if err != nil && err != io.EOF {
		fmt.Println("Error reading:", err)
		return
	}

	fmt.Printf("Read %d bytes: %q\n", n, buf[:n]) // Output: "acegi"

	want := string([]byte("acegi"))

	result := string(buf[:n])
	if want != result {
		t.Error("does not match")
	}
}
