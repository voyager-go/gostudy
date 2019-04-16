package main

import (
	"fmt"
	"errors"
)
// Go 语言的错误处理思想及设计包含以下特征：

// 一个可能造成错误的函数，需要返回值中返回一个错误接口（error）。如果调用是成功的，错误接口将返回 nil，否则返回错误。
// 在函数调用后需要检查错误，如果发生错误，进行必要的错误处理。

func main() {
	fmt.Println(div(1, 0))
}
// 定义除数为0的错误
var errDivisionByZero = errors.New("division by zero")

func div(dividend, divisor int) (int, error) {
	if divisor == 0 {
		return 0, errDivisionByZero
	}
	// 正常计算返回空错误
	return dividend / divisor, nil
}