package main

import (
	"fmt"
	"math"
)

type Person struct {
	Name string
	Age  int
}

type Circle struct {
	radius float64
}

func (p Person) speak()  {
	fmt.Println(p.Name + "是个大好人")
}

func (p Person) count() {
	y := 0
	for x := 1; x <= 100; x++  {
		y += x
	}
	fmt.Println(y)
}

func (p Person) count1(n int) {
	y := 0
	for x := 1; x <= n; x++  {
		y += x
	}
	fmt.Println(y)
}

func (p Person) getSum(x, y int) {
	fmt.Println(x + y)
}

func (c Circle) area() float64  {
	return c.radius * c.radius * math.Pi
}

type UserController struct {

}

type MethodUtils struct {

}

type Method struct {
	x int
	y int
	shape string
}

type Method1 struct {
	x float64
	y float64
	shape string
}

func (mu MethodUtils) Print()  {
	for x := 1; x <= 10; x++ {
		for y := 1; y <= 8; y++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}
}

func (mu MethodUtils) Print2(m int, n int) {
	for x := 1; x <= m; x++ {
		for y := 1; y <= n; y++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}
}

func (mu MethodUtils) area(width, length float64) float64  {
	return width * length
}

func (mu MethodUtils) judge(num int)  {
	if num%2 == 0 {
		fmt.Printf("%v是偶数", num)
	} else {
		fmt.Printf("%v是奇数", num)

	}
}

func (method Method) test(m, n int, shape string)  {
	for x := 1; x <= m; x++ {
		for y := 1; y <= n; y++ {
			fmt.Print(shape)
		}
		fmt.Println("")
	}
}

func (method Method1) test1(m, n float64, operator string) float64{
	res := 0.0
	switch operator {
	case "+":
		res = m + n
		return  res
	case "-":
		res = m - n
		return  res
	case "*":
		res = m * n
		return res
	case "/":
		res = m / n
		return res
	default:
		fmt.Println("操作符有误")
		return res
	}
}
func main()  {
	//var person Person
	//person.Name = "张文杰"
	//person.Age  = 25
	//person.speak()
	//person.count()
	//person.count1(10)
	//
	//var (
	//	a, b, c int
	//)
	//fmt.Scanln(&a)
	//person.count1(a)
	//fmt.Scanln(&b)
	//fmt.Scanln(&c)
	//person.getSum(b, c)
	//var (
	//	circle Circle
	//	res   float64
	//)
	//circle.radius = 3.2
	//res = circle.area()
	//fmt.Println(res)
	var mu MethodUtils
	mu.Print()
	mu.Print2(2, 3)
	area := mu.area(10.4 , 2.3)
	fmt.Println(area)
	mu.judge(9)
	var test Method
	test.test(10, 9, "@")
	var x Method1
	res := x.test1(1.0, 2.4, "+")
	fmt.Println(res)
}

