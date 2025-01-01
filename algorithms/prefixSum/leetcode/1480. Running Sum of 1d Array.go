package leetcode

func runningSum(nums []int) []int {
	arr := make([]int, len(nums))
	sum := 0
	for index, item := range nums {
		sum += item
		arr[index] = sum
	}
	return arr
}
