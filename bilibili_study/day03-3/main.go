package main

import (
	"fmt"
	"unsafe"
)
// 内存对齐

type TaskOne struct {
	a int
	b int64
	c bool
	d complex64
	e string
	f byte
	g int32
	h int8
}

type TaskTwo struct {
	a int
	c bool
	f byte
	h int8
	g int32
	b int64
	e string
	d complex64
}
func main() {
	// task one size => 64, align => 8
	fmt.Printf("task one size => %d, align => %d\n", unsafe.Sizeof(TaskOne{}), unsafe.Alignof(TaskOne{}))
	// task two size => 48, align => 8
	fmt.Printf("task two size => %d, align => %d\n", unsafe.Sizeof(TaskTwo{}), unsafe.Alignof(TaskTwo{}))
}
