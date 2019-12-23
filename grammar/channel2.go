package main

import (
	"fmt"
)
/**
关闭一个未初始化(nil) 的 channel 会产生 panic
重复关闭同一个 channel 会产生 panic
向一个已关闭的 channel 中发送消息会产生 panic
从已关闭的 channel 读取消息不会产生 panic，且能读出 channel 中还未被读取的消息，若消息均已读出，则会读到类型的零值。从一个已关闭的 channel 中读取消息永远不会阻塞，并且会返回一个为 false 的 ok-idiom，可以用它来判断 channel 是否关闭
关闭 channel 会产生一个广播机制，所有向 channel 读取消息的 goroutine 都会收到消息
**/
//func main()  {
	//ch := make(chan int, 10)
	//ch <- 10
	//ch <- 12
	//close(ch)
	//for x := range ch {
	//	fmt.Println(x)
	//}
	//x, ok := <-ch
	//fmt.Println(x, ok)
	
	
//}

func fibonacii(c, quit chan int)  {
	x, y := 0, 1
	for  {
		select {
		case c <- x:
			x, y = y, x+y
		case <- quit:
			fmt.Println("quit")
			return 
			
		}
	}
}

func main()  {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacii(c, quit)
}
