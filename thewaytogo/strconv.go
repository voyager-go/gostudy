package main

import (
	"strings"
	"fmt"
)

// HasPrefix 判断字符串 s 是否以 prefix 开头：
// HasSuffix 判断字符串 s 是否以 prefix 开头：
// Contains 判断字符串 s 是否包含 substr
// Index 返回字符串 str 在字符串 s 中的索引（str 的第一个字符的索引），-1 表示字符串 s 不包含字符串 str
// LastIndex 返回字符串 str 在字符串 s 中最后出现位置的索引（str 的第一个字符的索引），-1 表示字符串 s 不包含字符串 str：
// IndexRune 如果需要查询非 ASCII 编码的字符在父字符串中的位置，建议使用以下函数来对字符进行定位：
func main()  {
	str := "strings"
	//println(strings.HasPrefix(str, "s"))
	//println(strings.HasSuffix(str, "s"))
	//fmt.Println(str)
	//fmt.Println(strings.Contains(str, "s"))
	//fmt.Println(strings.Index(str, "t"))
	//fmt.Println(strings.LastIndex(str, "s"))

	//fmt.Println(strings.IndexRune(str, 5))
	oldStr := "s"
	newStr := "l"
	fmt.Println(strings.Replace(str, oldStr, newStr, 2))
	strOne := "stringSs"
	fmt.Println(strings.Count(strOne, "s"))
	strTwo := "SSSTRINGS"
	fmt.Println(strings.ToLower(strTwo))
	strThree := "sssss"
	fmt.Println(strings.ToUpper(strThree))
	fmt.Println(strings.NewReader(strOne))
}
