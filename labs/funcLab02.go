package main

import (
	"fmt"
	"flag"
)
// 匿名函数
// 定义命令行参数 skill，从命令行输入— skill 可以将空格后的字符串传入 skillParam 指针变量
var skillParam = flag.String("skill", "", "skill to perform")
func main() {
	fun := func(data string) {
		fmt.Println("hello", data)
	}
	fun("world")
	// 匿名函数用作回调函数
	// 使用匿名函数打印切片内容
	visit([]int{1, 2, 3, 4, 5}, func(v int) {
		fmt.Println(v)
	})


	// 解析命令行参数，解析完成后，skillParam 指针变量将指向命令行传入的值
	flag.Parse()
	// 定义一个从字符串映射到 func()的map，然后填充这个 map
	var skill = map[string]func(){
		"fire": func() {
			fmt.Println("chicken fire")
		},
		"run": func() {
			fmt.Println("soldier run")
		},
		"fly": func() {
			fmt.Println("angel fly")
		},
	}
	// skillParam 是一个 *string 类型的指针变量，使用 *skillParam 获取到命令行传过来的值，并在 map 中查找对应命令行参数指定的字符串的函数
	if fun, ok := skill[*skillParam]; ok {
		fun()
	}else {
		fmt.Println("skill not found")
	}
}
// 遍历切片的每个元素, 通过给定函数进行元素访问
func visit(list []int, f func(int))  {
	for _,v := range list {
		f(v)
	}
}
