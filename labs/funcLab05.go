package main

import (
	"bytes"
	"fmt"
)

// 所谓可变参数，是指参数数量不固定的函数形式。Go 语言支持可变参数特性，函数声明和调用时没有固定数量的参数，同时也提供了一套方法进行可变参数的多级传递。
// 遍历可变参数列表——获取每一个参数的值

// 定义一个函数, 参数数量为0~n, 类型约束为字符串
func joinStrings(slist ...string) string {
	// 定义一个字节缓冲, 快速地连接字符串
	var b bytes.Buffer
	// 遍历可变参数列表slist, 类型为[]string
	for _, s := range slist {
		// 将遍历出的字符串连续写入字节数组
		b.WriteString(s)
	}
	// 将连接好的字节数组转换为字符串并输出
	return b.String()
}

// 获得可变参数类型——获得每一个参数的类型
func printTypeValue(slist ...interface{}) string {
	var b bytes.Buffer
	for _, s := range slist{
		str := fmt.Sprintf("%v", s)
		var typeString string
		// 对s进行类型断言
		switch s.(type) {
		case bool:
			typeString = "bool"
		case string:
			typeString = "string"
		case int:
			typeString = "int"
		}
		// 写值字符串前缀
		b.WriteString("value: ")
		// 写入值
		b.WriteString(str)
		// 写类型前缀
		b.WriteString(" type: ")
		// 写类型字符串
		b.WriteString(typeString)
		// 写入换行符
		b.WriteString("\n")
	}
	return b.String()
}


// 在多个可变参数函数中传递参数
// 实际打印的函数
func rawPrint(rawList ...interface{}) {
	// 遍历可变参数切片
	for _, a := range rawList {
		// 打印参数
		fmt.Println(a)
	}
}
// 打印函数封装
func print(slist ...interface{}) {
	// 将slist可变参数切片完整传递给下一个函数
	rawPrint(slist...)
}

func main() {
	// 输入3个字符串, 将它们连成一个字符串
	fmt.Println(joinStrings("pig", "and", "rat"))

	// 将不同类型的变量通过printTypeValue()打印出来
	fmt.Println(printTypeValue(100, "str", true))

	//ss := fmt.Sprintf("%v", 100)
	//fmt.Println(ss)

	print("hello", 10000, false)
}
