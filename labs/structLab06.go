package main

import "fmt"

// 方法和函数的统一调用
// 声明结构体
type class struct {

}
// 为结构体添加Do方法
func (c *class) Do(v int) {
	fmt.Println("call method do:", v)
}

// 普通函数的Do
func funcDo(v int)  {
	fmt.Println("call method do:", v)
}



// 事件注册
// 实例化一个通过字符串映射函数切片的map
var eventByName = make(map[string][]func(interface{}))
// 注册事件，提供事件名和回调函数
func RegisterEvent(name string, callback func(interface{}))  {
	// 通过名字查找事件列表
	list := eventByName[name]
	// 在列表切片中添加函数
	list = append(list, callback)
	// 将修改的事件列表切片保存回去
	eventByName[name] = list
}

// 调用事件
func CallEvent(name string, param interface{})  {
	// 通过名字找到事件列表
	list := eventByName[name]
	// 遍历这个事件的所有回调
	for _, callback := range list {
		// 传入参数调用回调
		callback(param)
	}
}

//使用事件系统

// 声明角色的结构体
type Actor struct {

}
// 为角色添加一个事件处理函数
func (a *Actor) OnEvent (param interface{}){
	fmt.Println("actor event:", param)
}

// 全局事件
func GlobalEvent(param interface{}) {
	fmt.Println("global event:", param)
}
func main() {
	// 声明一个函数回调
	var delegate func(int)
	// 创建结构体实例
	c := new(class)
	// 将回调设为c的Do方法
	delegate = c.Do
	// 调用
	delegate(100)
	// 将回调设为普通函数
	delegate = funcDo
	delegate(200)

	// 实例化一个角色
	a := new(Actor)
	// 注册名为OnSkill的回调
	RegisterEvent("OnSkill", a.OnEvent)
	// 再次在OnSkill上注册全局事件
	RegisterEvent("OnSkill", GlobalEvent)
	// 调用事件  所有注册的同名函数都会被调用
	CallEvent("Onskill", 11111)
}