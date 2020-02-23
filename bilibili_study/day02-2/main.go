package main

import (
	"fmt"
	"reflect"
)

func main() {
	// 数组定义
	// var 变量名 [元素数量]类型(T)
	// var a [3]int
	// var b [5]int
	// 长度是数组类型的一部分
	// a = b  => 会引发 panic 此时 a 与 b 是两种不同的类型
	// 数组的初始化
	// 1. 定义时使用初始化列表的方式进行初始化
	var cityArray = [4]string{"北京", "上海", "广州", "深圳"}
	fmt.Println(cityArray)
	// 2. 使用编译器自行推导数组的长度进行初始化
	var booleanArray = [...]bool{true, true, false, false, true}
	fmt.Println(booleanArray)
	// 3. 使用索引值的方式进行初始化
	var langArray = [...]string{2: "PHP", 4: "Java", 6: "Python", 8: "Golang"}
	fmt.Println(langArray)
	// [  PHP  Java  Python  Golang]
	fmt.Printf("langArray的类型是：%T", langArray)

	// 数组遍历
	// 常规遍历
	for i := 0; i < len(cityArray); i++ {
		fmt.Println(cityArray[i])
	}
	// for range 遍历
	for index, val := range booleanArray {
		fmt.Println(index, val)
	}

	// 多维数组
	// 多维数组外层可以用三个点让编译器推导 内层不可以 var cs = [...][3]string{}
	// 二维数组
	var cities = [3][2]string{
		{"南京", "徐州"},
		{"合肥", "芜湖"},
		{"宁波", "杭州"},
	}
	fmt.Println(cities)
	// 取 芜湖
	fmt.Println(cities[1][1])
	// 二维数组遍历
	for _, v1 := range cities  {
		for _, v2 := range v1 {
			fmt.Println(v2)
		}
	}

	// 计算题
	// 0: 求数组{1, 3, 5, 7, 8}所有元素的和
	var numSum int
	for _, v1 := range [...]int{1, 3, 5, 7, 8}  {
		numSum += v1
	}
	fmt.Println(numSum)
	// 1: 找出数组中和为指定值的两个元素的下标，比如从数组[1, 3, 5, 7, 8]中找出和为8的两个元素的下标分别为(0,3)和(1,2)
	var targetVal int = 8
	var egArray   = [...]int{1, 3, 5, 7, 8}
	for index, v1 := range egArray {
		tmpVal := targetVal - v1
		if isExists, existsIndex := IsExistsItemInArray(tmpVal, egArray);isExists == true {
			fmt.Println(index, existsIndex)
		}
	}
}

func IsExistsItemInArray(item interface{}, array interface{}) (bool, int) {
	switch reflect.TypeOf(array).Kind() {
	case reflect.Array:
		tmpArray := reflect.ValueOf(array)
		for i := 0; i < tmpArray.Len(); i++{
			// 当两个interface的真实值深度相等时，两个interface深度相等
			if reflect.DeepEqual(item, tmpArray.Index(i).Interface()) {
				return true,i
			}
		}
	}
	return false,0
}
// var i1 interface{}
// i1 = "hello"
// var i2 interface{}
// i2 = "hello"
// fmt.Println(reflect.DeepEqual(i1, i2))
