package main

import (
	"crypto/md5"
	"fmt"
	"io"
)

func main() {
	// 方案一
	str := "123456"
	data := []byte(str)
	hashStr := md5.Sum(data)
	md5Str := fmt.Sprintf("%x", hashStr)
	fmt.Println(md5Str)
	// e10adc3949ba59abbe56e057f20f883e

	// 方案二
	w := md5.New()
	io.WriteString(w, str)
	newMd5Str := fmt.Sprintf("%x",  w.Sum(nil))
	fmt.Println(newMd5Str)
	// e10adc3949ba59abbe56e057f20f883e
}
