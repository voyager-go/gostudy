package main

import "fmt"

// 为结构体添加方法
type Bag struct {
	items []int
}

func Insert(b *Bag, itemid int)  {
	b.items = append(b.items, itemid)
}
/*
接收器
func (接收器变量 接收器类型) 方法名(参数列表) (返回参数) {
    函数体
}
 */
func (b *Bag)NewInsert(itemid int)  {
	b.items = append(b.items, itemid)
}
// 指针类型的接收器
// 定义属性结构
type Property struct {
	value int
}
// 设置属性值
func (p *Property) SetValue(v int)  {
	// 修改p的成员变量
	p.value = v
}
// 取属性值
func (p *Property) getValue() int  {
	return p.value
}

// 非指针类型的接收器
type Pointerr struct {
	X int
	Y int
}
// 非指针接收器的加方法
func (p Pointerr) Add(other Pointerr) Pointerr {
	// 成员值与参数相加后返回新的结构
	return Pointerr{p.X + other.X, p.Y + other.Y}
}
func main()  {
	// 过程式
	bag := new(Bag)
	Insert(bag, 10001)
	fmt.Println(bag)

	// 结构体方法
	b := new(Bag)
	b.NewInsert(1002)
	fmt.Println(b)

	// 实例化属性
	p := new(Property)
	p.SetValue(1000)
	fmt.Println(p.getValue())


	// 初始化点
	p1 := Pointerr{1, 1}
	p2 := Pointerr{2, 2}
	result := p1.Add(p2)
	fmt.Println(result)
}