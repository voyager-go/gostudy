package main

import (
	"math/rand"
	"fmt"
	"time"
	"text/tabwriter"
	"os"
)
// math/rand
func main() {
	// 1. 首先设置初始化种子
	rand.Seed(time.Now().UnixNano())
	// 2. 返回一个随机的int值
	r := rand.Int()
	fmt.Println("返回一个随机的int值", r)
	// 3. 返回一个[0，n]的随机int值
	r  = rand.Intn(88)
	fmt.Println("返回一个[0，n]的随机int值", r)
	// 4. 返回一个int32类型的非负的31位伪随机数
	fmt.Println(rand.Int31())
	// 5. 返回一个取值范围在[0,n]的伪随机int32值，如果n<=0会panic
	fmt.Println(rand.Int31n(33))
	// 6. 返回一个取值范围在[0, 1]的伪随机float64值
	fmt.Println(rand.Float64())
	// 7. 返回一个服从标准正态分布（标准差=1，期望=0）、取值范围在[-math.MaxFloat64, +math.MaxFloat64]的float64值
	fmt.Println(rand.NormFloat64())
	// 8. 返回一个服从标准指数分布（率参数=1，率参数是期望的倒数）、取值范围在(0, +math.MaxFloat64]的float64值
	fmt.Println(rand.ExpFloat64())
	// 9. 返回一个有n个元素的，[0,n)范围内整数的伪随机排列的切片
	fmt.Println(rand.Perm(8))
	fmt.Println("==========================")
	exampleRand()
	fmt.Println("==========================")
	exampleShuffle()
}

func exampleRand()  {
	// 初始化一个Source，代表一个生成均匀分布在范围[0, 1<<63)的int64值的（伪随机的）资源
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	// 初始化一个过滤器
	w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, '-', 0)
	defer w.Flush()
	show := func(name string, v1, v2, v3 interface{}) {
		fmt.Fprintf(w, "%s\t%v\t%v\t%v\n", name, v1, v2, v3)
	}
	show("Uint32", r.Uint32(), r.Uint32(), r.Uint32())
	show("Intn(10)", r.Intn(10), r.Intn(10), r.Intn(10))
	show("Int31n(10)", r.Int31n(10), r.Int31n(10), r.Int31n(10))
	show("Int63n(10)", r.Int63n(10), r.Int63n(10), r.Int63n(10))
	show("Perm", r.Perm(5), r.Perm(5), r.Perm(5))

}
// 随机交换数切片的位置, n应为切片的长度，n < 0 会panic
func exampleShuffle()  {
	numbers := []byte("12345")
	letters := []byte("ABCDE")
	rand.Shuffle(len(numbers), func(i, j int) {
		numbers[i], numbers[j] = numbers[j], numbers[i]
		letters[i], letters[j] = letters[j], letters[i]
	})
	for i := range numbers {
		fmt.Printf("%c: %c\n", letters[i], numbers[i])
	}
}
