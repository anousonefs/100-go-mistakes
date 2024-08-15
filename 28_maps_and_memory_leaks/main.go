package main

import (
	"crypto/rand"
	"fmt"
	"runtime"
)

func main() {
	map1()
	println("---------")
	map2()

	/*
					   adding n elements to a map and then deleting all the elements means keeping the same number of buckets in memory. So, we must remember that because a Go map can only grow in size, so does its memory consumption. There is no automated strategy to shrink it. If this leads to high memory consumption, we can try different
				options such as:
					   1. forcing Go to re-create the map
		( every hour, we can build a new map, copy all the elements, and release the previous one. The main drawback of this option is that following the copy and until the next garbage collection, we may consume twice the current memory for a short period. )

				     2. using pointers to check if it can be optimized.
	*/

	/* NOTE: a map cannot shrink. */

}

func map1() {
	n := 1_000_000
	m := make(map[int][128]byte)
	printAlloc()
	for i := 0; i < n; i++ {
		m[i] = randBytes()
	}
	/* After adding 1 million elements, the value of runtime.hmap.B equals 18, which means 2^18 = 262,144 buckets */
	printAlloc()
	for i := 0; i < n; i++ {
		delete(m, i)
	}
	runtime.GC()
	printAlloc()
	runtime.KeepAlive(m)
}

func map2() {
	n := 1_000_000
	m := make(map[int]*[128]byte)
	printAlloc()
	for i := 0; i < n; i++ {
		b := randBytes()
		m[i] = &b
	}
	printAlloc()
	for i := 0; i < n; i++ {
		delete(m, i)
	}
	runtime.GC()
	printAlloc()
	runtime.KeepAlive(m)
}

func randBytes() [128]byte {
	var b [128]byte
	_, err := rand.Read(b[:])
	if err != nil {
		panic(err)
	}
	return b
}

func printAlloc() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%.2f MB\n", float64(m.Alloc)/1024/1024)
}
