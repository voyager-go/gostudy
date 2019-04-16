package main

import "fmt"

// 声明结构体内嵌
type BasicColor struct {
	R, G, B float32
}
type Color struct {
	Basic BasicColor
	Alpha float32
}
// 内嵌写法
type NewColor struct {
	BasicColor
	Alpha float32
}
func main() {
	var c Color
	c.Basic.R = 1
	c.Basic.G = 1
	c.Basic.B = 0
	c.Alpha   = 1
	fmt.Printf("%+v", c)


	//-----
	var co NewColor
	co.R = 1
	co.G = 1
	co.B = 0
	fmt.Printf("%+v", co)
}