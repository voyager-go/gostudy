package main

import (
	"net"
	"fmt"
	"bufio"
	"io"
	"time"
)

func main() {
	var tcpAddr  *net.TCPAddr
	tcpAddr, _ = net.ResolveTCPAddr("tcp", "127.0.0.1:8929")
	tcpListener, _ := net.ListenTCP("tcp", tcpAddr)
	defer tcpListener.Close()
	fmt.Println("Server ready to read ...")
	for  {
		tcpConn, err := tcpListener.AcceptTCP()
		if err != nil {
			fmt.Println("Accept error : ", err)
			continue
		}
		fmt.Println("A client connectd : " + tcpConn.RemoteAddr().String())
		go tcpPipe(tcpConn)
	}
}
func tcpPipe(conn *net.TCPConn)  {
	ipStr := conn.RemoteAddr().String()
	defer func() {
		fmt.Println(" Disconnected : " + ipStr)
		conn.Close()
	}()

	reader := bufio.NewReader(conn)
	i := 0
	for  {
		message, err := reader.ReadString('\n')
		if err != nil || err == io.EOF {
			break
		}
		fmt.Println(string(message))
		time.Sleep(time.Second * 3)
		msg := time.Now().String() + conn.RemoteAddr().String() + " Server Say hello world! \n"
		b   := []byte(msg)
		conn.Write(b)
		i++
		if i > 10 {
			break
		}
	}
}