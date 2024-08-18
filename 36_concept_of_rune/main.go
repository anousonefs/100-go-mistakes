package main

import "fmt"

func main() {
	s := "hello"
	fmt.Println(len(s)) // 5

	s1 := "汉"
	fmt.Println(len(s1)) // 3

	s2 := string([]byte{0xE6, 0xB1, 0x89})
	fmt.Printf("%s\n", s2)

	/*
			   1. A charset is a set of characters, whereas an encoding describes how to translate a charset into binary.
		     2. In Go, a string references an immutable slice of arbitrary bytes.
		     3. Go source code is encoded using UTF-8. Hence, all string literals are UTF-8 strings. But because a string can contain arbitrary bytes, if it’s obtained from somewhere else (not the source code), it isn’t guaranteed to be based on the UTF-8 encoding.
		     4. A rune corresponds to the concept of a Unicode code point, meaning an item represented by a single value.
		     5. Using UTF-8, a Unicode code point can be encoded into 1 to 4 bytes.
		     6. Using len on a string in Go returns the number of bytes, not the number of runes.
	*/
}
