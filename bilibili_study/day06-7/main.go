// @Title
// @Description
// @Author       幼儿猿
// @Datetime     2020/2/27 12:59 上午
package main

import "fmt"

// 空接口
// 空接口是没有定义任何方法的接口
// 因此任何类型都实现了空接口
// 空接口的变量可以存储任意类型的变量
//type X interface{}

func main() {
	//var x X
	var x interface{}
	var m string = "Hello"
	x = m
	fmt.Printf("%T, %v\n", x, x)
	var b bool = true
	x = b
	fmt.Printf("%T, %v\n", x, x)
}
