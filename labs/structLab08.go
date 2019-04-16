package main

import "fmt"

// 可飞行的
type Flying struct {}

func (f *Flying) Fly() {
	fmt.Println("can fly")
}

// 可行走的
type Walkable struct{}

func (w *Walkable) Walk() {
	fmt.Println("can walk")
}

// 人类
type Human struct {
	Walkable // 可以行走
}

// 鸟类
type Bird struct {
	Flying // 可以飞行
	Walkable // 可以行走
}

func main() {
	// 实例化鸟类
	b := new(Bird)
	b.Fly()
	b.Walk()

	// 实例化人类
	p := new(Human)
	p.Walk()
}