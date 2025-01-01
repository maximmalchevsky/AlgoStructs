package main

import (
	"fmt"
)

func main() {
	a := []int{10, 4, 8, 3}
	fmt.Println(leftRightDifference(a))
}

func leftRightDifference(nums []int) []int {
	arr := make([]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		arr[i+1] = nums[i] + arr[i]
	}
	return arr
}
