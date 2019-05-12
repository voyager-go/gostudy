package main

import (
	"bytes"
	"strconv"
)

func main() {
	Chr(63)
	//fmt.Println(Bin2Hex("acx"))
}
// ********* String Functions ************
// php: addslashes — 使用反斜线引用字符串
func Addslashes ( str string) string  {
	var buff bytes.Buffer
	for _, char := range str {
		switch char {
		case '\'','"','\\':
			buff.WriteRune('\\')
		}
		buff.WriteRune(char)
	}
	return buff.String()
}
// php: bin2hex — 函数把包含数据的二进制字符串转换为十六进制值
func Bin2Hex(str string) (string, error) {
	i, err := strconv.ParseInt(str, 2, 0)
	if err != nil {
		return "", err
	}
	return strconv.FormatInt(i, 16), nil
}
// php: chr — 返回相对应于 ascii 所指定的单个字符
func Chr(ascii int) string {
	return string(ascii)
}