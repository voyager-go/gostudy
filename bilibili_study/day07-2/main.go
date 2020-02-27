// @Title
// @Description
// @Author       幼儿猿
// @Datetime     2020/2/27 10:48 下午
package main

import (
	"fmt"
	"reflect"
)

// 反射
func reflectType(x interface{}) {
	t := reflect.TypeOf(x)
	fmt.Println(x, t, t.Name(), t.Kind())
}
func main() {
	var a int = 100
	reflectType(a)
	var b string = "Hello"
	reflectType(b)
	var c int8 = 109
	reflectType(c)
	var d float64 = 1.22
	reflectType(d)
	// 复合类型
	var e struct{}
	reflectType(e)
	var cat Cat
	reflectType(cat)
	var f []int
	reflectType(f)
	var g map[string]int
	reflectType(g)
}

type Cat struct{}
