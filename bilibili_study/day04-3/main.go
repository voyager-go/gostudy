package main

import (
	"fmt"
	"strconv"
)

// Go语言中strconv包实现了基本数据类型和其字符串表示的相互转换
func main() {
	// Atoi()函数用于将字符串类型的整数转换为int类型
	a, _ := strconv.Atoi("100")
	fmt.Println(a)
	// Itoa()函数用于将int类型数据转换为对应的字符串表示
	// a 的典故【扩展阅读】这是C语言遗留下的典故。C语言中没有string类型而是用字符数组(array)表示字符串，所以Itoa对很多C系的程序员很好理解
	b := strconv.Itoa(100)
	fmt.Printf("type => %T, value => %v\n", b, b)
	// Parse类函数用于转换字符串为给定类型的值：ParseBool()、ParseFloat()、ParseInt()、ParseUint()
	c, _ := strconv.ParseBool("1")
	fmt.Println(c)
	// 返回字符串表示的整数值，接受正负号。
	// base指定进制（2到36），如果base为0，则会从字符串前置判断，”0x”是16进制，”0”是8进制，否则是10进制；
	// bitSize指定结果必须能无溢出赋值的整数类型，0、8、16、32、64 分别代表 int、int8、int16、int32、int64；
	d, _ := strconv.ParseInt("100", 0, 64)
	fmt.Println(d)
	// ParseUint类似ParseInt但不接受正负号，用于无符号整型

	// Format系列函数实现了将给定类型数据格式化为string类型数据的功能
	e := strconv.FormatInt(100, 8)
	// 返回i的base进制的字符串表示。base 必须在2到36之间，结果中会使用小写字母’a’到’z’表示大于10的数字
	fmt.Println(e)
}
