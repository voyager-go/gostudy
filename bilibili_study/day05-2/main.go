// @Title
// @Description
// @Author       幼儿猿
// @Datetime     2020/2/25 11:30 下午
package main

import "fmt"

// 匿名结构体
// 匿名结构体是不常用的 需要注意的是不能出现两个相同类型 违反了结构体不能出现相同属性的原则
type Person1 struct {
	string
	int
}

// 地址
type Address struct {
	Province, City string
}

// 嵌套结构体
// 人类
type Person struct {
	name, gender string
	age          int
	//Addr         Address
	// 与匿名结构体结合
	Address
}

func main() {
	// 匿名结构体
	var person Person1
	person.int = 18
	person.string = "张三"
	fmt.Println(person)

	// 嵌套结构体
	p := Person{
		name:   "幼儿猿",
		gender: "male",
		age:    25,
		Address: Address{
			Province: "安徽",
			City:     "宿州",
		},
	}
	fmt.Printf("%#v\n", p)
	fmt.Println(p.Address.City) // 通过嵌套的匿名结构体字段访问其内部的字段
	fmt.Println(p.Province)     // 直接访问匿名结构体中的字段
}
