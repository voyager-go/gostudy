package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

// Go语言标准库flag基本使用
func main() {
	if len(os.Args) > 0 {
		for index, arg := range os.Args {
			fmt.Printf("args[%d]=%v\n", index, arg)
		}
	}
	// 将上面的代码执行go build -o "args_demo"编译之后，执行: $ ./args_demo a b c d
	// os.Args是一个存储命令行参数的字符串切片，它的第一个元素是执行文件的名称

	// 定义命令行参数的方式
	var (
		name    string
		age     int
		married bool
		delay   time.Duration
	)
	flag.StringVar(&name, "name", "张三", "姓名")
	flag.IntVar(&age, "age", 20, "年龄")
	flag.BoolVar(&married, "married", false, "是否已婚")
	flag.DurationVar(&delay, "d", 0, "延迟的时间间隔")
	// 解析命令行参数
	// 定义好命令行flag参数后，需要通过调用flag.Parse()来对命令行参数进行解析
	flag.Parse()
	fmt.Println(name, age, married, delay)
	// 返回命令行后的其它参数
	fmt.Println(flag.Args())
	// 返回命令行后的其它参数个数
	fmt.Println(flag.NArg())
	// 返回使用的命令行参数个数
	fmt.Println(flag.NFlag())
}
