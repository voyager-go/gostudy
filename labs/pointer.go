package main

import "fmt"

var intVarD = 8

func swap(a, b *int)  {
	*a, *b = *b, *a
}
func main()  {
	ptr := &intVarD
	fmt.Println(ptr)
	fmt.Println(*ptr)
	// 准备一个字符串类型
	var house = "Malibu Point 10880, 90265"
	// 对字符串取地址, ptr1类型为*string
	ptr1 := &house
	fmt.Printf("%p\n", ptr1)
	fmt.Println(*ptr1)
	intValE, intVarF := 22, 33
	swap(&intValE, &intVarF)
	fmt.Println(intValE, intVarF)
	// 创建指针的另一种方法——new() 函数
	strA := new(string)
	*strA = "nihao"
	println(*strA)
}