package main

import "fmt"

// Go语言的关键字 type 可以将各种基本类型定义为自定义类型，基本类型包括整型、字符串、布尔等。结构体是一种复合的基本类型，通过 type 定义为自定义类型后，使结构体更便于使用。
type Pointer struct {
	X int
	Y int
}
// 创建指针类型的结构体
type Player struct {
	Name string
	HealthPointer int
	MagicPointer  int
}

// 取结构体的地址实例化
type Command struct {
	Name string
	Var  *int
	Comment string
}
func main() {
	// 基本的实例化形式
	var p Pointer
	p.X = 10
	p.Y = 20
	fmt.Println(p)
	tank := new(Player)
	tank.Name = "Canon"
	tank.HealthPointer = 300
	fmt.Println(tank)

	var version int = 1
	cmd := &Command{}
	cmd.Name = "version"
	cmd.Var  = &version
	cmd.Comment = "show version"
	fmt.Println(cmd)
}