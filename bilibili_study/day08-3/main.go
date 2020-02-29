// @Title
// @Description
// @Author       幼儿猿
// @Datetime     2020/2/29 10:59 下午
package main

import (
	"fmt"
	"strconv"
	"sync"
)

// Go 语言中内置的 map 不是并发安全的
// 开启少量几个goroutine的时候可能没什么问题，当并发多了之后执行上面的代码就会报fatal error: concurrent map writes错误
// Go语言的sync包中提供了一个开箱即用的并发安全版map–sync.Map。开箱即用表示不用像内置的map一样使用make函数初始化就能直接使用。同时sync.Map内置了诸如Store、Load、LoadOrStore、Delete、Range等操作方法
var m map[string]int = map[string]int{}
var m2 sync.Map
var wg sync.WaitGroup

func get(key string) int {
	return m[key]
}
func set(key string, val int) {
	m[key] = val
}
func main() {
	for i := 0; i < 30; i++ {
		wg.Add(1)
		go func(i int) {
			key := strconv.Itoa(i)
			m2.Store(key, i+100)
			//set(key, i+100)
			val, _ := m2.Load(key)
			fmt.Printf("k=:%v,v:=%v\n", key, val)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
