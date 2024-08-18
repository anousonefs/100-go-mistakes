package main

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

func main() {
	/* we’re doing I/O or not, we should first check whether we could implement a whole workflow using bytes instead of strings and avoid the price of additional conversions. */
	demo()
}

func getBytes(reader io.Reader) ([]byte, error) {
	b, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return sanitizeByte(b), nil
}

func sanitizeString(s string) string {
	return strings.TrimSpace(s)
}

func sanitizeByte(b []byte) []byte {
	return bytes.TrimSpace(b)
}

func demo() {
	b := []byte{'a', 'b', 'c'}
	s := string(b)
	b[1] = 'x'
	fmt.Println(s)         // a,b,c
	fmt.Println(string(b)) // a,x,c
}
