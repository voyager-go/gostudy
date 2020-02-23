package main

import (
	"fmt"
	"sort"
)

func main() {
	// 切片
	var a []int
	var b = []bool{true, false, true}
	fmt.Println(a, b)
	// 基于数组得到切片
	var c = [5]int{7, 2, 4, 5, 9}
	var d = c[1:3]
	fmt.Println(d)
	// 查看类型
	fmt.Printf("切片d: %T\n", d)
	// 从切片得到切片
	e := d[:]
	// 查看当前容量 d => [2,4]
	// 为什么当前容量为 4 ？ 因为容量和 d 一样，而 d 的容量是从 c 切片分离出的起点一直到终点
	//    c: 7, 2, 4, 5, 9
	//    d:    2, 4
	// d容量:    --------- 从起点到结束都是容量
	fmt.Println("当前 e 的容量：", cap(e))
	// : 相当于 0:len(slice)
	f := d[0:len(d)]
	fmt.Println(e, f)
	// make 函数构造切片
	// 5表示切片的长度 10表示容量
	var g = make([]int, 5, 10)
	fmt.Println(g)
	fmt.Printf("%T\n", g)
	fmt.Println("切片长度：", len(g))
	fmt.Println("切片容量：", cap(g))

	var h []string     // 声明 string 类型切片 => nil
	var j = []string{} // 声明并且初始化 和上面的单纯声明是不同的
	var k = make([]string, 0)
	fmt.Println(h == nil)
	fmt.Println(j == nil)
	fmt.Println(k == nil)
	// 判断一个切片是否为空用 len() 是否等于 0
	fmt.Println(len(k))
	// 切片的赋值拷贝
	var l = make([]int, 4)
	var m = l
	m[2] = 199
	// 切片是引用类型
	// 结果为 [0 0 199 0]  说明 切片l 与 切片m 底层都指向了同一个数组
	fmt.Println(l)
	// 切片的扩容操作   切片一定要初始化以后才可以使用
	var n []int // 此时它并没有申请内存
	// n[0] = 100
	// panic: runtime error: index out of range [0] with length 0 数组越界
	fmt.Println(n)
	// 切片扩容演示
	for i := 0; i < 10; i++ {
		n = append(n, i)
		fmt.Printf("slice => %v, len => %d, cap => %d, ptr => %p\n", n, len(n), cap(n), n)
	}
	var p = []int{12, 14, 15, 16}
	n = append(n, p...)
	fmt.Println(n)
	// 切片的 copy
	aa := []int{1, 3, 4, 5, 6, 7}
	bb := make([]int, len(aa))
	copy(bb, aa)
	fmt.Println(bb)
	// 切片是引用类型
	aa[1] = 33
	fmt.Println(aa, bb) // [1 33 4 5 6 7] [1 3 4 5 6 7]

	// 使用切片删除元素
	cc := []int{1, 2, 3, 4, 5}
	// 删除 3
	cc = append(cc[:2], cc[3:]...)
	fmt.Println(cc)
	// 公式为 append(cc[:index], [index+1:]...)

	// 使用 sort 函数对切片进行排序
	var dd = [5]int{7, 5, 6, 2, 5}
	// 这里需要传一个切片 而切片底层指向的是数组 dd
	sort.Ints(dd[:])
	fmt.Println(dd)
}
