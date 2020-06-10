package main

import (
	"crypto/md5"
	"fmt"
	"io"
)

func main() {
	var passwd = []byte("123456")
	hash := md5.Sum(passwd)
	fmt.Printf("%x\n", hash)

	h := md5.New()
	io.WriteString(h, "Hello")
	io.WriteString(h, " World")
	fmt.Printf("%x\n", h.Sum(nil))
}
