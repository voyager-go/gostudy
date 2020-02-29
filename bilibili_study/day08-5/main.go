// @Title
// @Description
// @Author       幼儿猿
// @Datetime     2020/2/29 11:33 下午
package main

import "fmt"

// 两个 goroutine 和两个 channel 协同工作
// f1 生成 0-100 的数字放入 chan1 中
// 通过单向通道来限制函数只能发送或者只能接收的行为
func f1(ch chan<- int) {
	// 这里的 ch 只能接收
	for i := 0; i < 100; i++ {
		ch <- i
	}
	close(ch)
}
func f2(ch1 <-chan int, ch2 chan<- int) {
	// 此处的 ch1 只能发送 ch2 只能接收
	for {
		val, ok := <-ch1
		if !ok {
			break
		}
		ch2 <- val * val
	}
	close(ch2)
}

func main() {
	ch1 := make(chan int, 100)
	ch2 := make(chan int, 200)
	go f1(ch1)
	go f2(ch1, ch2)
	for val := range ch2 {
		fmt.Println(val)
	}
}
