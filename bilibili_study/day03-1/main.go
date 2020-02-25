package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Printf("current time => %v\n", now)
	year   := now.Year()
	month  := now.Month()
	day    := now.Day()
	hour   := now.Hour()
	minute := now.Minute()
	second := now.Second()
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
	timestamp1 := now.Unix() // 时间戳
	fmt.Println(timestamp1)
	timestamp2 := now.UnixNano() // 纳秒时间戳
	fmt.Println(timestamp2)
	// 使用time.Unix()函数可以将时间戳转为时间格式
	timeObj := time.Unix(1582511059, 0)
	fmt.Println(timeObj)
	year1 := timeObj.Year()
	fmt.Printf("%d\n", year1)

	// 时间操作
	// 当前时间间隔一个小时
	hourLater := time.Now().Add(time.Hour)
	fmt.Println(hourLater.Hour())
	// 求两个时间之间的差值 Sub
	// 判断两个时间是否相等 Equal
	// 如果t代表的时间点在u之前，返回真；否则返回假 func (t Time) Before(u Time) bool
	// 如果t代表的时间点在u之后，返回真；否则返回假 func (t Time) After(u Time) bool

	// 定时器  使用time.Tick(时间间隔)来设置定时器，定时器的本质上是一个通道（channel）
	//ticker := time.Tick(time.Second)
	//for i := range ticker {
	//	fmt.Println(i)
	//}

	// 时间格式化
	// 时间类型有一个自带的方法Format进行格式化，需要注意的是Go语言中格式化时间模板不是常见的Y-m-d H:M:S而是使用Go的诞生时间2006年1月2号15点04分（记忆口诀为2006 1 2 3 4）。也许这就是技术人员的浪漫吧
	// 记忆口诀 2006 1 2 3 4
	fmt.Println(now.Format("2006-01-02 15:04:05 Mon Jan")) // 24小时制
	// 12小时制
	fmt.Println(now.Format("2006-01-02 03:04:05.000 PM Mon Jan"))
	fmt.Println(now.Format("2006/01/02 15:04"))
	fmt.Println(now.Format("15:04 2006/01/02"))
	fmt.Println(now.Format("2006/01/02"))

	// 解析字符串格式的时间
	// 加载时区
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println(err)
		return
	}
	// 按照指定时区和指定格式解析字符串时间
	obj, err := time.ParseInLocation("2006/01/02 15:04:05", "2019/08/04 14:15:20", loc)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(obj)
	fmt.Println(obj.Sub(now))
}