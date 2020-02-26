// @Title
// @Description
// @Author       幼儿猿
// @Datetime     2020/2/27 12:37 上午
package main

import "fmt"

// 接口
// 接口是一种类型，接口是一种抽象类型
// 一个对象只要全部实现了接口中的方法，那么就实现了这个接口
type Cat struct {
	name string
}
type Dog struct {
	name string
}

func (c Cat) walk() {
	fmt.Println("我是猫咪，我会走猫步！")
}

func (d Dog) walk() {
	fmt.Println("我是狗狗，我会走狗步！")
}

// walker 接口 一般接口以 er 结尾
type walker interface {
	walk()
}

func main() {
	var (
		c Cat
		d Dog
		w walker
	)
	c.name = "花花"
	w = c // 可以把cat实例直接赋值给w
	w.walk()
	fmt.Println(c)
	fmt.Println(w)
	d.name = "旺财"
	w = d // 可以把dog实例直接赋值给w
	w.walk()
	fmt.Println(d)
	fmt.Println(w)
}
