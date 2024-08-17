package main

import "os"

func main() {
}

// wrong
/* the defer calls are executed not during each loop iteration
but when the readFiles function returns. */
func readFiles(ch <-chan string) error {
	for path := range ch {
		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()
		// Do something with file
	}
	return nil
}

// solution
func readFiles2(ch <-chan string) error {
	for path := range ch {
		if err := readFile2(path); err != nil {
			return err
		}
	}
	return nil
}

// easy for test
func readFile2(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	// Do something with file
	return nil
}

// solution
func readFiles3(ch <-chan string) error {
	for path := range ch {
		err := func() error {
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()
			return nil
		}()
		if err != nil {
			return err
		}
	}
	return nil
}
