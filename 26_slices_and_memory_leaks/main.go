package main

import (
	"fmt"
	"runtime"
)

func main() {
	/* consumeMessages() */
	sliceAndPointer()
}

func consumeMessages() {
	for {
		msg := receiveMessage() // Do something with msg
		storeMessageType(getMessageType(msg))
	}
}

func getMessageType(msg []byte) []byte {
	return msg[:5] // len: 5, cap: 1 000 000
}

// solution
func getMessageType2(msg []byte) []byte {
	dst := make([]byte, 5)
	copy(dst, msg)
	return dst // len: 5, cap: 5
}

func receiveMessage() []byte {
	return []byte{}
}

func storeMessageType(_ []byte) {
}

// slice and pointers

func printAlloc() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%d KB\n", m.Alloc/1024)
}

type Foo struct {
	v []byte
}

func sliceAndPointer() {
	foos := make([]Foo, 1_000)
	printAlloc()
	for i := 0; i < len(foos); i++ {
		foos[i] = Foo{
			v: make([]byte, 1024*1024),
		}
	}
	printAlloc()
	two := keepFirstTwoElementsOnly2(foos)
	runtime.GC()
	printAlloc()
	runtime.KeepAlive(two)
}

func keepFirstTwoElementsOnly(foos []Foo) []Foo {
	return foos[:2]
}

// use slice copy
func keepFirstTwoElementsOnly2(foos []Foo) []Foo {
	res := make([]Foo, 2)
	copy(res, foos)
	return res
}

// make pointer field to nil
func keepFirstTwoElementsOnly3(foos []Foo) []Foo {
	for i := 2; i < len(foos); i++ {
		foos[i].v = nil
	}
	return foos[:2]
}
