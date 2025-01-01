package leetcode

func pivotIndex(nums []int) int {
	sumRight := 0
	for _, num := range nums {
		sumRight += num
	}

	sumLeft := 0
	for i, num := range nums {
		sumRight -= num
		if sumRight == sumLeft {
			return i
		}
		sumLeft += num
	}

	return -1
}
