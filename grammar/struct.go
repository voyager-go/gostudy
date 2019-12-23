// 结构体的操作练习
// struct和其他语言的class具有同等地位
// golang没有extends继承是通过匿名字段
// 面向对象编程 => 面向接口编程
package main

import "fmt"

type User struct {
	Name   string
	Age    int
	Gender string
	Score  [3]int
}
type Cat struct {
	Name  string
	Color string
	Age   int
	Behavior string
}
type Struct struct {
	Per *int
	Name string
	Age  int
	Score [] float64
	Hobby [3] float64
	Friend map[string]string
}

func main()  {
	//user := User{
	//	Name: "youeryuan",
	//	Age: 25,
	//	Gender: "male",
	//}
	//fmt.Println(user)
	//user.Age = 30
	//fmt.Println(user.Age)
	//user.Score = [3]int{99, 80, 60}
	//fmt.Println(user)
	//cat := Cat{
	//	Name: "小白",
	//	Age: 5,
	//	Color: "Black",
	//}
	//fmt.Println(cat)
	//var p1 Struct
	//fmt.Println(p1)
	//if p1.Per == nil {
	//	fmt.Println(true)
	//}
	// struct 中包含 slice
	//p1.Score = append(p1.Score, 120.2)
	//fmt.Println(p1)
	//p1.Score = make([] float64, 10)
	//p1.Score[0] = 10
	//fmt.Println(p1)
	// struct中包含map
	//p1.Friend = make(map[string]string)
	//p1.Friend["Name"] = "NC"
	//fmt.Println(p1)
	//结构体为值类型，默认为值拷贝
	//创造结构体变量的四种方式
	//var U *User = new(User)
	//(*U).Name  = "youeryuan"
	//(*U).Score = [...]int{100, 20, 30}
	//fmt.Println(*U)
	var person *User = &User{"yanfang", 23, "female", [3]int{1, 2, 3}}
	fmt.Println(person)
}