package main

import "fmt"

func Add(a, b int)  {
	fmt.Printf("The sum of %d and %d is %d\n", a, b, a + b)
}

func CallBack(y int, f func(int, int))  {
	f(y, 2)
}

func main() {
	CallBack(1, Add)
}