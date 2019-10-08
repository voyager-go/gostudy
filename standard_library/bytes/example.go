package main

import (
	"bytes"
	"fmt"
	"os"
	"encoding/base64"
	"io"
)

// bytes包实现了操作[]byte的常用函数。本包的函数和strings包的函数相当类似
func main() {
	x := []byte("abc")
	y := []byte("a")
	if bytes.Compare(x, y) > 0 {
		fmt.Println("x > y")
	}else {
		fmt.Println(bytes.Compare(x, y))
	}
	write()
	newBuffer()
	exampleBuffer()
	exampleBufferReader()
	var a, b []byte
	// 根据[]byte创建 reader
	bytes.NewReader([]byte("shan gui"))
	// 比较a和b, 返回 0: a等于b, 1: a包含b, -1: a不包含b
	compareNum := bytes.Compare(a, b)
	fmt.Printf("a与b的比较结果为:%d\n", compareNum)
	// 判断a与b是否相同
	equalBool := bytes.Equal(a, b)
	fmt.Printf("a与b的比较结果为:%v\n", equalBool)
	// 判断a与b是否相同，忽略大小写
	equalFoldBool := bytes.EqualFold(a, b)
	fmt.Printf("a与b的比较结果为:%v\n", equalFoldBool)
}

func write()  {
	buf := bytes.Buffer{}
	// 增加buffer容量
	buf.Grow(64)
	// 写入字符串
	buf.WriteString("Hello ")
	// 写入单个byte
	buf.WriteByte('W')
	// 写入单个rune
	buf.WriteRune('o')
	// 写入byte
	buf.Write([]byte("rld"))
	fmt.Printf("%s\n", buf.Bytes())

}

func newBuffer()  {
	str  := "Hello golang"
	buf  := bytes.NewBufferString(str)
	buff := bytes.NewBuffer([]byte(str))
	fmt.Printf("%s\n", buf.Bytes())
	fmt.Printf("%s\n", buff.Bytes())
}

func exampleBuffer()  {
	var buf bytes.Buffer
	buf.Write([]byte("Example  "))
	fmt.Fprintf(&buf, "Buffer\n")
	buf.WriteTo(os.Stdout)
}

func exampleBufferReader()  {
	buf := bytes.NewBufferString("R29waGVycyBydWxlIQ==")
	dec := base64.NewDecoder(base64.StdEncoding, buf)
	io.Copy(os.Stdout, dec)
}
