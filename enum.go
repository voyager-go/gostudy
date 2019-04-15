package main

import "fmt"
// Go 语言中现阶段没有枚举，可以使用 const 常量配合 iota 模拟枚举，请看下面的代码：
// Weapon -> 兵器
type Weapon int

const (
	Arrow Weapon = iota	// 开始生成枚举值 默认为0
	Shuriken
	SniperRifle
	Rifle
	Blower
)
// 声明芯片类型
type ChipType int

const (
	None ChipType = iota
	CPU
	GPU
)

func (c ChipType) String() string  {
	switch c {
	case None:
		return "None"
	case CPU:
		return "CPU"
	case GPU:
		return "GPU"
	}
	return "N/A"
}
func main() {
	// 输出所有枚举值
	fmt.Println(Arrow, Shuriken, SniperRifle, Rifle, Blower)
	// 使用枚举类型并赋初值
	var weapon Weapon  = Rifle
	fmt.Println(weapon)
	//输出 CPU 的值并按整型格式输出。
	// 当这个类型需要显示为字符串时，Go 语言会自动寻找 String() 方法并进行调用。
	fmt.Printf("%s %d", CPU, CPU)
	fmt.Printf("%s %d", CPU.String(), CPU)

}