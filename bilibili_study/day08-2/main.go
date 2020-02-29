// @Title
// @Description
// @Author       幼儿猿
// @Datetime     2020/2/29 10:29 下午
package main

import (
	"fmt"
	"sync"
	"time"
)

// 读写锁
// 比较读写锁与互斥锁的效率
// 读写锁非常适合读多写少的场景，如果读和写的操作差别不大，读写锁的优势就发挥不出来
var (
	x  int64
	wg sync.WaitGroup
	//lock sync.Mutex
	relock sync.RWMutex
)

func read() {
	//lock.Lock()
	relock.RLock()
	time.Sleep(time.Millisecond) // 读一毫秒
	relock.RUnlock()
	//lock.Unlock()
	wg.Done()
}

func write() {
	//lock.Lock()
	relock.Lock()
	x += 1
	time.Sleep(time.Millisecond * 10) // 写 10 毫秒
	//lock.Unlock()
	relock.Unlock()
	wg.Done()
}

func main() {
	start := time.Now()
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go read()
	}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go write()
	}
	wg.Wait()
	fmt.Println(time.Now().Sub(start))
}
