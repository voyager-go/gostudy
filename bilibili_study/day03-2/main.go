package main

import "fmt"

func main() {
	a := 10
	b := &a	// *int 类型(int 指针)
	fmt.Printf("%v, %p\n", a, &a)
	fmt.Printf("%v, %p\n", b, &b)
	// 变量 b 是一个 int 类型的指针(*int)，它存储的是变量 a 的内存地址
	c := *b // 指针取值，根据指针取内存地址
	fmt.Printf("值为 => %v\n", c)
	fmt.Printf("类型 => %T\n", &c)
	func1(a)
	fmt.Println(a)
	func2(&a)
	fmt.Println(a)
	var d *int 	// nil
	d = new(int) // 初始化
	fmt.Println(*d)
	*d = 100
	fmt.Println(*d)
	// new 很少使用
	// new 与 make 的区别
	// 1. 二者都是用来做内存分配的
	// 2. make 只用于 slice map channel的初始化，返回的还是这三个引用类型的本身
	// 3. new  只用于类型的内存分配，并且对应的值为类型零值，返回的是指向类型的指针
	e := new(bool)
	fmt.Println(*e) // false
}

func func1(x int)  {
	x = 100
}

func func2(y *int)  {
	*y = 100
}