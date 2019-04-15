package main

import "fmt"

// 数组（Array）是一段固定长度的连续内存区域。
// 在 Go 语言中，数组从声明时就确定，使用时可以修改数组成员，但是数组大小不可变化。
// 数组声明
var team [3]string
func main() {
	team[0] = "hammer"
	team[1] = "soldier"
	team[2] = "mum"
	fmt.Println(team)
	// 数组的声明与初始化
	var fruits = [...]string{"apple", "pear", "banana"}
	for k, v := range fruits{
		fmt.Println(k, v)
	}
}