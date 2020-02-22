package main

import (
	"fmt"
)

func main() {
	n1 := 101
	fmt.Printf("十进制数：%d\n", n1)
	fmt.Printf("二进制数：%b\n", n1)
	fmt.Printf("八进制数：%o\n", n1)
	fmt.Printf("十六进制数：%x\n", n1)
	n2 := 0777
	fmt.Printf("十进制数：%d\n", n2)
	fmt.Printf("八进制数：%o\n", n2)
	// 十六进制
	n3 := 0x123450f
	fmt.Printf("十进制数：%d\n", n3)
	fmt.Printf("十六进制数：%x\n", n3)
	// 查看变量类型 %T
	fmt.Printf("十进制数：%T\n", n3)
	fmt.Printf("字符串类型：%#v\n", "Hello Wolrd")
}
