package main

import (
	"fmt"
	"log"
	"os"
)

// Go语言标准库log介绍
func main() {
	log.Println("这是一条很普通的日志。")

	// 配置日志前缀
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	log.Println("这是一条很普通的日志")
	log.SetPrefix("youeryuan=>")
	log.Println("这是一条很普通的日志")

	//log标准库中还提供了一个创建新logger对象的构造函数–New，支持我们创建自己的logger实例
	logger := log.New(os.Stdout, "<Golang>", log.Lshortfile|log.Ldate|log.Ltime)
	logger.Println("自定义日志")

	// Go内置的log库功能有限，例如无法满足记录不同级别日志的情况，我们在实际的项目中根据自己的需要选择使用第三方的日志库，如logrus、zap等
}

// 向 new.log 中写入日志
func init() {
	// SetOutput函数用来设置标准logger的输出目的地，默认是标准错误输出
	logFile, err := os.OpenFile("./new.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644) // 即用户具有读写权限，组用户和其它用户具有只读权限
	if err != nil {
		fmt.Println("open log file failed, err:", err)
		return
	}
	log.SetOutput(logFile)
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
}
