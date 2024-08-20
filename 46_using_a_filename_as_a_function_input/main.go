package main

import (
	"bufio"
	"io"
	"os"
	"strings"
	"testing"
)

func main() {

}

// wrong: hard to reuse and test
func countEmptyLinesInFile(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
	}
	return 0, nil
}

// solution: reuseable and easy to test
func countEmptyLines(reader io.Reader) (int, error) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
	}
	return 0, nil
}

func TestCountEmptyLines(t *testing.T) {
	_, _ = countEmptyLines(
		strings.NewReader(
			`foo
       bar
       baz`),
	)
}
