package main

import (
	"fmt"
	"math"
)

/**
  *整型分为以下两个大类：
  *按长度分为：int8、int16、int32、int64
  *还有对应的无符号整型：uint8、uint16、uint32、uint64
 **/
 /**
  *Go语言支持两种浮点型数：float32 和 float64。这两种浮点型数据格式遵循 IEEE 754 标准：
  *float32 的浮点数的最大范围约为 3.4e38，可以使用常量定义：math.MaxFloat32。
  *float64 的浮点数的最大范围约为 1.8e308，可以使用一个常量定义：math.MaxFloat64。
 **/
 /**
  *字符串中的每一个元素叫做“字符”，在遍历或者单个获取字符串元素时可以获得字符。
  *Go 语言的字符有以下两种：
  *一种是 uint8 类型，或者叫 byte 型，代表了 ASCII 码的一个字符。
  *另一种是 rune 类型，代表一个 UTF-8 字符。当需要处理中文、日文或者其他复合字符时，则需要用到 rune 类型。rune 类型实际是一个 int32。
 **/
// 布尔型数据在 Go 语言中以 bool 类型进行声明，布尔型数据只有 true（真）和 false（假）两个值。
// 将字符串的值以双引号书写的方式是字符串的常见表达方式，被称为字符串字面量（string literal）。这种双引号字面量不能跨行。如果需要在源码中嵌入一个多行字符串时，就必须使用`字符
func main() {
	fmt.Println(math.MaxFloat32)
	fmt.Printf("%f\n", math.Pi)
	fmt.Printf("%.2f\n", math.Pi)
	str := "Hello World" +
		" 你好世界"
	strAnother := `Hello World
你好世界`
	fmt.Println(str)
	fmt.Println(strAnother)
	var varByteA byte = 'a' //ASCII => 97
	var varByteB rune = 'b'
	varByteC := 'c'
	// %T 打印类型
	fmt.Printf("%T", varByteA)
	fmt.Println()
	fmt.Printf("%T", varByteB)
	fmt.Println()
	fmt.Printf("%T", varByteC)
}