package main

import "fmt"

// 结构体内嵌初始化时，将结构体内嵌的类型作为字段名像普通结构体一样进行初始化
// 车轮
type Wheel struct {
	Size int  // 尺寸
}
// 引擎
type Engine struct {
	Power int		// 功率
	Type  string	// 类型
}
// 车
type Car struct {
	Wheel
	Engine
	// 初始化内嵌匿名结构体
	// 中控
	CenterControl struct{
		Display string	// 显示器
		USB string		// usb口
	}
}

func main() {
	c := Car{
		// 初始化轮子
		Wheel: Wheel{
			Size: 18,
		},
		Engine: Engine{
			Type: "1.5T",
			Power: 143,
		},
		CenterControl: struct {
			Display string
			USB     string
		}{Display: "AOC", USB: "3.0"},
	}
	fmt.Println("%+v\n", c)
}