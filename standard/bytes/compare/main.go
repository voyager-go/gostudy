package main

import (
	"bytes"
	"fmt"
	"sort"
)

func main() {
	var a, b []byte
	if bytes.Compare(a, b) < 0 {
		fmt.Println("a less b")
	}
	if bytes.Compare(a, b) <= 0 {
		fmt.Println("a less or equal b")
	}
	if bytes.Compare(a, b) > 0 {
		fmt.Println("a greater b")
	}
	if bytes.Equal(a, b) {
		fmt.Println("a equal b")
	}
	if !bytes.Equal(a, b) {
		fmt.Println("a not equal b")
	}
	if bytes.Equal(a, nil) {
		fmt.Println("a equal nil slice")
	}

	// Binary search to find a matching byte slice.
	var needle = []byte{12}
	//fmt.Println("needle:", needle)
	var haystack = [][]byte{
		{1},
		{2},
		{3},
		{4},
		{12},
	}
	index := sort.Search(len(haystack), func(i int) bool {
		return bytes.Compare(haystack[i], needle) >= 0
	})
	fmt.Println(index)
	if index < len(haystack) && bytes.Equal(haystack[index], needle) {
		fmt.Println(index)
	}
}
