// @Title
// @Description
// @Author       幼儿猿
// @Datetime     2020/2/27 11:57 下午
package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1) // 计数+1 启动一个 goroutine
	go walk() // 开启了一个 goroutine 去执行这个函数
	fmt.Println("I am next job")
	wg.Wait() // 阻塞， 接着等所有的任务执行完成
}

func walk() {
	fmt.Println("I am walking now ...")
	wg.Done() // 任务完成 通知wg 把计数器-1
}
