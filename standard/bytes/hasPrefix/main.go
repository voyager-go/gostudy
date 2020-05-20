package main

import (
	"bytes"
	"fmt"
)

func main() {
	//var a = []byte{12}
	//var a = []byte{13}
	var a = []byte{14}
	var b = []byte{12, 13, 14}
	var c = []byte{15}
	// 判断s是否有前缀切片prefix
	if bytes.HasPrefix(b, a) {
		fmt.Println("b has prefix a")
	} else {
		fmt.Println("b has not prefix a")
	}

	if bytes.HasSuffix(b, a) {
		fmt.Println("b has suffix a")
	} else {
		fmt.Println("b has not suffix a")
	}

	if bytes.Contains(b, c) {
		fmt.Println("b has contains c")
	} else {
		fmt.Println("b has not contains c")
	}
}
