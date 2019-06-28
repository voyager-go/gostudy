package main

import (
	"fmt"
	"sort"
)

// 不要使用 new，永远用 make 来构造 map
func main() {
	var mapLit map[string]int
	var mapAssigned map[string]int
	mapLit = map[string]int{"one": 1, "two": 2}
	mapCreated := make(map[string]float32)
	mapAssigned = mapLit
	mapCreated["k1"] = 4
	mapCreated["k2"] = 3.4
	mapAssigned["two"] = 3

	fmt.Printf("Map literal at \"one\" is: %d\n", mapLit["one"])
	fmt.Printf("Map created at \"k2\" is: %f\n", mapCreated["k2"])
	fmt.Printf("Map assigned at \"two\" is: %d\n", mapLit["two"])
	fmt.Printf("Map literal at \"ten\" is: %d\n", mapLit["ten"])

	fmt.Println("====================================")

	mf := map[int]func() int{
		1: func() int { return 10 },
		2: func() int { return 20 },
		5: func() int { return 50 },
	}
	fmt.Println(mf)

	map_a := make(map[string]int, 10)
	map_a["a"] = 1
	map_a["b"] = 2
	map_a["c"] = 3
	// 判断map中的某个key是否存在
	if _, ok := map_a["b"]; ok {
		fmt.Println(map_a["b"])
	}
	// 删除某个key
	delete(map_a, "a")
	if _, ok := map_a["a"]; ok {
		fmt.Println(map_a["a"])
	}

	// for range ...
	map_b := make(map[int]float32)
	map_b[2] = 2.2
	map_b[3] = 2.3
	map_b[4] = 2.4
	map_b[5] = 2.5
	// 注意 输出的 map 不是按照 key 的顺序排列的，也不是按照 value 的序排列的。
	for key, val := range map_b {
		fmt.Printf("key is: %d - value is: %f\n", key, val)
	}

	capitals := map[string]string{"Italy": "Rome", "France": "Paris", "Japan": "Tokyo"}
	for key := range capitals {
		fmt.Println("Map item: Capital of", key, "is", capitals[key])
	}

	// 假设我们想获取一个 map 类型的切片，我们必须使用两次 make() 函数，第一次分配切片，第二次分配 切片中每个 map 元素
	items := make([]map[int]int, 5)
	for i := range items {
		items[i] = make(map[int]int, 1)
		items[i][1] = 2
	}
	fmt.Printf("Version A: Value of items: %v\n", items)
	items2 := make([]map[int]int, 5)
	for _, item2 := range items2 {
		item2 = make(map[int]int, 1)
		item2[1] = 2
	}
	fmt.Printf("Version B: Value of items: %v\n", items2)

	// map 的排序map 默认是无序的，不管是按照 key 还是按照 value 默认都不排序
	// 如果你想为 map 排序，需要将 key（或者 value）拷贝到一个切片，再对切片排序（使用 sort 包，详见第 7.6.6 节），然后可以使用切片的 for-range 方法打印出所有的 key 和 value。
	var (
		barVal = map[string]int{
			"alpha": 34, "bravo": 56, "charlie": 23,
			"delta": 87, "echo": 56, "foxtrot": 12,
			"golf": 34, "hotel": 16, "indio": 87,
			"juliet": 65, "kili": 43, "lima": 98,
		}
	)
	fmt.Println("unsorted:")
	for k, v := range barVal {
		fmt.Printf("Key: %v, Value: %v / ", k, v)
	}
	keys := make([]string, len(barVal))
	i := 0
	for k, _ := range barVal {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	fmt.Println("\nsorted:")
	for _, k := range keys {
		fmt.Printf("Key: %v, Value: %v / ", k, barVal[k])
	}

	// 将 map 的键值对调
	invMap := make(map[int]string, len(barVal))
	for k, v := range barVal {
		invMap[v] = k
	}
	fmt.Println("\ninverted:")
	for k, v := range invMap {
		fmt.Printf("Key: %v, Value: %v / ", k, v)
	}
}
