package leetcode

func isCovered(ranges [][]int, left int, right int) bool {
	arr := make([]int, 52)

	for i := 0; i < len(ranges); i++ {
		arr[ranges[i][0]]++
		arr[ranges[i][1]+1]--
	}

	for i := 1; i < 52; i++ {
		arr[i] += arr[i-1]
	}

	for i := left; i <= right; i++ {
		if arr[i] < 1 {
			return false
		}
	}
	return true
}
