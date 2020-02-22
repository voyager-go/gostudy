package main

import "fmt"

const  (
	n1 = iota
	n2
	_
	n3
)
// 插队
const (
	a1 = iota
	a2 = 100
	a3 = iota
	a4
)
func main() {
	fmt.Println(n1, n2, n3)
	fmt.Println(a1, a2, a3, a4)
}

