package main

import "fmt"

/**
 * 切片的复制与追加
 */
func main() {
	slice_from := []int{1,2,3}
	slice_to   := make([]int, 10)
	fmt.Println(slice_from, slice_to)
	num := copy(slice_to, slice_from)
	fmt.Println(slice_to)
	fmt.Printf("Copied %d elements \n", num)
	slice_from = append(slice_from, 4,5,6)
	fmt.Println(slice_from)
}