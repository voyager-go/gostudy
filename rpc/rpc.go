package main

import (
	"net/rpc"
	"net"
	"log"
	"fmt"
)

type HelloService struct {}

func (p *HelloService) Hello(request string, reply *string) error {
	*reply = "I am rpc.go!" + request
	fmt.Println(request)
	return nil
}

func main()  {
	rpc.RegisterName("HelloService", new(HelloService))
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	conn, err := listener.Accept()
	if err != nil {
		log.Fatal("Accept error:", err)
	}

	rpc.ServeConn(conn)
}