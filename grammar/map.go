package main

import "fmt"

func main()  {
	//var m map[string]string
	//m := make(map[string]string)
	//m["name"] = "youeryuan"
	//m["gender"] = "man"
	//m["age"] = "25"
	//fmt.Println(m)
	//// curd
	//m["job"] = "coder"
	//fmt.Println(m)
	//// 删除
	//delete(m, "gender")
	//fmt.Println(m)
	//// 遍历删除
	//for x, _ := range m {
	//	delete(m, x)
	//	fmt.Println(m)
	//}
	// map 切片
	//var slice []map[string]string
	slice := make([]map[string]string, 1)
	m1 := make(map[string]string)
	m2 := make(map[string]string)
	m1["name"] = "youeryuan"
	m1["job"] = "coder"
	m2["name"] = "yanfang"
	m2["job"] = "driver"
	slice = append(slice, m1, m2)
	fmt.Println(slice)
}
