package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	str :=   `This is a raw string` + `\n\` + `会被原样输出`
	fmt.Println(str)
	StrPartOne := "Hello"
	StrPartTwo := " Word"
	// 字符串拼接的几种方式
	// https://segmentfault.com/a/1190000012978989
	// 方式一
	fmt.Println(StrPartOne + StrPartTwo)
	// 方式二
	fmt.Println(strings.Join([]string{StrPartOne,StrPartTwo}, ""))
	// 方式三，可以拼接数字
	fmt.Println(fmt.Sprintf("%s%s", StrPartOne, StrPartTwo))
	// 方式四，性能较好
	buf := bytes.Buffer{}
	buf.WriteString(StrPartOne)
	buf.WriteString(StrPartTwo)
	fmt.Println(&buf)
}
// 在循环中使用加号 + 拼接字符串并不是最高效的做法，更好的办法是使用函数 strings.Join()（第 4.7.10 节），有没有更好地办法了？
// 有！使用字节缓冲（bytes.Buffer）拼接更加给力（第 7.2.6 节）