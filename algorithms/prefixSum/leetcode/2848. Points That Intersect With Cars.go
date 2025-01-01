package leetcode

func numberOfPoints(nums [][]int) int {
	arr := make([]int, 102)
	ans, count := 0, 0
	for i := 0; i < len(nums); i++ {
		arr[nums[i][0]]++
		arr[nums[i][1]+1]--
	}

	for i := 0; i < 102; i++ {
		count += arr[i]
		if count != 0 {
			ans++
		}
	}
	return ans
}
