package main

import "fmt"

var ten int = 11

func main() {
	if ten > 10 {
		fmt.Println(">10")
	}else {
		fmt.Println("<=10")
	}
	step := 2
	for ; step > 0; step-- {
		fmt.Println(step)
	}
	// 遍历数组、切片——获得索引和元素
	for key, value := range []int{1, 2, 3, 4} {
		fmt.Printf("key:%d value:%d\n", key, value)
	}
	// 遍历字符串——获得字符
	strA := "hello 你好"
	for key, value := range strA {
		fmt.Printf("key:%d value:0x%x\n", key, value)
	}
	// 遍历map——获得map的键和值
	m := map[string]int{
		"hello": 100,
		"world": 200,
	}
	for key, value := range m {
		fmt.Println(key, value)
	}
	// 遍历通道（channel）——接收通道数据
	ch := make(chan int)
	go func() {
		ch <- 1
		ch <- 2
		ch <- 3
		close(ch)
	}()
	for v := range ch {
		fmt.Println(v)
	}
	// switch
	// Go 语言改进了 switch 的语法设计，避免人为造成失误。Go 语言的 switch 中的每一个 case 与 case 间是独立的代码块，不需要通过 break 语句跳出当前 case 代码块以避免执行到下一行
	var a = "hello"
	switch a {
	case "hello":
		fmt.Println(1)
	case "world":
		fmt.Println(2)
	default:
		fmt.Println(0)
	}
	var b = "mum"
	switch b {
	case "mum", "daddy":
		fmt.Println("family")
	}
}
