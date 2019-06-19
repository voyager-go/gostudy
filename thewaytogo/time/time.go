package main

import (
	"fmt"
	"time"
)
var week time.Duration
func main()  {
	t := time.Now()
	fmt.Println(t)
	fmt.Println(t.Date())
	fmt.Println(t.Day())
	fmt.Println(t.Month())
	fmt.Println(t.Year())
	// 格式化时间
	fmt.Printf("%02d,%02d,%04d\n", t.Day(), t.Month(), t.Year())
	t = time.Now().UTC()
	fmt.Println(t)
	// 1e9 表示10的9次方 也就是10亿
	week = 60 * 60 * 24 * 7 * 1e9
	weekFromNow := t.Add(time.Duration(week))
	fmt.Println(weekFromNow)
	fmt.Println(t.Format("01 March 2019 22:13"))
}