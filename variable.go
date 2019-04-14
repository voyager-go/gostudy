package main

import "fmt"
// 常规变量声明与赋值
var intVar int = 100
var strVar string = "hello world"
var spliceVar []float32
var callbackVal func() bool
var structVal struct{
	xx int
}
// 多变量声明
var(
	intVarA int = 20
	strVarB string
)

// 让编译器去推导类型
// 等号右边的部分在编译原理里被称做右值（rvalue）
var intA  = 200;
// 匿名变量
// 在使用多重赋值时，如果不需要在左值中接收变量，可以使用匿名变量（anonymous variable）
func GetData() (int, int)  {
	return 100, 200
}
func main()  {
	a,_ := GetData()
	_,b := GetData()
	c,d := GetData()
	fmt.Println(intA)
	fmt.Println(intVar)
	fmt.Println(intVarA)
	fmt.Println(a, b, c, d)
	// 短变量赋值 在函数中使用
	shortVar := 1000
	fmt.Println(shortVar)
}