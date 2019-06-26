package main

import (
	"fmt"
	"time"
)

func main(){
	ch := make(chan string)
	go sendData(ch)
	go getData(ch)
	time.Sleep(1e9)
}
func sendData(ch chan string){
	ch <- "zhangwenjie"
	ch <- "liusiji"
	ch <- "duandapao"
	ch <- "shumengnan"
}

func getData(ch chan string){
	var input string
	for {
		input = <- ch
		fmt.Printf("%s\n", input)
	}
}