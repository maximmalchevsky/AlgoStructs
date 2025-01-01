package leetcode

import "math"

func leftRightDifference(nums []int) []int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}

	leftSum := 0
	for i, num := range nums {
		nums[i] = int(math.Abs(float64(sum - leftSum - num - leftSum)))
		leftSum += num
	}

	return nums
}
