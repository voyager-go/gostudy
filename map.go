package main

import (
	"fmt"
	"sort"
	"sync"
)

// 在业务和算法中需要使用任意类型的关联关系时，就需要使用到映射，如学号和学生的对应、名字与档案的对应等。
// Go 语言提供的映射关系容器为 map，map使用散列表（hash）实现。
// 提示
// 大多数语言中映射关系容器使用两种算法：散列表和平衡树。

// 散列表可以简单描述为一个数组（俗称“桶”），数组的每个元素是一个列表。根据散列函数获得每个元素的特征值，将特征值作为映射的键。如果特征值重复，表示元素发生碰撞。碰撞的元素将被放在同一个特征值的列表中进行保存。散列表查找复杂度为 O(1)，和数组一致。最坏的情况为 O(n)，n 为元素总数。散列需要尽量避免元素碰撞以提高查找效率，这样就需要对“桶”进行扩容，每次扩容，元素需要重新放入桶中，较为耗时。

// 平衡树类似于有父子关系的一棵数据树，每个元素在放入树时，都要与一些节点进行比较。平衡树的查找复杂度始终为 O(log n)。
func main() {
	scene := make(map[string]int)
	scene["route"] = 66
	fmt.Println(scene["route"])
	// 尝试查找一个不存在的键，那么返回的将是 ValueType 的默认值。
	mapA := scene["route1"]
	fmt.Println(mapA)
	// 某些情况下，需要明确知道查询中某个键是否在 map 中存在，可以使用一种特殊的写法来实现
	mapB, ok := scene["route2"]
	fmt.Println(mapB, ok)
	// map 还有一种在声明时填充内容的方式
	words := map[string]string{
		"W": "forward",
		"A": "left",
		"D": "right",
		"S": "backward",
	}
	fmt.Println(words)
	// map 的遍历
	for k, v := range words {
		fmt.Println(k, v)
	}
	// 只遍历值 将不需要的键改为匿名变量形式
	for _, v := range scene{
		fmt.Println(v)
	}
	// 只遍历键
	for k := range words {
		fmt.Println(k)
	}
	// 排序遍历
	scene["foo"] = 33
	scene["bar"] = 44
	// 声明 sceneList 为字符串切片，以缓冲和排序 map 中的所有元素
	var sceneList []string
	for k := range scene {
		sceneList = append(sceneList, k)
	}
	// 对切片进行排序
	sort.Strings(sceneList)
	fmt.Println(sceneList)
	// 使用 delete() 函数从 map 中删除键值对
	fmt.Println(scene)
	delete(scene, "foo")
	fmt.Println(scene)


	// sync.Map
	// 无须初始化，直接声明即可。
	// sync.Map 不能使用 map 的方式进行取值和设置等操作，而是使用 sync.Map 的方法进行调用。Store 表示存储，Load 表示获取，Delete 表示删除。
	// 使用 Range 配合一个回调函数进行遍历操作，通过回调函数返回内部遍历出来的值。Range 参数中的回调函数的返回值功能是：需要继续迭代遍历时，返回 true；终止迭代遍历时，返回 false。

	var newScene sync.Map
	// 将键值对保存到sync.Map
	newScene.Store("green", 98)
	newScene.Store("london", 100)
	newScene.Store("egypt", 188)
	// 从sync.Map中根据键取值
	fmt.Println(newScene.Load("london"))
	// 根据键删除对应的键值对
	newScene.Delete("london")
	// 遍历所有sync.Map中的键值对
	newScene.Range(func(key, value interface{}) bool {
		fmt.Println("iterate: ", key, value)
		return true
	})

}