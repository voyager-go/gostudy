// @Title
// @Description
// @Author       幼儿猿
// @Datetime     2020/2/26 2:39 下午
package main

import (
	"fmt"
	"sort"
)

func main() {
	nums1 := []int{1, 3}
	nums2 := []int{2, 5}
	newNums := append(nums1, nums2...)
	sort.Ints(newNums)
	middleIndex := len(newNums) / 2
	var result float64
	if len(newNums)%2 == 0 {
		index1, index2 := middleIndex-1, middleIndex
		result = (float64)(newNums[index1]+newNums[index2]) / 2
	} else {
		index := int(middleIndex)
		result = float64(newNums[index])
	}
	fmt.Println(result)
}
