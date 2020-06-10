package main

import "bytes"
import "fmt"

func main() {
	var a = []byte{1, 2, 2, 3, 3, 2, 3, 2, 3, 55, 11}
	var b = []byte{2, 3}
	// Count计算s中有多少个不重叠的sep子切片
	fmt.Printf("共有%d个不重叠的子切片\n", bytes.Count(a, b))
	// 子切片sep在s中第一次出现的位置，不存在则返回-1
	fmt.Println("子切片第一次出现的位置：", bytes.Index(a, b))
	// 字符c在s中第一次出现的位置，不存在则返回-1
	fmt.Println("字符在切片中第一次出现的位置：", bytes.IndexByte(a, 55))
	var t = []byte("Golang hello world!")
	fmt.Println(string(bytes.Title(t)))
}
