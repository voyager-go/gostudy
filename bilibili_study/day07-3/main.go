// @Title
// @Description
// @Author       幼儿猿
// @Datetime     2020/2/27 11:03 下午
package main

import (
	"fmt"
	"reflect"
)

// 反射
func main() {
	var a float64 = 3.33
	v := reflect.ValueOf(a)
	fmt.Println(v, v.Kind())
	var b int32 = 22
	reflectValue(b)
	var c string = "Hello"
	setVal(&c)
	fmt.Printf("%v, %T\n", c, c)
}

// 如果根据变量反射获取原本类型的值
func reflectValue(x interface{}) {
	v := reflect.ValueOf(x)
	switch v.Kind() {
	case reflect.Int32:
		x = int32(v.Int())
		fmt.Printf("%v, %T\n", x, x)
		break
	case reflect.Float64:
		x = float64(v.Float())
		fmt.Printf("%v, %T\n", x, x)
		break
	}
}

// 通过反射设置变量的值
func setVal(x interface{}) {
	v := reflect.ValueOf(x)
	switch v.Elem().Kind() {
	case reflect.Float64:
		v.Elem().SetFloat(1.234)
		break
	case reflect.String:
		v.Elem().SetString("World")
		break
	}
}
