package main

import "fmt"
// 匿名
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
// 多个常量在一行
// 遇到 const iota 重置为 0
// 每增加一行常量复制 iota 新增 1
const (
	b1, b2 = iota + 1, iota + 2 // 1, 2
	b3, b4 = iota + 1, iota + 2 // 2, 3
)
// 定义数量级
const (
	_  = iota
	KB = 1 << (10 * iota) // (10 * 1)    这里 << 表示向左移10 位 10000000000 二进制转换十进制为 1024
	MB = 1 << (10 * iota) // (10 * 2)    左移20 位
	GB
	TB
	PB
)
func main() {
	fmt.Println(n1, n2, n3)
	fmt.Println(a1, a2, a3, a4)
	fmt.Println(b1, b2, b3, b4)
	fmt.Println(PB)
}

