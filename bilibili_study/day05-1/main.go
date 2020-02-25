package main

import "fmt"

// 方法与接收者

type Person struct {
	name, gender string
	age          int
}

// NewPerson 是 Person类型的构造函数
func NewPerson(name, gender string, age int) *Person {
	return &Person{
		name:   name,
		gender: gender,
		age:    age,
	}
}

// Dream 是位 Person 类型定义方法
func (p Person) Dream() {
	fmt.Printf("%s 有一个不写代码的梦想！\n", p.name)
}

// 值类型接收者
func (p Person) SetAge(age int) {
	p.age = age
}

// 指针类型接收者
func (p *Person) SetAge1(age int) {
	p.age = age
}

func main() {
	// Go 中的方法是作用于特定类型变量的函数
	var p *Person
	p = NewPerson("张三", "female", 30)
	p.Dream()
	// 值类型接收者未改变 Person 的 age
	p.SetAge(15)
	fmt.Println(p.age)
	// 指针类型接收者改变了 Person 的 age
	p.SetAge1(20)
	fmt.Println(p.age)
}
