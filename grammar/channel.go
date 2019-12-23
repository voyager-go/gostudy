package main

import (
	"fmt"
	"time"
)

func testChannel(ch chan int)  {
	ch <- 1
	fmt.Println("this is goroutine1")
	time.Sleep(time.Second)
}

func main()  {
	ch := make(chan int, 1)
	go testChannel(ch)
	fmt.Println("this is goroutine2")
	time.Sleep(2 * time.Second)
}