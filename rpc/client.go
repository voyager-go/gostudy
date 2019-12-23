package main

import (
	"net/rpc"
	"log"
	"fmt"
)

func main()  {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil{
		log.Fatal("dialing:", err)
	}
	var reply string
	var argString string
	argString = "I am from client.go!"
	err = client.Call("HelloService.Hello", argString, &reply)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply)
}