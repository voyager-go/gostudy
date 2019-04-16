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

	var e error
	// 创建一个错误实例 包含文件名和行号
	e = newParseError("funcLab06.go", 1)
	// 通过error接口查看错误描述
	fmt.Println(e.Error())
	// 根据接口具体的类型，获取详细错误信息
	switch detail := e.(type) {
	case *ParseError:
		fmt.Printf("Filename: %s Line: %d\n", detail.Filename, detail.Line)
	default:
		fmt.Println("other error")
	}
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

// 在解析中使用自定义错误

// 声明一个解析错误
type ParseError struct {
	Filename string
	Line     int
}

// 实现error接口，返回错误描述
func (e *ParseError) Error() string  {
	return fmt.Sprintf("%s:%d", e.Filename, e.Line)
}

// 创建一些解析错误
func newParseError(filename string, line int) error {
	return &ParseError{filename, line}
}