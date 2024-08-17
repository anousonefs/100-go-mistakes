package main

import "fmt"

func main() {
	m := map[int]bool{
		0: true,
		1: false,
		2: true}

	for k, v := range m {
		if v {
			m[10+k] = true
		}
	}
	/* The result of this code is unpredictable */
	fmt.Println(m)

	/* m2 := copyMap(m) */
	/* for k, v := range m { */
	/* 	m2[k] = v */
	/* 	if v { */
	/* 		m2[10+k] = true */
	/* 	} */
	/* } */
	/* fmt.Println(m2) */

	/* 1. The data being ordered by keys */
	/* 2. Preservation of the insertion order */
	/* 3. A deterministic iteration order */
	/* 4. An element being produced during the same iteration in which it’s added */

}

func copyMap(original map[int]bool) map[int]bool {
	newMap := make(map[int]bool)
	for k, v := range original {
		newMap[k] = v
	}
	return newMap
}
