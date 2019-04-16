package main

import (
	"fmt"
	"net/http"
	"strings"
	"os"
	"io/ioutil"
)

// Go 语言可以对任何类型添加方法。给一种类型添加方法就像给结构体添加方法一样，因为结构体也是一种类型。
// 为基本类型添加方法
// 将int定义为MyInt类型
type MyInt int

// 为MyInt添加IsZero()方法
func (m MyInt) IsZero() bool {
	return m == 0
}

// 为MyInt添加Add()方法
func (m MyInt) Add(other int) int{
	return other + int(m)
}

func main() {
	var b MyInt
	fmt.Println(b.IsZero())
	b = 1
	fmt.Println(b.IsZero())
	fmt.Println(b.Add(2))

	// http包中的类型方法
	client := &http.Client{}
	// 创建一个http请求
	req, err := http.NewRequest("GET", "http://www.baidu.com/", strings.NewReader("key=value"))
	// 发现错误就立即打印并退出
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		return
	}
	// 为标头添加信息
	req.Header.Add("User-Agent", "myClient")
	// 开始请求
	resp, err := client.Do(req)

	// 处理请求的错误
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		return
	}

	data, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(data))
	defer resp.Body.Close()
}
