package main

import (
	"bytes"
	"encoding/base64"
	"log"
	"fmt"
)

func main() {
	// 声明内容
	origin := []byte("Hello World")
	// 声明Buffer
	buffer := bytes.Buffer{}
	// 自定一个64字节的字符串
	customEncode := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
	// 使用给出的字符集生成一个*base64.Encoding，字符集必须是32字节的字符串
	e := base64.NewEncoding(customEncode)
	// 创建一个新的base64流编码器
	w := base64.NewEncoder(e, &buffer)
	// 写入
	if _, err := w.Write(origin); err != nil{
		log.Fatal(err)
	}
	// 关闭
	if err := w.Close(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("base64编码内容: ", string(buffer.Bytes()))
	// 创建一个新的base64流解码器
	originEncode := base64.StdEncoding.EncodeToString(origin)
	fmt.Println("base64编码内容: ", originEncode)
	// 使用标准的base64编码字符集解码
	originBytes, err := base64.StdEncoding.DecodeString(originEncode)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println("base64解码内容: ", string(originBytes))
	// 获取数据进行base64编码后的最大长度
	ne := base64.StdEncoding.EncodedLen(len(origin))
	// 声明[]byte
	dst := make([]byte, ne)
	// 将src的数据编码后存入dst，最多写EncodedLen(len(src))字节数据到dst，并返回写入的字节数
	base64.StdEncoding.Encode(dst, origin)
	fmt.Println("base64编码内容: ", string(dst))
	// 获取base64编码的数据解码后的最大长度
	nd := base64.StdEncoding.DecodedLen(len(dst))
	// 声明[]byte
	originText := make([]byte, nd)
	if _, err := base64.StdEncoding.Decode(originText, dst); err != nil {
		log.Println(err)
	}
	fmt.Println("base64解码内容: ", string(originText))
	// 创建与enc相同的新编码，指定的填充字符除外，或者nopadding禁用填充。填充字符不能是'\r'或'\n'，不能包含在编码的字母表中，并且必须是等于或小于'\xff'的rune
	base64.StdEncoding.WithPadding(base64.StdPadding)
}
