package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
)

func main() {
	// map 映射
	// 只声明 不初始化 初始值就是 nil
	var a map[int]string
	fmt.Println(a == nil)

	// map 的初始化
	var b = make(map[int]string, 8)
	fmt.Println(b == nil)
	// map 中添加键值对
	b[0] = "Hello"
	b[1] = "World"
	fmt.Printf("b的格式化=> %#v\n", b)
	fmt.Printf("b的类型=> %T\n", b)
	// 声明 map 的同时完成初始化
	c := map[int]bool{ 1:true, 2:false, 3:true, 5: false}
	fmt.Printf("map c => %#v\n", c)

	// 判断 map 中的某个键是否存在
	var nameMap = map[int]string{2:"张文杰", 3:"王艳芳", 4:"幼儿猿", 7:"李月华", 8:"王大梦"}
	if val, ok := nameMap[4];ok {
		fmt.Println(val + "在 map 中")
	}else {
		fmt.Println("没有找到!")
	}

	// map 的遍历
	// map 是无序的 与添加的键值对顺序无关
	for  k, v := range nameMap{
		fmt.Println(k, v)
	}
	// map 的删除
	delete(nameMap, 2)
	fmt.Println(nameMap)

	// 生成一个 map
	var numsMap = make(map[string]int, 100)
	for i:=1; i<=50; i++ {
		key := fmt.Sprintf("person%02d", i)
		val := rand.Intn(100)
		numsMap[key] = val
	}
	fmt.Println(numsMap)

	// 遍历map => 无序
	for k, v := range numsMap {
		fmt.Println(k, v)
	}
	// 按照从小到大的顺序遍历
	// 1: 先取出所有的 key 放在切片中
	var keys = make([]string, 0, 100)
	for k := range numsMap{
		keys = append(keys, k)
	}
	//fmt.Println(keys)
	// 2. 对 keys 进行排序
	sort.Strings(keys) // 目前是一个有序的切片
	//fmt.Println(keys)
	// 3. 按照排序后的 key 对 map 进行排序遍历
	for _, k := range keys {
		fmt.Println(k + " =>", numsMap[k])
	}

	// 元素为 map 的切片
	var mapSlice = make([]map[string]int, 8) // 此时只完成了切片的初始化
	fmt.Println(mapSlice[0] == nil) // 此时切片中的 map 尚未完成初始化
	// 完成 map 的初始化
	mapSlice[0] = make(map[string]int, 8)
	mapSlice[0]["hello"] = 222
	fmt.Println(mapSlice)

	// 值为切片的 map 对上述内容类似
	var sliceMap = make(map[string][]int, 10) // 此时只完成了 map 的初始化
	if v, ok := sliceMap["a"]; ok{
		fmt.Println(v)
	}else {
		// sliceMap 中没有 a 这个 key
		// 对切片进行初始化
		sliceMap["a"] = make([]int, 10)
		sliceMap["a"][2] = 222
		sliceMap["a"][4] = 444
		fmt.Println(sliceMap)
	}

	// 小练习
	// 统计一个字符串中每个字母出现的次数 how do you do
	var wordsMap = make(map[string]int, 30)
	var words    = "how do you do"
	words = strings.Replace(words, " ", "", -1)
	var wordsRune = []rune(words)
	for i := 0; i < len(wordsRune); i++ {
		if _, ok := wordsMap[string(wordsRune[i])]; ok {
			wordsMap[string(wordsRune[i])] += 1
		}else {
			wordsMap[string(wordsRune[i])] = 1;
		}
	}
	fmt.Println(wordsMap)
}
