package main

import (
	"bytes"
	"strconv"
	"fmt"
)
// 使用Golang实现PHP的部分函数
func main() {
	//Chr(63)
	//fmt.Println(Bin2Hex("acx"))

	//var arrayMap = make(map[string]interface{})
	//arrayMap["red"] = "apple"
	//arrayMap["green"] = "tree"
	//fmt.Println(ArrayKeys(arrayMap))

	//var newArray = Array("apple", "pear", "banana")
	//fmt.Println(len(newArray))

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
// php: array_keys - 返回数组中部分的或所有的键名
func ArrayKeys(data map[string]interface{}) []string{
	if len(data) < 1 {
		return []string{}
	}
	var resData []string
	for index  := range data {
		resData = append(resData, index)
	}
	return  resData
}
// php: array — 新建一个数组
func Array(val ...interface{})[]interface{}  {
	return val
}
// php: count — 计算数组中的单元数目，或对象中的属性个数
func Count(data []interface{})int  {
	return len(data)
}
// php: array_column — 返回数组中指定的一列
//func ArrayColumn(arrayMap map[string]map[string]interface{}, columnKey string) (r []interface{}) {
//	fmt.Println(arrayMap)
//	return
//}