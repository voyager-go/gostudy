package main
import "fmt"

func main() {
	function1()
	af()
	f()
}

func function1() {
	fmt.Printf("In function1 at the top\n")
	defer function2()
	fmt.Printf("In function1 at the bottom!\n")
}

func function2() {
	fmt.Printf("Function2: Deferred until the end of the calling function!\n")
}


// 使用 defer 的语句同样可以接受参数，下面这个例子就会在执行 defer 语句时打印 0：
func af() {
	i := 0
	defer fmt.Println(i)
	i++
	return
}

// 当有多个 defer 行为被注册时，它们会以逆序执行（类似栈，即后进先出）：
func f() {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}
}



// 关键字 defer 允许我们进行一些函数执行完成后的收尾工作，例如：

// 关闭文件流 （详见 第 12.2 节）
// open a file
// defer file.Close()
// 解锁一个加锁的资源 （详见 第 9.3 节）
// mu.Lock()
// defer mu.Unlock()
// 打印最终报告
// printHeader()
// defer printFooter()
// 关闭数据库链接
// open a database connection
// defer disconnectFromDB()