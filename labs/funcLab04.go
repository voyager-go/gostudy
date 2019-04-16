package main

import "fmt"
// 闭包是引用了自由变量的函数，被引用的自由变量和函数一同存在，即使已经离开了自由变量的环境也不会被释放或者删除，在闭包中可以继续使用这个自由变量
// 在闭包内部修改引用的变量
func main() {
	str := "hello word"
	foo := func() {
		str = "hello dude"
	}
	foo()
	fmt.Println(str)

	accumulator := Accumulate(1)
	fmt.Println(accumulator())
	fmt.Println(accumulator())
	fmt.Printf("%p\n", accumulator)
	accumulator2 := Accumulate(10)
	fmt.Println(accumulator2())
	fmt.Printf("%p\n", accumulator2)

	//=====================================
	// 创建一个玩家生成器
	generator := playerGen("high noon")
	// 返回玩家的名字和血量
	name, hp := generator()
	fmt.Println(name, hp)
}

// 闭包的记忆效应
// 提供一个值, 每次调用函数会指定对值进行累加
func Accumulate(value int) func() int  {
	return func() int {
		value++
		return value
	}
}

// 闭包实现生成器
func playerGen(name string) func()(string, int)  {
	// 血量一直是150
	hp := 150
	// 返回创建的闭包
	return func() (string, int) {
		// 将变量引用到闭包中
		return name,hp
	}
}