package main

import (
	"net"
	"bufio"
	"fmt"
	"io"
	"time"
)

func main()  {
	var tcpAddr *net.TCPAddr
	tcpAddr, _ = net.ResolveTCPAddr("tcp", "127.0.0.1:8929")
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		fmt.Println("Client connect error !" + err.Error())
		return
	}
	defer conn.Close()
	fmt.Println(conn.LocalAddr().String() + " : Client connected!")
	onMessageReceived(conn)
}

func onMessageReceived(conn *net.TCPConn)  {
	reader := bufio.NewReader(conn)
	b := []byte(conn.LocalAddr().String() + " Say hello to Server...\n")
	conn.Write(b)
	for {
		msg, err := reader.ReadString('\n')
		fmt.Println("ReadString")
		fmt.Println(msg)
		if err != nil || err == io.EOF {
			fmt.Println(err)
			break
		}
		time.Sleep(time.Second * 2)
		fmt.Println("Writing...")
		b := []byte(conn.LocalAddr().String() + " write data to Server...\n")
		_, err = conn.Write(b)
		if err != nil {
			fmt.Println(err)
			break
		}
	}
}
