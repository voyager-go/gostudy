package main

import (
	"fmt"
	"strings"
)

func main() {
	// 字符串长度
	s1 := "Hello World"
	fmt.Println("字符串长度：", len(s1))
	s2 := ", 你好 世界"
	fmt.Println(fmt.Sprintf("字符串拼接：%s%s", s1, s2))
	s3 := s1 + s2
	fmt.Println("字符串拼接：", s3)
	// 字符串数组拼接
	a := []string{"a", "b", "c"}
	fmt.Println("字符串数组拼接：", strings.Join(a, "+"))
	s4 := "1\\2\\3\\4\\5"
	fmt.Println(s4)
	fmt.Println("字符串分隔:", strings.Split(s4, "\\"))
	fmt.Println("字符串包含:", strings.Contains(s3, "llo"))
	fmt.Println("字符串前缀:", strings.HasPrefix(s4, "1"))
	fmt.Println("字符串后缀:", strings.HasSuffix(s4, "5"))
	s5 := "abscdabc"
	fmt.Println("字符b 在 abscdabc 中首先出现的位置：", strings.Index(s5, "b"))
	fmt.Println("字符b 在 abscdabc 中最后出现的位置：", strings.LastIndex(s5, "b"))

}
