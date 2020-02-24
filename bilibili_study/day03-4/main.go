package main

import "fmt"

type Person struct {
	Name, Gender string
	Age          int
	Married      bool
}

func main() {
	// 匿名结构体
	var user struct {
		Name    string
		Job     string
		Married bool
	}
	user.Name = "幼儿猿"
	user.Job = "coder"
	user.Married = false
	fmt.Printf("%#v\n", user)

	// 结构体指针
	var p = new(Person)
	(*p).Name = "张文杰"
	(*p).Married = false
	(*p).Age = 25
	(*p).Gender = "male"
	fmt.Printf("%#v\n", *p)
	// Go 语言为结构体指针提供了一个简写的语法糖
	p.Name = "张二狗"
	p.Gender = "female"
	p.Age = 23
	fmt.Printf("%#v\n", *p)

	// 取结构体的地址进行实例化
	var per = &Person{}
	fmt.Printf("type 为 %T\n", per)
	fmt.Printf("%#v\n", *per)
	per.Name = "王二"
	per.Married = true
	per.Gender = "female"
	fmt.Printf("%#v\n", *per)

	// 结构体初始化的两种方式
	// 1. 键值对初始化
	var p1 = Person{
		Name:    "李四",
		Gender:  "male",
		Age:     30,
		Married: false,
	}
	fmt.Printf("%#v\n", p1)

	// 2. 值列表进行初始化 此处需要注意的是 列表的值必须与结构体的 key 一一对应
	var p2 = Person{
		"王刚",
		"male",
		24,
		true,
	}
	fmt.Printf("%#v\n", p2)
	// 结构体字段占用一块连续的内存

	// 结构体是值类型 尤其是较为复杂的结构体值拷贝的时候内存开销较大 所以使用时尽量使用结构体指针
	p3 := newPerson("狗蛋儿", "female", 12, false)
	fmt.Printf("%#v\n", p3)
}

// 结构体的构造函数 通常以 new 开始
func newPerson(name, gender string, age int, married bool) *Person {
	return &Person{
		Name:    name,
		Gender:  gender,
		Age:     age,
		Married: married,
	}
}
