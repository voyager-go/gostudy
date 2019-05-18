package main

import "fmt"

// 传递变长参数
func main() {
	minN := min(3, 4, 5, 6)
	fmt.Println(minN)
}
func min(n ...int) int {
	if len(n) == 0 {
		return 0
	}
	minN := n[0]
	for _, m := range n {
		if m < minN {
			minN = m
		}
	}
	return minN
}