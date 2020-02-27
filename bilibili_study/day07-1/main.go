// @Title
// @Description
// @Author       幼儿猿
// @Datetime     2020/2/27 5:02 下午
package main

import "fmt"

// 一个接口的值是由一个具体类型和具体类型的值两部分组成的，这两部分分别称为接口的动态类型和动态值

// 类型断言
func main() {
	var x interface{}
	x = "Hello World"
	if v, ok := x.(string); ok {
		fmt.Printf("%T, %s\n", v, v)
	} else {
		fmt.Println("类型断言失败！")
	}
	x = 100
	switch x.(type) {
	case string:
		fmt.Printf("是 string 类型. %T\n", x)
		break
	case int:
		fmt.Printf("是 int 类型%T\n", x)
		break
	case bool:
		fmt.Printf("是 bool 类型%T\n", x)
		break
	default:
		fmt.Println("不知道什么类型")
	}
}
