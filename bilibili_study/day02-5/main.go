package main

import "fmt"
var globalVal = 3
func main() {
	fmt.Println(intSum(12, 13, 14))
	fmt.Println(intSum1(12, 14))
	testDefer()
	fmt.Println("结果为：", calc(1, 2, add))
	fmt.Println(globalVal)
	useGlobalOne()
	// 全局变量在函数体内部被改变
	useGlobalTwo()
}
func useGlobalOne()  {
	globalVal++
	fmt.Println(globalVal)
}
func useGlobalTwo()  {
	fmt.Println(globalVal)
}
/**
 * 可变参数
 * 注意可变参数在函数体中是一个切片类型
 */
func intSum(n ...int) int {
	var sumRes int
	for _, val := range n {
		sumRes += val
	}
	return sumRes
}
/**
 * 如果返回值中已经定义了返回的变量  则函数体中不需要重新定义该变量 return 后也加不加上这个变量都可以
 */
func intSum1(a, b int) (res int) {
	res = a + b
	return
}
// defer 延迟执行 defer 后面的语句会在 return 将要执行前执行
// 最先定义的 defer 最后执行
func testDefer() int {
	fmt.Println("begin ===== ")
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	fmt.Println("end   ===== ")
	return 456
}


func add(a, b int) int {
	return a + b
}
// 函数可以当参数 也可以当返回值
func calc(a, b int, op func(int, int)int) int {
	 return op(a, b)
}