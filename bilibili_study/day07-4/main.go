// @Title
// @Description
// @Author       幼儿猿
// @Datetime     2020/2/27 11:26 下午
package main

import (
	"fmt"
	"reflect"
)

type student struct {
	Name  string `json:"name" ini:"s_name"`
	Score int    `json:"score" ini:"s_score"`
}

// 结构体反射
func main() {
	// 通过反射获取结构体中的所有字段信息
	stu := student{
		Name:  "张大仙",
		Score: 80,
	}
	t := reflect.TypeOf(stu)
	fmt.Printf("value => %v, type => %T\n", t, t)
	for i := 0; i < t.NumField(); i++ {
		// 根据结构体的索引去取字段
		field := t.Field(i)
		fmt.Printf("name: %s, tag: %s, type: %v, index: %d\n", field.Name, field.Tag, field.Type, field.Index)
		fmt.Println(t.Field(i).Tag.Get("json"), t.Field(i).Tag.Get("ini"))
	}
	// 根据名字取结构体字段
	fieldObj, ok := t.FieldByName("Score")
	if ok {
		fmt.Println(fieldObj.Tag)
	}

	printMethods(stu)
}

// 根据反射去获取结构体中的方法
func printMethods(x interface{}) {
	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)
	for i := 0; i < t.NumMethod(); i++ {
		method := t.Method(i)
		fmt.Println(method.Name, method.Type, method.Func)

		// 通过反射调用方法传参必须是 []reflect.Value 类型
		var arg []reflect.Value
		v.Method(i).Call(arg) // 调用方法
	}
}

func (s student) Sleep() {
	fmt.Println("I am sleeping ...")
}
func (s student) Eat() {
	fmt.Println("I am eating ...")
}
