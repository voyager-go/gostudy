package main

import "fmt"
// go 中的字符有两种 byte 和 rune
func main() {
	// byte uint8 的别名 对应 ASCII码
	// rune int32 的别名
	var c1 byte = 'c'
	var c2 rune = 'c'
	fmt.Println(c1, c2)
	fmt.Printf("c1:%T, c2:%T\n", c1, c2)
	// %c 是把一个数字按照字符给打印出来
	s := "Hello 文杰"
	for i := 0 ; i < len(s); i++ {
		fmt.Printf("%c\n", s[i])
	}

	for _, str := range s {
		fmt.Printf("%c\n", str)
	}
}
