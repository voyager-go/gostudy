// @Title
// @Description
// @Author       幼儿猿
// @Datetime     2020/2/29 10:18 下午
package main

import (
	"fmt"
	"sync"
)

// 互斥锁
//一种常用的控制共享资源的方法，能够保证同时只有一个 goroutine可以访问共享资源
var (
	x    int64
	wg   sync.WaitGroup
	lock sync.Mutex
)

func addNum() {
	for i := 0; i < 5000; i++ {
		lock.Lock() // 加锁
		x++
		lock.Unlock() // 释放锁
	}
	wg.Done()
}
func main() {
	wg.Add(2)
	go addNum()
	go addNum()
	wg.Wait()
	fmt.Println(x)
}
