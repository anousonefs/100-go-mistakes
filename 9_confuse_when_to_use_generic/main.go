package main

import (
	"fmt"
	"sort"
)

func main() {

	/* When it makes our code more complex—Generics are never mandatory, and as Go developers, we have lived without them for more than a decade. If we’re writing generic functions or structures and we figure out that it doesn’t make our code clearer, we should probably reconsider our decision for that particular use case */

	m := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	m2 := map[int]string{
		1: "haha",
		2: "a",
		3: "bb",
	}
	keysAny, err := getKeysByAny(m)
	if err != nil {
		panic(err)
	}
	fmt.Printf("keysAny: %v\n", keysAny)

	keysGeneric := getKeysByGeneric(m)
	fmt.Printf("keysGeneric: %v\n", keysGeneric)

	keysGeneric2 := getKeysByGeneric(m2)
	fmt.Printf("keysGeneric2: %v\n", keysGeneric2)

	// sort demo

	s := SliceFn[int]{
		S: []int{3, 2, 1, 5, 8, 4},
		Compare: func(a, b int) bool {
			return a < b
		},
	}
	sort.Sort(s)
	fmt.Printf("sorted_int: %v\n", s.S)

	s2 := SliceFn[string]{
		S: []string{"sone", "a", "Anousone", "house", "golang", "Flutter"},
		Compare: func(a, b string) bool {
			return a < b
		},
	}
	sort.Sort(s2)
	fmt.Printf("sorted_string: %v\n", s2.S)
}

// if we use any it can error in runtime
func getKeysByAny(m any) ([]any, error) {
	switch t := m.(type) {
	default:
		return nil, fmt.Errorf("unknown type: %T", t)
	case map[string]int:
		var keys []any
		for k := range t {
			keys = append(keys, k)
		}
		return keys, nil
	case map[int]string:
		var keys []any
		for k := range t {
			keys = append(keys, k)
		}
		return keys, nil
	}
}

// use generic for safe in complie time
func getKeysByGeneric[K comparable, V any](m map[K]V) []K {
	var keys []K
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

type Node[T any] struct {
	Val  T
	next *Node[T]
}

func (n *Node[T]) Add(next *Node[T]) {
	n.next = next
}

/* If we want to use generics with methods, it’s the receiver that needs to be a type parameter */

/* type Foo struct {} */
/**/
/* func (Foo) bar[T any](t T) { */
/* } */
/**/

type SliceFn[T any] struct {
	S       []T
	Compare func(T, T) bool
}

func (s SliceFn[T]) Len() int {
	return len(s.S)
}

func (s SliceFn[T]) Less(i, j int) bool {
	return s.Compare(s.S[i], s.S[j])
}
func (s SliceFn[T]) Swap(i, j int) {
	s.S[i], s.S[j] = s.S[j], s.S[i]
}
