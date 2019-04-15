package main

import (
	"container/list"
	"fmt"
)

func main() {
	// 声明一个列表
	listA := list.New()
	// 向列表尾部追加
	listA.PushBack("canon")
	// 向列表头部追加
	listA.PushFront(67)
	// 尾部追加后保存元素句柄
	ele := listA.PushBack("first")
	// 在first之后追加second
	listA.InsertAfter("second", ele)
	// 在first之前追加init
	listA.InsertBefore("init", ele)
	// 移除 element 变量对应的元素 => first的句柄 所以first被移除
	listA.Remove(ele)
	// 遍历列表——访问列表的每一个元素
	for i := listA.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
}