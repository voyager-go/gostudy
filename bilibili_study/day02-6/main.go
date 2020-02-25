package main

import (
	"fmt"
	"strings"
)

func main() {
	// 匿名函数与闭包
	// 1. 匿名函数的使用 变量赋值方式
	var sayHello = func() {
		fmt.Println("say Hello World ....")
	}
	sayHello()

	// 2. 匿名函数的使用 直接调用方式
	func(){
		fmt.Println("say nice to meet you ...")
	}()

	// 闭包 = 函数 + 外部变量
	ret := a("幼儿猿") // ret 此时就是一个闭包
	ret() // 相当于执行了 a 内部的匿名函数
	
	// 深入闭包

	suffixFunc := makeSuffixFunc("苍老师") // 此时 suffixFunc 就是一个闭包
	filename := suffixFunc(".avi")
	fmt.Println(filename)

	// panic 错误
	testPanic()
	keepPanic()
}

func a(name string) func() {
	return func() {
		fmt.Println("Hello ", name)
	}
}
// 为文件名添加后缀
func makeSuffixFunc(name string) func(suffix string) string {
	return func(suffix string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

// 使用 panic 来处理错误
// panic 后程序将无法继续往下执行
// 所以要使用 defer 配合 recover 来处理异常
// 注意 defer recover 要在 panic 之前使用
func testPanic()  {
	fmt.Println("begin")
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("在 panic 之前使用 defer 来延迟执行 recover 处理错误...")
		}
	}()
	panic("test panic error")
}
func keepPanic()  {
	fmt.Println("panic 继续执行 ....")
}