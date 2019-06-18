package main

import "fmt"

/**
 *切片拥有 长度 和 容量。
 *切片的长度就是它所包含的元素个数。
 *切片的容量是从它的第一个元素开始数，到其底层数组元素末尾的个数。
 *切片 s 的长度和容量可通过表达式 len(s) 和 cap(s) 来获取。
 */
func main() {
	sl := make([]int, 0, 10)
	for i :=0; i < cap(sl); i++ {
		sl = sl[0:i+1]
		sl[i] = i
		fmt.Printf("The length of slice is %d\n", len(sl))
		fmt.Println(sl)
	}
	var ar = [10]int{0,1,2,3,4,5,6,7,8,9}
	fmt.Println(cap(ar[5:7]))

}
