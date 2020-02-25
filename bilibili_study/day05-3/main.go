// @Title
// @Description
// @Author       幼儿猿
// @Datetime     2020/2/25 11:56 下午
package main

import "fmt"

// 结构体的继承
type Animal struct {
	Name string
}
type Dog struct {
	Color string
	*Animal
}

// Animal 的行走方法
func (a *Animal) walk() {
	fmt.Printf("%s可以行走...\n", a.Name)
}
func (d *Dog) Bark() {
	fmt.Printf("%s可以犬吠...\n", d.Name)
}
func main() {
	dog := Dog{
		Color: "Black",
		Animal: &Animal{
			Name: "旺财",
		},
	}
	dog.Bark()
	dog.walk()
}
