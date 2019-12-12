package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"os"
	"log"
	"math/rand"
	"time"
	"strings"
)

type Person struct {
	Name string
	Height string
	Weight string
	Age    string
	Avatar string
}

func main()  {
	str := FetchPage("https://vip.jiayuan.com/cjjllist.php")
	WriteToTxt(str)
}

func ParseHtml()  {
	
}
// 把页面写入到一个文本中
func WriteToTxt(html string)  {
	var (
		fileName string
		file	 *os.File
		err		 error
		n		 int
	)
	fileName = RandString(10) + ".html"
	file, err = os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	n, err = file.WriteString(html)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(n)
}
// 生成随机字符串
func RandString(len int) string {
	var r *rand.Rand
	r = rand.New(rand.NewSource(time.Now().Unix()))
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		b := r.Intn(26) + 65
		bytes[i] = byte(b)
	}
	return strings.ToLower(string(bytes))
}
// 抓取页面
func FetchPage(url string) string {
	var (
		resp *http.Response
		err   error
		html  []byte
	)
	resp, err =  http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	if resp.StatusCode != http.StatusOK {
		fmt.Println(resp.StatusCode)
	}
	
	html, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return string(html)
}