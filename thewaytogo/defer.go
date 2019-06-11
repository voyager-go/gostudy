package main
import "fmt"

func main() {
	fmt.Println(td1(), td2(), td3())
}

func td1() int  {
	n := 3
	defer func() {
		n =  2
	}()
	n++
	return n	// 按值传递，n的copy已经传给了"返回值"，defer对n做修改，并不能影响到n的copy(返回值)
}

func td2() (result int){
	defer func() {
		result++
	}()	// 闭包内的result是引用，在原始result上的操作
	return 0
}

func td3() (r int)  {
	defer func(r int) {
		r = r + 3
	}(r)	// 按值传递，在r的copy上的操作，r还是0
	return
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
var ii int = 0
func dFun() {
	defer func() {
		fmt.Println(ii)
	}()
	ii++
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