package main

import "fmt"

func main() {
	fmt.Println("defer begin")

	// 将defer放入延迟调用栈
	defer fmt.Println(1)
	defer fmt.Println(2)
	// 最后一个放入的，位于栈顶，最新调用
	defer fmt.Println(3)

	fmt.Println("defer end")
}

// defer 语句正好是在函数退出时执行的语句，所以使用 defer 能非常方便地处理资源释放问题。
