package main

import "fmt"

var sliceA = [3]int{1,2,3}
var sliceB = []int{4,5,6}
// 声明字符串切片
var sliceStr []string
// 声明一个空切片
var numListEmpty = []int{}
var highBuilding [30] int
func main() {
	fmt.Println(sliceA, sliceA[1:2])
	for i := 0; i < 30; i++ {
		highBuilding[i] = i + 1
	}
	fmt.Println(highBuilding[10:15])
	// 中间到尾部的所有元素
	fmt.Println(highBuilding[20:])
	// 开头到中间的所有元素
	fmt.Println(highBuilding[:20])
	// 表示原有的切片
	fmt.Println(sliceB)
	fmt.Println(sliceB[:])
	// 重置切片，清空拥有的元素
	fmt.Println(sliceB[0:0])
	// 输出切片大小
	fmt.Println(len(sliceStr), len(numListEmpty))
	// 切片判定空的结果
	fmt.Println(sliceStr == nil)
	// numListEmpty 已经被分配到了内存，但没有元素，因此和 nil 比较时是 false。
	fmt.Println(numListEmpty == nil)
	// 使用 make() 函数构造切片
	sliceC := make([]int, 2)
	sliceD := make([]int, 2, 10)
	fmt.Println(sliceC, sliceD)
	// sliceC 和 sliceD 均是预分配 2 个元素的切片，只是 sliceD 的内部存储空间已经分配了 10 个，但实际使用了 2 个元素。
	// 容量不会影响当前的元素个数，因此 sliceC 和 sliceD 取 len 都是 2。
	fmt.Println(len(sliceC), len(sliceD))


	// append()为切片添加元素
	// Go 语言的内建函数 append() 可以为切片动态添加元素。每个切片会指向一片内存空间，这片空间能容纳一定数量的元素。当空间不能容纳足够多的元素时，切片就会进行“扩容”。“扩容”操作往往发生在 append() 函数调用时。
	// 切片在扩容时，容量的扩展规律按容量的 2 倍数扩充，例如 1、2、4、8、16……
	var numbers []int
	for i := 0; i < 10; i++ {
		numbers = append(numbers, i)
		//  cap() 函数查看切片的容量情况
		fmt.Printf("len: %d cap: %d pointer: %p\n", len(numbers), cap(numbers), numbers)
	}
	// append() 函数除了添加一个元素外，也可以一次性添加很多元素。
	var car []string
	car = append(car, "oldDriver")
	car = append(car, "Ice", "Sniper", "Monk")
	fmt.Println(car)
	// 添加切片
	team := []string{"pig", "flyingCake", "chicken"}
	// 在team后面加上了...，表示将 team 整个添加到 car 的后面。
	car = append(car, team...)
	fmt.Println(car)


	// Go语言copy()：切片复制（切片拷贝）
	// 设置元素数量为1000
	const elementCount  = 1000
	// 预分配足够多的元素切片
	srcData := make([]int, elementCount)
	// 将切片赋值
	for i := 0; i < elementCount; i++ {
		srcData[i] = i
	}
	// 引用切片数据
	refData := srcData
	// 预分配足够多的元素切片
	copyData := make([]int, elementCount)
	// 将数据分配到新的切片空间
	copy(copyData, srcData)
	// 修改原始数据的第一个元素
	srcData[0] = 999
	// 打印引用切片的第一个元素
	fmt.Println(refData[0])
	// 打印复制切片的第一个和最后一个元素
	fmt.Println(copyData[0], copyData[elementCount - 1])
	// 复制原始数据从4到6(不包含)
	copy(copyData, srcData[4:6])
	for i :=0; i < 5; i++ {
		fmt.Printf("%d", copyData[i])
	}
	fmt.Println("============")
	// Go 语言并没有对删除切片元素提供专用的语法或者接口，需要使用切片本身的特性来删除元素
	seq := []string{"a","b","c","d","e"}
	// 指定删除位置
	index := 2
	// 查看删除位置之前的元素和之后的元素
	fmt.Println(seq[:index],seq[index+1:])
	// 将删除点前后的元素连接起来
	seq = append(seq[:index], seq[index+1:]...)
	// Go 语言中切片删除元素的本质是：以被删除元素为分界点，将前后两个部分的内存重新连接起来。
	fmt.Println(seq)
}