package main

import (
	"fmt"
	"io"
	"strings"
	"unicode"
)

// LowerCaseReader is a custom reader that reads from another reader,
// converts all characters to lowercase, and maintains filtering logic.
type LowerCaseReader struct {
	reader io.Reader
}

// Read reads from the underlying reader, converts characters to lower case,
// and returns the modified data.
func (lcr *LowerCaseReader) Read(p []byte) (n int, err error) {
	// Create a buffer to store the original data
	buf := make([]byte, len(p))
	n, err = lcr.reader.Read(buf)
	if err != nil {
		return n, err
	}

	// Process the buffer: convert to lowercase
	var j int
	for i := 0; i < n; i++ {
		char := buf[i]

		// Use unicode.ToLower to convert character to lowercase
		lowerChar := byte(unicode.ToLower(rune(char)))

		// Only write the character at the even positions
		if i%2 == 0 { // Filter: Keep characters at even indices
			p[j] = lowerChar
			j++
		}
	}

	// Return the number of processed (lowercased and filtered) bytes
	return j, err
}

// The main function is for demonstrating how the Reader works (not strictly necessary for test)
func IoDemo() {
	lcr := &LowerCaseReader{reader: strings.NewReader("aBcDeFgHiJ")}
	buf := make([]byte, 10) // Create a buffer to read into

	n, err := lcr.Read(buf)
	if err != nil && err != io.EOF {
		fmt.Println("Error reading:", err)
		return
	}

	fmt.Printf("Read %d bytes: %q\n", n, buf[:n]) // Output: "acegi"
}
