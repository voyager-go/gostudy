package main

import "fmt"

// 使用键值对填充结构体的例子
type People struct {
	name string
	child *People
}

// 多个值列表初始化结构体的例子
type Address struct {
	Province string
	City string
	ZipCode int
	PhoneNumber string
}

// 使用匿名结构体的例子
// 打印消息类型 传入匿名结构体
func printMsgType(msg *struct{
	id int
	data string
})  {
	// 使用动词%T打印msg的类型
	fmt.Printf("%T\n", msg)
}
func main() {
	relation := &People{
		name: "爷爷",
		child: &People{
			name: "爸爸",
			child: &People{
				name: "我",
			},
		},
	}

	fmt.Println(relation.child.child.name)


	addr := Address{
		"四川",
		"成都",
		61000,
		"0",
	}
	fmt.Println(addr)


	// 实例化一个匿名结构体
	msg := &struct {
		id int
		data string
	}{
		1024,
		"hello",
	}
	printMsgType(msg)
}