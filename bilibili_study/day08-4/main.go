// @Title
// @Description
// @Author       幼儿猿
// @Datetime     2020/2/29 11:17 下午
package main

import "fmt"

func main() {
	var ch1 chan int        // 通道是一个引用类型  需要初始化之后再进行使用
	ch1 = make(chan int, 1) // 有缓冲区通道
	//ch1 = make(chan int) // 无缓冲区通道 是阻塞的通道
	ch1 <- 10
	x := <-ch1
	fmt.Println(x)
	close(ch1)
}
